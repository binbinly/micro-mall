"use strict";

const Service = require("egg").Service;
const moment = require("moment");
const EtherToken = require("../utils/ether_token");
const Tron = require("../utils/tron");
const { Decimal } = require("decimal.js");

class BalanceService extends Service {
  table = "chain_balance";

  /**
   * 发送交易至汇总钱包
   * @returns
   */
  async sendTransaction() {
    const { app, service } = this;
    const list = await app.mysql.query(
      `select * from ${this.table} where status = 1 order by amount desc limit 100`
    );
    if (list.length == 0) {
      return;
    }
    //汇总钱包地址
    const receiveAddress = await service.config.value(
      service.config.tron_receive_key
    );
    if (!receiveAddress) {
      return;
    }
    list.forEach(async (item) => {
      const account = await service.account.detail(item.address);
      if (!account) return;

      if (item.type == 1) {
        await this.sendEthToken(item, account);
      } else if (item.type == 2) {
        await this.sendTronToken(item, account, receiveAddress);
      }
    });
  }

  /**
   * 以太坊 收款钱包发送代币至汇总钱包
   * @param {*} item
   * @returns
   */
  async sendEthToken(item, account) {
    const { app, ctx } = this;

    const txObj = JSON.parse(item.tx);
    if (!txObj) {
      return;
    }
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    let insertData = {
      from: txObj.from,
      to: txObj.to,
      value: item.amount,
      created_at: now,
    };

    const ether = new EtherToken(app.config.bsc.web3Url);
    const txData = ether.signTransaction(txObj, app.aesEcbDecrypt(account.key));
    ether
      .sendTransaction(txData)
      .on("transactionHash", async (hash) => {
        insertData.hash = hash;
        return await app.mysql.insert("chain_trans_package", insertData);
      })
      .on("receipt", async (receipt) => {
        await app.mysql.update(
          "chain_trans_package",
          { status: 1 },
          { where: { hash: receipt.transactionHash } }
        );
        //账户地址初始化
        this.initStatus(item.address, item.contract);
      })
      .on("error", async (error) => {
        ctx.logger.error("sendTransaction err:%s", error);
      });
  }

  /**
   * 波场 收款钱包发送代币至汇总钱包
   * @param {*} item
   * @param {*} receiveAddress
   * @returns
   */
  async sendTronToken(item, account, receiveAddress) {
    const { app } = this;

    const tron = new Tron(
      app.config.tron.api,
      app.config.tron.apiKey,
      app.aesEcbDecrypt(account.key)
    );
    const txid = await tron.sendTokenTransaction(
      item.contract,
      item.address,
      receiveAddress
    );
    if (txid) {
      const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
      await app.mysql.insert("chain_trans_tron_package", {
        from: item.address,
        to: receiveAddress,
        value: item.amount,
        txid,
        status: 1,
        created_at: now,
      });
      //初始化
      await this.initStatus(item.address, item.contract);
    } else if (txid === 0) {
      //钱包中已木有余额，同样初始化
      await this.initStatus(item.address, item.contract);
    }
  }

  /**
   * 发送 gas 至收款钱包
   */
  async sendGasTransaction() {
    const { app, service } = this;
    const list = await app.mysql.query(
      `select * from ${this.table} where status = 0 and amount > ? order by amount desc limit 100`,
      [app.config.minAmount]
    );
    if (list.length == 0) {
      return;
    }
    //发送gas钱包
    const etherWallet = await service.config.getWallet(
      service.config.ether_wallet_key
    );
    if (!etherWallet) {
      return;
    }
    const tronWallet = await service.config.getWallet(
      service.config.tron_wallet_key
    );
    if (!tronWallet) {
      return;
    }
    //汇总钱包地址
    const etherReceiveAddress = await service.config.value(
      service.config.ether_receive_key
    );
    if (!etherReceiveAddress) {
      return;
    }
    const ether = new EtherToken(app.config.bsc.web3Url);
    const tron = new Tron(
      app.config.tron.api,
      app.config.tron.apiKey,
      tronWallet.key
    );
    list.forEach(async (item) => {
      if (item.type == 1) {
        ether.contractInit(app.config.contractPath, item.contract);
        await this.sendEthGas(ether, item, etherReceiveAddress, etherWallet);
      } else if (item.type == 2) {
        await this.sendTronGas(tron, item, tronWallet);
      }
      await app.sleep(1000);
    });
  }

  /**
   * 发送 eth 至收款钱包以转出到汇总钱包
   * @param {EtherToken} ether
   * @param {*} item
   * @param {*} receiveAddress
   * @param {*} wallet
   */
  async sendEthGas(ether, item, receiveAddress, wallet) {
    const { ctx, app } = this;
    //获取主账号 eth 是否足够
    const mainEth = await ether.getBalance(wallet.address);
    //获取当前账户的 eth
    const eth = await ether.getBalance(item.address);
    const txObj = await ether.buildTokenTransaction(
      item.contract,
      item.address,
      receiveAddress
    );
    const gasEth = ether.calcGas(txObj.gasPrice, txObj.gasLimit);
    if (mainEth < gasEth) {
      throw new Error("eth gas insufficient");
    }
    if (eth >= gasEth) {
      return;
    }
    //构建发送gas交易对象
    const txGasObj = await ether.buildEthTransaction(
      wallet.address,
      item.address,
      gasEth
    );
    const nonce = await ether.getAddressNonce(wallet.address);
    txGasObj.nonce = nonce;
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    let insertData = {
      to: txGasObj.to,
      value: new Decimal(txGasObj.value).div(10 ** 12).toString(),
      created_at: now,
    };
    //签名交易
    const txData = ether.signTransaction(txGasObj, wallet.key);
    //发送交易
    ether
      .sendTransaction(txData)
      .on("transactionHash", async (hash) => {
        insertData.hash = hash;
        return await app.mysql.insert("chain_trans_gas_package", insertData);
      })
      .on("receipt", async (receipt) => {
        await app.mysql.update(
          "chain_trans_gas_package",
          { is_success: 1 },
          { where: { hash: receipt.transactionHash } }
        );
      })
      .on("error", async (error) => {
        ctx.logger.error("sendTransaction gas err:%s", error);
      });
    //设置提现中状态
    await app.mysql.update(this.table, {
      id: item.id,
      status: 1,
      tx: JSON.stringify(txObj),
    });
  }

  /**
   * 发送 trx 给收款钱包
   * @param {Tron} tron
   * @param {*} item
   * @param {*} wallet
   * @returns
   */
  async sendTronGas(tron, item, wallet) {
    const { app } = this;
    //获取主账号trx是否足够
    const mainTrx = await tron.getBalance(wallet.address);
    //获取当前账户的trx
    const trx = await tron.getBalance(item.address);
    //获取当前交易需要的trx
    const gasTrx = await tron.calcTronGas(
      item.contract,
      app.config.tron.gasAddress || wallet.address,
      item.address,
      item.amount
    );
    if (mainTrx < gasTrx) {
      throw new Error("tron gas insufficient");
    }
    //设置提现中状态
    await app.mysql.update(this.table, {
      id: item.id,
      status: 1,
    });
    if (trx >= gasTrx) {
      //手续费足够
      return;
    }
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    //需要转入手续费
    const trxToAmount = gasTrx;
    const trxRes = await tron.sendTransaction(item.address, trxToAmount);
    if (trxRes.result) {
      await app.mysql.insert("chain_trans_tron_gas_package", {
        to: item.address,
        value: trxToAmount,
        txid: trxRes.txid,
        created_at: now,
      });
    }
  }

  //增加币
  async add(address, contract, token, amount, type, conn) {
    const balance = await conn.get(this.table, { address, contract });
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    if (balance) {
      //更新钱包币数量
      return await conn.update(this.table, {
        id: balance.id,
        amount: balance.amount + amount,
        updated_at: now,
      });
    } else {
      return await conn.insert(this.table, {
        address,
        contract,
        token,
        amount,
        type,
        created_at: now,
        updated_at: now,
      });
    }
  }

  async initStatus(address, contract) {
    await this.app.mysql.update(
      this.table,
      { status: 0, amount: 0 },
      {
        where: {
          address,
          contract,
        },
      }
    );
  }
}

module.exports = BalanceService;

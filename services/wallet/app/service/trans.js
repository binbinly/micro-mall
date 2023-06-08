"use strict";

const Service = require("egg").Service;
const moment = require("moment");
const Ether = require("../utils/ether");

class TransService extends Service {
  table = "chain_trans_eth";

  /**
   * api获取以太坊代币交易记录
   * @param {*} chainId
   * @param {*} url
   * @param {*} contractaddress
   * @param {*} apikey
   * @returns
   */
  async getTokenTxListByApi(chainId, url, apikey, contractaddress, token) {
    const { ctx, app, service } = this;
    const db = ctx.app.mysql;
    const key = "eth_" + chainId + contractaddress;
    const info = await service.config.detail(key);
    if (!info) {
      const ether = new Ether(app.config.bsc.web3Url);
      const value = await ether.getBlockNumber();
      await db.insert(service.config.table, { key, value });
      return;
    }
    const startblock = info.value;
    let endBlock = startblock;

    const offset = 100;
    let page = 1;
    while (true) {
      try {
        const list = await ctx.helper.getEthTransactionList(url, {
          module: "account",
          action: "tokentx",
          contractaddress,
          startblock,
          page,
          offset,
          sort: "asc",
          endblock: 99999999,
          apikey,
        });
        if (list.length == 0) {
          break;
        }
        let m = {};
        for (const transaction of list) {
          if (!m[transaction.to.toLowerCase()]) {
            m[transaction.to.toLowerCase()] = [transaction];
          } else {
            m[transaction.to.toLowerCase()].push(transaction);
          }
          if (transaction.blockNumber > endBlock) {
            endBlock = transaction.blockNumber;
          }
        }
        await this.match(chainId, contractaddress, token, m);
        await db.update(service.config.table, { id: info.id, value: endBlock });
        if (list.length < offset) {
          break;
        }
        await app.sleep(1000);
        page++;
      } catch (e) {
        ctx.logger.error("getTokenTxListByApi usdt err:%s", e);
        break;
      }
    }
  }

  /**
   * api获取以太坊交易记录
   * @param {*} chainId
   * @param {*} url
   * @param {*} contractaddress
   * @param {*} apikey
   * @returns
   */
  async getTxListByApi(chainId, url, apikey, contractaddress, token) {
    const { ctx, app, service } = this;
    const db = ctx.app.mysql;
    const key = "eth_" + chainId + contractaddress;
    const info = await service.config.detail(key);
    const ether = new Ether(app.config.bsc.web3Url);
    if (!info) {
      const value = await ether.getBlockNumber();
      await db.insert(service.config.table, { key, value });
      return;
    }
    const startblock = info.value;
    let endBlock = startblock;

    const offset = 100;
    let page = 1;
    while (true) {
      try {
        const list = await ctx.helper.getEthTransactionList(url, {
          module: "account",
          action: "txlist",
          address: contractaddress,
          startblock,
          page,
          offset,
          sort: "asc",
          endblock: "99999999",
          apikey,
        });
        console.log("list", startblock, list.length);
        if (list.length == 0) {
          break;
        }
        let m = {};
        for (let transaction of list) {
          if (transaction.to.toLowerCase() != contractaddress.toLowerCase()) {
            continue;
          }
          if (transaction.isError != "0") {
            continue;
          }
          const inputData = ether.decodeERC20TransferParameters(transaction);
          if (!inputData) continue;
          transaction.to = inputData.to;
          transaction.value = inputData.amount;
          if (!m[transaction.to.toLowerCase()]) {
            m[transaction.to.toLowerCase()] = [
              ctx.helper.handleEthTransaction(transaction),
            ];
          } else {
            m[transaction.to.toLowerCase()].push(
              ctx.helper.handleEthTransaction(transaction)
            );
          }
          if (transaction.blockNumber > endBlock) {
            endBlock = transaction.blockNumber;
          }
        }
        await this.match(chainId, contractaddress, token, m);
        await db.update(service.config.table, { id: info.id, value: endBlock });
        if (list.length < offset) {
          break;
        }
        await app.sleep(1000);
        page++;
      } catch (e) {
        ctx.logger.error("getTxListByApi err:%s", e);
        break;
      }
    }
  }

  /**
   * 交易检查，是否确认
   * @returns
   */
  async check() {
    const { app } = this;
    const limit = 50;
    const list = await app.mysql.select(this.table, {
      where: { is_confirm: 0 },
      limit,
    });
    if (list.length == 0) {
      return 0;
    }
    const ether = new Ether(this.app.config.bsc.web3Url);
    //当前区块高度
    const curBlock = await ether.getBlockNumber();
    list.forEach(async (item) => {
      if (curBlock - item.block_number < 15) {
        return;
      }
      //获取交易详情
      const receipt = await ether.getTransactionReceipt(item.hash);
      //交易成功，并且必须相隔15个区块
      const diffBlock = curBlock - receipt.blockNumber;
      if (receipt.status == true && diffBlock > 15) {
        //交易确认
        await this.submit(item, curBlock);
      }
    });
  }

  /**
   * 匹配平台钱包地址的所有交易并保存入库
   * @param {*} chain_id
   * @param {*} contract
   * @param {*} token
   * @param {*} list
   * @returns
   */
  async match(chain_id, contract, token, list) {
    const { app, ctx } = this;
    const addresses = Object.keys(list);
    if (addresses.length == 0) {
      return;
    }
    //检查出平台的钱包地址
    const results = await app.mysql.query(
      "select address from chain_account where type=1 and address in (?)",
      [addresses]
    );
    if (results.length == 0) {
      return;
    }
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    results.forEach(async (val) => {
      //匹配交易
      let transList = list[val.address.toLowerCase()];
      if (transList) {
        transList.forEach(async (trans) => {
          if (await app.mysql.get(this.table, { hash: trans.hash })) {
            return;
          }
          //保存交易
          await app.mysql.insert(this.table, {
            from: trans.from,
            to: trans.to,
            value: ctx.helper.coinToDB(trans.value),
            token,
            contract,
            hash: trans.hash,
            block_number: trans.blockNumber,
            timestamp: trans.timeStamp,
            chain_id,
            created_at: now,
            updated_at: now,
          });
        });
      }
    });
  }

  /**
   * 交易确认
   * @param {*} item
   * @param {*} curBlock
   * @returns
   */
  async submit(item, curBlock) {
    const { app, service } = this;
    //事务处理
    const conn = await app.mysql.beginTransaction(); // 初始化事务
    try {
      //链上余额记录，转账到汇总钱包使用
      await service.balance.add(
        item.to,
        item.contract,
        item.token,
        item.value,
        1,
        conn
      );
      //修改状态
      await conn.update(this.table, {
        id: item.id,
        is_confirm: 1,
        confirm_block: curBlock,
      });
      await conn.commit(); // 提交事务
    } catch (err) {
      await conn.rollback(); // 一定记得捕获异常后回滚事务！！
      throw err;
    }
  }
}

module.exports = TransService;

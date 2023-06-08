"use strict";

const TronWeb = require("tronweb");
const { ethers } = require("ethers");
const { Decimal } = require("decimal.js");

const AbiCoder = ethers.utils.AbiCoder;
const ADDRESS_PREFIX = "41";

module.exports = {
  /**
   * 生成以太坊钱包
   * @returns
   */
  createEtherWallet() {
    const privateKey = ethers.utils.randomBytes(32);
    const wallet = new ethers.Wallet(privateKey);
    const key = ethers.BigNumber.from(privateKey)._hex;
    const encKey = this.app.aesEcbEncrypt(key);
    return { address: wallet.address, key: encKey };
  },

  /**
   * 生成波场钱包
   * @returns
   */
  async createTronWallet() {
    const info = await TronWeb.createAccount();
    const encKey = this.app.aesEcbEncrypt(info.privateKey);
    return { address: info.address.base58, key: encKey };
  },

  /**
   * 币入库放大 1000000倍
   * @param {*} value
   * @returns
   */
  coinToDB(value) {
    return parseInt(value * 1000000, 10);
  },

  /**
   * 币出库缩小 1000000倍
   * @param {*} value
   * @returns
   */
  coinGetDB(value) {
    return parseFloat(value / 1000000);
  },

  /**
   * 通过api获取 以太坊 交易
   * @param {*} url
   * @param {*} params
   * @returns
   */
  async getEthTransactionList(url, params) {
    const queryString = this.paresQuery(params);
    try {
      const res = await this.ctx.curl(url + queryString, {
        dataType: "json",
        timeout: 5000,
      });
      if (res.status != 200) {
        return [];
      }
      if (res.data.status != 1) {
        return [];
      }
      let list = [];
      for (const transaction of res.data.result) {
        list.push(this.handleEthTransaction(transaction));
      }
      return list;
    } catch (err) {
      this.logger.error(err);
      return [];
    }
  },

  /**
   * 获取波场链交易
   * @param {*} url
   * @returns
   */
  async getTronTransactionList(url) {
    const { ctx } = this;
    try {
      const res = await this.ctx.curl(url, {
        // 自动解析 JSON response
        dataType: "json",
        // 3 秒超时
        timeout: 5000,
        headers: {
          "TRON-PRO-API-KEY": ctx.app.config.tron.apiKey,
          "Content-Type": "application/json",
        },
      });
      let list = {};
      let lastBlock = 0;
      if (res.status != 200) {
        return { list, next: "", lastBlock };
      }
      if (res.data.data.length == 0) {
        return { list, next: "", lastBlock };
      }
      for (const transaction of res.data.data) {
        const trans = this.handleTronTransaction(transaction);
        if (!trans) continue;
        if (!list[trans.to]) {
          list[trans.to] = [trans];
        } else {
          list[trans.to].push(trans);
        }
        if (trans.block_timestamp > lastBlock) {
          lastBlock = trans.block_timestamp;
        }
      }
      const next = res.data.meta.links ? res.data.meta.links.next : "";
      return { list, next, lastBlock };
    } catch (err) {
      this.logger.error(err);
      return { list: {}, next: "", lastBlock: 0 };
    }
  },

  /**
   * 波场链交易处理
   * @param {*} transaction
   * @returns
   */
  handleTronTransaction(transaction) {
    if (transaction.ret[0].contractRet != "SUCCESS") {
      return false;
    }
    let data = transaction.raw_data.contract[0].parameter.value.data;
    if (!data) {
      return false;
    }
    //特殊处理，不然报错，41 => 0x
    data = data.replace("0000041", "0000000");
    const dataArr = this.decodeTronParams(["address", "uint256"], data, true);
    return {
      txid: transaction.txID,
      block_timestamp: transaction.block_timestamp,
      block_number: transaction.blockNumber,
      from: TronWeb.address.fromHex(
        transaction.raw_data.contract[0].parameter.value.owner_address
      ),
      to: TronWeb.address.fromHex(dataArr[0]),
      value: dataArr[1].toString(),
    };
  },

  /**
   * 以太坊交易处理
   * @param {*} transaction
   * @returns
   */
  handleEthTransaction(transaction) {
    transaction["gasToEth"] = new Decimal(transaction.gas)
      .mul(transaction.gasPrice)
      .div(10 ** 18)
      .toFixed();
    if (transaction.tokenSymbol !== undefined) {
      transaction["value"] = new Decimal(transaction.value)
        .div(10 ** transaction.tokenDecimal)
        .toFixed();
    } else {
      transaction.value = new Decimal(transaction.value)
        .div(10 ** 18)
        .toFixed();
    }
    return transaction;
  },

  paresQuery(obj) {
    let queryString = "?";
    Object.keys(obj).map((key) => {
      queryString += key + "=" + obj[key] + "&";
    });
    return queryString;
  },

  //types:参数类型列表，如果函数有多个返回值，列表中类型的顺序应该符合定义的顺序
  //output: 解码前的数据
  //ignoreMethodHash：对函数返回值解码，ignoreMethodHash填写false，如果是对gettransactionbyid结果中的data字段解码时，ignoreMethodHash填写true
  decodeTronParams(types, output, ignoreMethodHash) {
    if (!output || typeof output === "boolean") {
      ignoreMethodHash = output;
      output = types;
    }

    if (ignoreMethodHash && output.replace(/^0x/, "").length % 64 === 8)
      output = "0x" + output.replace(/^0x/, "").substring(8);

    const abiCoder = new AbiCoder();

    if (output.replace(/^0x/, "").length % 64)
      throw new Error(
        "The encoded string is not valid. Its length must be a multiple of 64."
      );
    return abiCoder.decode(types, output).reduce((obj, arg, index) => {
      if (types[index] == "address")
        arg = ADDRESS_PREFIX + arg.substr(2).toLowerCase();
      obj.push(arg);
      return obj;
    }, []);
  },
};

"use strict";

const TronWeb = require("tronweb");

class Tron {
  tronWeb;

  constructor(api = "", apiKey = "", privateKey = "") {
    this.tronWeb = new TronWeb({
      fullHost: api,
      headers: { "TRON-PRO-API-KEY": apiKey },
      privateKey,
    });
  }

  /**
   * 发送代币交易
   * @param contract
   * @param from
   * @param to
   * @returns
   */
  async sendTokenTransaction(contract, from, to) {
    //加载合约
    const contractInstance = await this.tronWeb.contract().at(contract);
    //获取当前账户的token余额
    const balance = await contractInstance.balanceOf(from).call();
    if (balance <= 0) {
      return 0;
    }
    // //获取当前账号的trx是否够
    // const trx = await this.tronWeb.trx.getBalance(from);
    // //预估当前交易的手续费
    // const gas = this.calcTronGas(contract, from, to, balance);
    // if (gas < trx) {
    //   //余额不足
    //   return;
    // }
    //发送交易
    const txid = await contractInstance.transfer(to, balance).send({
      feeLimit: 1000000,
    });
    return txid;
  }

  /**
   * 获取tron转移所需手续费，from地址必须激活，否则无法计算
   * @param {string} from
   * @param {string} to
   * @param {(number| string)} value
   * @returns
   */
  async calcTronGas(contract, from, to, amount) {
    const parameter1 = [
      { type: "address", value: to },
      { type: "uint256", value: amount },
    ];
    const transaction =
      await this.tronWeb.transactionBuilder.triggerConstantContract(
        contract,
        "transfer(address,uint256)",
        {},
        parameter1,
        from
      );
    return transaction.energy_used;
  }

  /**
   * 获取交易的详细信息，包括交易的费用和交易事件
   * @param {*} txid
   * @returns
   */
  async getTransactionInfo(txid) {
    return await this.tronWeb.trx.getTransactionInfo(txid);
  }

  /**
   * 获取当前账户的trx
   * @param {*} address
   * @returns
   */
  async getBalance(address) {
    return await this.tronWeb.trx.getBalance(address);
  }

  /**
   * 发送 trx 交易
   * @param {*} to
   * @param {*} amount
   * @returns
   */
  async sendTransaction(to, amount) {
    return await this.tronWeb.trx.sendTransaction(to, amount);
  }

  /**
   *返回最新区块信息
   */
  async getCurrentBlock() {
    const block = await this.tronWeb.trx.getCurrentBlock();
    return block.block_header.raw_data.timestamp;
  }
}

module.exports = Tron;

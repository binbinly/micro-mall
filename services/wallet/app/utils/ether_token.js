"use strict";

const fs = require("fs");
const Ether = require("./ether");
const { Decimal } = require("decimal.js");

class EtherToken extends Ether {
  contractInstance;

  /**
   * 合约初始化
   * @memberof TokenModel
   */
  contractInit(path, address, name = "token") {
    const abiStr = fs.readFileSync(path, "utf8");
    const abi = JSON.parse(abiStr);
    this.contractInstance = new this.web3.eth.Contract(abi[name], address);
  }

  /**
   * 获取代币余额
   * @param {string} address
   * @returns
   */
  async getTokenAmount(address, decimal = 18) {
    const tokenAmount = await this.contractInstance.methods
      .balanceOf(address)
      .call();
    return new Decimal(tokenAmount).div(10 ** decimal).toString();
  }

  /**
   * 构建转账参数
   * @param {string} to
   * @param {(string|number)} tokenAmount
   * @returns
   */
  async buildTransactionAbiData(to, tokenAmount, decimal = 18) {
    tokenAmount = new Decimal(tokenAmount).mul(10 ** decimal).toString();
    const abiData = await this.contractInstance.methods
      .transfer(to, tokenAmount)
      .encodeABI();
    return abiData;
  }

  /**
   * 计算发送代币所需的eth
   * @param {string} from
   * @param {string} to
   * @param {(string|number)} tokenAmount
   * @returns
   * @memberof TokenModel
   */
  async calcTokenGas(from, to, tokenAmount, decimal = 18) {
    const gasPrice = await this.getGasPrice();
    tokenAmount = new Decimal(tokenAmount).mul(10 ** decimal).toString();
    const gasLimit = await this.contractInstance.methods
      .transfer(to, tokenAmount)
      .estimateGas({ from });
    return {
      gasPrice,
      gasLimit,
      gasToEth: this.web3.utils.fromWei(
        new Decimal(gasPrice).mul(gasLimit).toFixed(),
        "ether"
      ),
    };
  }

  /**
   * 构建代币转移TxObj
   * @param {string} contractAddress
   * @param {string} from
   * @param {string} to
   * @param {(string| number)} [amount='all']
   * @returns
   */
  async buildTokenTransaction(contractAddress, from, to, amount = "") {
    let data;
    if (amount == "") {
      //默认发送全部币
      amount = await this.getTokenAmount(from);
      data = await this.buildTransactionAbiData(to, amount);
    } else {
      let balance = await this.contractInstance.methods.balanceOf(from);
      balance = this.web3.utils.fromWei(balance, "ether");
      if (balance < amount) {
        throw new Error("余额不足");
      }
      data = await this.buildTransactionAbiData(to, amount);
    }
    const gasObj = await this.calcTokenGas(from, to, amount);
    const nonce = await this.getAddressNonce(from);

    return {
      from,
      to: contractAddress,
      nonce,
      gasPrice: gasObj.gasPrice,
      gasLimit: gasObj.gasLimit,
      data,
    };
  }
}

module.exports = EtherToken;

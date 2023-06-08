"use strict";

const Web3 = require("web3");
const Tx = require("ethereumjs-tx");
const { Decimal } = require("decimal.js");

class Ether {
  web3;

  erc20Method = "0xa9059cbb";

  erc20TransferABI = [
    {
      type: "address",
      name: "receiver",
    },
    {
      type: "uint256",
      name: "amount",
    },
  ];

  constructor(provider = "") {
    this.web3 = new Web3(provider);
  }

  /**
   * 获取eth 余额
   * @param {string} address
   * @returns
   */
  async getBalance(address) {
    const amount = await this.web3.eth.getBalance(address);
    return this.web3.utils.fromWei(amount, "ether");
  }

  /**
   * 获取转移所需手续费
   * @param {string} from
   * @param {string} to
   * @param {(number| string)} value
   * @returns
   */
  async calcEthGas(from, to, value) {
    const gasPrice = await this.getGasPrice();
    const gasLimit = await this.web3.eth.estimateGas({
      gasPrice,
      value: this.web3.utils.toWei(value.toString(), "ether"),
    });

    return {
      gasLimit,
      gasPrice,
      gasToEth: this.web3.utils.fromWei(
        new Decimal(gasPrice).mul(gasLimit).toFixed()
      ),
    };
  }

  /**
   * 计算最多可发送eth数量
   * @param {string} from
   * @param {string} to
   * @returns
   */
  async calcMaxSendEthAmount(from, to) {
    const gasObj = await this.calcEthGas(from, to, 0);
    const ethAmount = await this.getEthAmount(from);
    return new Decimal(ethAmount).mul(gasObj.gasToEth).toString();
  }

  /**
   * 获取 gasPrice
   * @param amount
   * @param def
   * @returns
   */
  async getGasPrice(amount = 0, def = false) {
    if (amount > 0) {
      return this.web3.utils.toWei(amount.toString(), "gwei");
    }
    if (def) {
      return this.web3.utils.toWei(def, "gwei");
    } else {
      return await this.web3.eth.getGasPrice();
    }
  }

  /**
   * 获取当前地址的 nonce 值，即当前已经交易的笔数
   * @param address
   * @returns
   */
  async getAddressNonce(address) {
    return await this.web3.eth.getTransactionCount(address);
  }

  /**
   * 获取链id
   * @returns
   */
  async getChainId() {
    return await this.web3.eth.getChainId();
  }

  /**
   * 获取区块高度
   * @returns
   */
  async getBlockNumber() {
    return await this.web3.eth.getBlockNumber();
  }

  /**
   * 根据交易哈希返回交易收据。
   * @param hash
   * @returns
   */
  async getTransactionReceipt(hash) {
    return await this.web3.eth.getTransactionReceipt(hash);
  }

  /**
   * 计算预期消耗的 gas
   * @param gasPrice
   * @param gasLimit
   * @returns
   */
  calcGas(gasPrice, gasLimit) {
    return this.web3.utils.fromWei(
      new Decimal(gasPrice).mul(gasLimit).toFixed(),
      "ether"
    );
  }

  /**
   * 构建Eth 转移
   * @param {string} from
   * @param {string} to
   * @param {(string|number)} [amount]
   */
  async buildEthTransaction(from, to, amount = "") {
    let gasObj = { gasLimit: 0, gasPrice: 0, gasToEth: 0 };

    if (amount === "") {
      amount = await this.getEthAmount(from);
      gasObj = await this.calcEthGas(from, to, amount);
      amount = new Decimal(amount).sub(gasObj.gasToEth).toString();
    } else {
      gasObj = await this.calcEthGas(from, to, amount);
    }

    return {
      from,
      to,
      gasPrice: gasObj.gasPrice,
      gasLimit: gasObj.gasLimit,
      value: this.web3.utils.toWei(amount.toString(), "ether"),
    };
  }

  /**
   * 签名交易
   * @param {object} transaction
   * @param {string} privateKey
   * @returns
   */
  signTransaction(transaction, privateKey) {
    Object.keys(transaction).map((key) => {
      if (["to", "data"].includes(key) === false) {
        transaction[key] = this.web3.utils.toHex(transaction[key]);
      }
    });
    const bufferKey = Buffer.from(privateKey.replace("0x", ""), "hex");
    const tx = new Tx(transaction);
    tx.sign(bufferKey);
    const serializedTx = tx.serialize();
    return `0x${serializedTx.toString("hex")}`;
  }

  /**
   * 发送交易
   * @param {string} Tx
   * @returns
   */
  sendTransaction(Tx) {
    return this.web3.eth.sendSignedTransaction(Tx);
  }

  /**
   * 地址校验
   * @param address
   * @returns
   */
  checkAddress(address) {
    address = this.web3.utils.toChecksumAddress(address);
    this.web3.utils.checkAddressChecksum(address);
    return address;
  }

  /**
   * 解析erc20交易input数据
   * @param {*} tx
   * @returns
   */
  decodeERC20TransferParameters(tx) {
    if (tx.methodId != this.erc20Method) {
      return false;
    }
    const decoded = this.web3.eth.abi.decodeParameters(
      this.erc20TransferABI,
      tx.input.slice(10)
    );
    return { to: decoded.receiver, amount: decoded.amount };
  }
}

module.exports = Ether;

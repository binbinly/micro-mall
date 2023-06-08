"use strict";

const Ether = require("./ether");
const fs = require("fs");

class BscPancake extends Ether {
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
   * pancake 中获取当前token 的币价
   * @returns
   */
  async getTokenPrice(token1, token2) {
    const unit = this.web3.utils.toWei("1", "ether");
    const amount = await this.contractInstance.methods
      .getAmountsOut(unit, [token1, token2])
      .call();
    return this.web3.utils.fromWei(amount[1], "ether");
  }
}

module.exports = BscPancake;

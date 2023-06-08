"use strict";

const Service = require("egg").Service;

class ConfigService extends Service {
  table = "config";
  ether_wallet_key = "ether_gas_wallet";
  tron_wallet_key = "tron_gas_wallet";
  ether_receive_key = "ether_receive_address";
  tron_receive_key = "tron_receive_address";

  /**
   * 详情
   * @param {*} address
   * @returns
   */
  async detail(key) {
    return await this.app.mysql.get(this.table, { key });
  }

  /**
   * value值
   * @param {*} key
   * @returns
   */
  async value(key) {
    const info = await this.app.mysql.get(this.table, { key });
    if (!info) {
      return "";
    }
    return info.value;
  }

  /**
   * 获取发送 gas 的钱包
   * @param {*} key
   * @returns
   */
  async getWallet(key) {
    const gasWallet = await this.value(key);
    if (!gasWallet) {
      return false;
    }
    const wallet = JSON.parse(gasWallet);
    return { address: wallet.address, key: this.app.aesEcbDecrypt(wallet.key) };
  }

  /**
   * 创建以太坊官方钱包
   * @returns
   */
  async createEth() {
    const account = await this.app.mysql.get(this.table, {
      key: this.ether_wallet_key,
    });
    if (account) {
      return true;
    }
    const wallet = this.ctx.helper.createEtherWallet();

    const result = await this.app.mysql.insert(this.table, {
      key: this.ether_wallet_key,
      value: JSON.stringify(wallet),
    });
    if (result.affectedRows === 1) {
      return true;
    }
    return false;
  }

  /**
   * 创建波场链官方钱包
   * @returns
   */
  async createTron() {
    const account = await this.app.mysql.get(this.table, {
      key: this.tron_wallet_key,
    });
    if (account) {
      return true;
    }
    const wallet = await this.ctx.helper.createTronWallet();

    //保存入库
    const result = await this.app.mysql.insert(this.table, {
      key: this.tron_wallet_key,
      value: JSON.stringify(wallet),
    });
    if (result.affectedRows === 1) {
      return true;
    }
    return false;
  }
}

module.exports = ConfigService;

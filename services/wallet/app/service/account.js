"use strict";

const Service = require("egg").Service;
const moment = require("moment");

class AccountService extends Service {
  table = "chain_account";

  /**
   * 账号详情
   * @param {*} address
   * @returns
   */
  async detail(address) {
    return await this.app.mysql.get(this.table, { address });
  }

  /**
   * 创建以太坊钱包
   * @param {*} user_id
   * @returns
   */
  async createEth(user_id) {
    const account = await this.app.mysql.get(this.table, {
      user_id,
      type: 1,
    });
    if (account) {
      return account.address;
    }
    const { address, key } = this.ctx.helper.createEtherWallet();

    //保存入库
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    const result = await this.app.mysql.insert(this.table, {
      user_id,
      address,
      key,
      type: 1,
      created_at: now,
      updated_at: now,
    });
    if (result.affectedRows === 1) {
      return address;
    }
    return false;
  }

  /**
   * 创建波场链钱包
   * @param {*} user_id
   * @returns
   */
  async createTron(user_id) {
    const account = await this.app.mysql.get(this.table, {
      user_id,
      type: 2,
    });
    if (account) {
      return account.address;
    }
    const { address, key } = await this.ctx.helper.createTronWallet();
    //保存入库
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    const result = await this.app.mysql.insert(this.table, {
      user_id,
      address,
      key,
      type: 2,
      created_at: now,
      updated_at: now,
    });
    if (result.affectedRows === 1) {
      return address;
    }
    return false;
  }
}

module.exports = AccountService;

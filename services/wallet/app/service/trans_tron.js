"use strict";

const Service = require("egg").Service;
const moment = require("moment");
const Tron = require("../utils/tron");

class TransTronService extends Service {
  table = "chain_trans_tron";

  /**
   * api获取波场代币交易记录
   * @param {*} api
   * @param {*} contractaddress
   * @returns
   */
  async getTransListByApi(api, contractaddress, token) {
    const { ctx, app, service } = this;
    const db = ctx.app.mysql;
    const key = "tron_" + contractaddress;
    const info = await service.config.detail(key);
    if (!info) {
      const tron = new Tron(app.config.tron.api, app.config.tron.apiKey);
      const value = await tron.getCurrentBlock();
      await db.insert(service.config.table, { key, value });
      return;
    }

    const startblock = info.value;
    let url =
      api +
      "/v1/contracts/" +
      contractaddress +
      "/transactions?min_block_timestamp=" +
      startblock +
      "&order_by=block_timestamp,asc&limit=100";
    let endBlock = startblock;
    while (true) {
      try {
        let { list, next, lastBlock } = await ctx.helper.getTronTransactionList(
          url
        );
        if (lastBlock > endBlock) {
          endBlock = lastBlock;
        }
        if (lastBlock == 0) {
          break;
        }

        await this.match(list, contractaddress, token);
        if (next == "") {
          break;
        }
        url = next;
        await db.update(service.config.table, { id: info.id, value: endBlock });
        await app.sleep(1000);
      } catch (e) {
        ctx.logger.error("getTransListByApi tron_usdt err:%s", e);
        break;
      }
    }
  }

  /**
   * 检查交易
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
    const tron = new Tron(app.config.tron.api, app.config.tron.apiKey);
    list.forEach(async (item) => {
      //获取交易详情
      const receipt = await tron.getTransactionInfo(item.txid);
      if (receipt.receipt && receipt.receipt.result == "SUCCESS") {
        //交易成功
        await this.submit(item);
      }
    });
  }

  /**
   * 匹配平台钱包地址的所有交易并保存入库
   * @param {*} list
   * @param {*} contractaddress
   * @param {*} token
   * @returns
   */
  async match(list, contractaddress, token) {
    const { app } = this;
    const addresses = Object.keys(list);
    if (addresses.length == 0) {
      return;
    }
    //检查出平台的钱包地址
    const results = await app.mysql.query(
      "select address from chain_account where type=2 and address in (?)",
      [addresses]
    );
    if (results.length == 0) {
      return;
    }
    const now = moment(Date.now()).format("YYYY-MM-DD HH:mm:ss");
    results.forEach(async (val) => {
      //匹配交易
      let transList = list[val.address];
      if (transList) {
        transList.forEach(async (trans) => {
          if (await app.mysql.get(this.table, { txid: trans.txid })) {
            return;
          }
          //保存交易
          await app.mysql.insert(this.table, {
            from: trans.from,
            to: trans.to,
            value: trans.value,
            txid: trans.txid,
            contract: contractaddress,
            token,
            block_timestamp: trans.block_timestamp,
            block_number: trans.block_number,
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
   * @returns
   */
  async submit(item) {
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
        2,
        conn
      );
      //修改状态
      await conn.update(this.table, {
        id: item.id,
        is_confirm: 1,
      });
      await conn.commit(); // 提交事务
    } catch (err) {
      await conn.rollback(); // 一定记得捕获异常后回滚事务！！
      throw err;
    }
  }
}

module.exports = TransTronService;

//获取币安链上的usdt合约的所有交易
module.exports = {
  schedule: {
    interval: "12s", // 1 分钟间隔
    type: "worker", // 指定所有的 worker 都需要执行
    immediate: true,
    disable: false,
  },
  async task(ctx) {
    //加锁5分钟
    const lock_key = "chain_bsc_usdt_job_lock";
    if (await ctx.app.redis.set(lock_key, 1, "EX", 300, "NX")) {
      await ctx.service.trans.getTxListByApi(
        ctx.app.config.bsc.chainId,
        ctx.app.config.bsc.api,
        ctx.app.config.bsc.apiKey,
        ctx.app.config.bsc.usdt,
        "USDT"
      );
      //释放锁
      await ctx.app.redis.del(lock_key);
    }
    await ctx.service.trans.check();
  },
};

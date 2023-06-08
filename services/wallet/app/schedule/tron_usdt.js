//波场usdt交易记录获取
module.exports = {
  schedule: {
    interval: "15s", // 1 分钟间隔
    type: "worker", // 指定所有的 worker 都需要执行
    immediate: true,
    disable: false,
  },
  async task(ctx) {
    //加锁5分钟
    const lock_key = "chain_tron_usdt_job_lock";
    if (await ctx.app.redis.set(lock_key, 1, "EX", 300, "NX")) {
      await ctx.service.transTron.getTransListByApi(
        ctx.app.config.tron.api,
        ctx.app.config.tron.usdt,
        "USDT"
      );
      //释放锁
      await ctx.app.redis.del(lock_key);
    }
    await ctx.service.transTron.check();
  },
};

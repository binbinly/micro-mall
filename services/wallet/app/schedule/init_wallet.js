//初始化
module.exports = {
  schedule: {
    interval: "1h", // 1 分钟间隔
    type: "worker", // 指定所有的 worker 都需要执行
    disable: true,
  },
  async task(ctx) {
    await ctx.service.config.createEth();
    await ctx.service.config.createTron();
    //释放锁
    await ctx.app.redis.del("chain_bsc_usdt_job_lock");
    await ctx.app.redis.del("chain_tron_usdt_job_lock");
  },
};

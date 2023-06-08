module.exports = {
  schedule: {
    interval: "1h", // 1 分钟间隔
    type: "worker", // 指定所有的 worker 都需要执行
    immediate: true,
    disable: true,
  },
  async task(ctx) {
    await ctx.service.balance.sendGasTransaction();

    await ctx.app.sleep(30000);

    await ctx.service.balance.sendTransaction();
  },
};

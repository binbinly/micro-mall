const TronWeb = require("tronweb");

const tronWeb = new TronWeb({
    fullHost: "https://api.trongrid.io",
    headers: { "TRON-PRO-API-KEY": "708d8d51-1470-40de-b23d-168b1c308747" },
  });

class AppBootHook {
    constructor(app) {
      this.app = app;
    }
  
    //应用启动完成
    async serverDidReady() {
        console.log('server run finish')
        const ctx = this.app.createAnonymousContext();
        // Address B transfers 10 USDT from address A to C: B calls transferFrom (A, C, 10)
        const trc20ContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t";//contract address
        try {
            let contract = await tronWeb.contract().at(trc20ContractAddress);
            await contract && contract.Transfer().watch((err, event) => {
                if(err)
                return console.error('Error with "Message" event:', err);
                if(event.result.value < 1000000) {
                    return
                }
                const data = {
                    from:TronWeb.address.fromHex(event.result.from),
                    to:TronWeb.address.fromHex(event.result.to),
                    value:event.result.value,
                    time:event.timestamp/1000,
                    txid:event.transaction
                }
                // console.log('event', data)
                if(data.value > 0) {
                    ctx.service.transTron.create(data)
                }
            });
        } catch(error) {
            console.error("trigger smart contract error",error)
        }
    }
  
    //应用即将关闭
    async beforeClose() {
      // Do some thing before app close.
    }
  }
  
  module.exports = AppBootHook;
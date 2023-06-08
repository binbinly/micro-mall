const Web3 = require("web3");
const fs = require("fs");
const Tx = require("ethereumjs-tx");
const { ethers } = require("ethers");
module.exports = {
  schedule: {
    interval: "30s", // 1 分钟间隔
    type: "worker", // 指定所有的 worker 都需要执行
    immediate: true,
    disable: false,
  },
  async task(ctx) {
    const CAKE = "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82";
    const USDT = "0x55d398326f99059ff775485246999027b3197955";
    const wallets = initWallet(10);
    const paths = [
      [USDT, CAKE],
      [CAKE, USDT],
    ];
    const wallet = wallets[Math.floor(Math.random() * wallets.length)];
    const path = paths[Math.floor(Math.random() * paths.length)];
    console.log("address", wallet.address, path[0]);

    const web3 = new Web3(ctx.app.config.bsc.web3Url);
    const gasPrice = await web3.eth.getGasPrice();
    const etherInst = contractInit(web3, ctx.app.config.contractPath, path[0]);
    console.log("etherInst", wallet.address, ctx.app.config.pancakeRouter);
    const allowance = await etherInst.methods
      .allowance(wallet.address, ctx.app.config.pancakeRouter)
      .call();
    console.log("allowance", allowance);
    const balance = await etherInst.methods.balanceOf(wallet.address).call();
    if (balance == 0) {
      console.log("end");
      return;
    }
    const rate = Math.ceil(Math.random() * 5);
    const amount = parseInt((balance * rate) / 10);
    console.log("amount: " + amount);
    if (allowance == 0) {
      //授权
      await sendAllowance(
        web3,
        etherInst,
        ctx.app.config.pancakeRouter,
        path[0],
        wallet,
        gasPrice
      );
      await ctx.app.sleep(5000);
    }

    const pancakeInst = contractInit(
      web3,
      ctx.app.config.contractPath,
      ctx.app.config.pancakeRouter,
      "pancake"
    );

    let timestamp = new Date().getTime();
    timestamp = parseInt(timestamp / 1000) + 1800;
    let data = "";
    if (path[0] == USDT) {
      //buy
      data = pancakeInst.methods
        .swapExactTokensForTokens(
          amount + "",
          "0",
          path,
          wallet.address,
          timestamp
        )
        .encodeABI();
    } else {
      //sell
      data = pancakeInst.methods
        .swapExactTokensForTokensSupportingFeeOnTransferTokens(
          amount + "",
          "0",
          path,
          wallet.address,
          timestamp
        )
        .encodeABI();
    }
    const gasLimit = "210000";
    const txObj = {
      from: wallet.address,
      to: ctx.app.config.pancakeRouter,
      data,
      gasLimit,
      gasPrice,
      nonce: await web3.eth.getTransactionCount(wallet.address),
    };
    // console.log("txObj: ", txObj);
    //签名交易
    const txData = signTransaction(web3, txObj, wallet.key);
    // return;
    //发送交易
    web3.eth
      .sendSignedTransaction(txData)
      .on("transactionHash", async (hash) => {
        console.log("hash", hash);
      })
      .on("receipt", async (receipt) => {
        // console.log("receipt", receipt);
        console.log("success");
      })
      .on("error", async (error) => {
        console.log("error", error);
      });
  },
};

const contractInit = function (web3, path, address, name = "token") {
  const abiStr = fs.readFileSync(path, "utf8");
  const abi = JSON.parse(abiStr);
  return new web3.eth.Contract(abi[name], address);
};

const sendAllowance = async (web3, inst, to, token, wallet, gasPrice) => {
  const data = inst.methods
    .approve(
      to,
      "115792089237316195423570985008687907853269984665640564039457584007913129639935"
    )
    .encodeABI();
  const gasLimit = "210000";
  const txObj = {
    from: wallet.address,
    to: token,
    data,
    gasLimit,
    gasPrice,
    nonce: await web3.eth.getTransactionCount(wallet.address),
  };
  //签名交易
  const txData = signTransaction(web3, txObj, wallet.key);
  //发送交易
  return await web3.eth.sendSignedTransaction(txData);
};

const signTransaction = (web3, transaction, privateKey) => {
  Object.keys(transaction).map((key) => {
    if (["to", "data"].includes(key) === false) {
      transaction[key] = web3.utils.toHex(transaction[key]);
    }
  });
  const bufferKey = Buffer.from(privateKey.replace("0x", ""), "hex");
  const tx = new Tx(transaction);
  tx.sign(bufferKey);
  const serializedTx = tx.serialize();
  return `0x${serializedTx.toString("hex")}`;
};

const initWallet = (n = 10) => {
  const filename = "wallet.json";
  try {
    const data = fs.readFileSync(filename);
    return JSON.parse(data.toString());
  } catch (err) {
    if (err.code === "ENOENT") {
      let list = [];
      for (let i = 0; i < n; i++) {
        const privateKey = ethers.utils.randomBytes(32);
        const wallet = new ethers.Wallet(privateKey);
        const key = ethers.BigNumber.from(privateKey)._hex;
        list.push({ address: wallet.address, key });
      }
      fs.writeFile(filename, JSON.stringify(list), (err) => {
        if (err) {
          console.error(err);
          return;
        }
      });
      return list;
    } else {
      throw err;
    }
  }
};

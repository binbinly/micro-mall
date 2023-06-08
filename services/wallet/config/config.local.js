const Web3 = require("web3");
const path = require("path");

module.exports = (appInfo) => {
  return {
    mysql: {
      // 单数据库信息配置
      client: {
        // host
        host: "127.0.0.1",
        // 端口号
        port: "3306",
        // 用户名
        user: "root",
        // 密码
        password: "root",
        // 数据库名
        database: "mall_ums",
      },
      // 是否加载到 app 上，默认开启
      app: true,
      // 是否加载到 agent 上，默认关闭
      agent: false,
    },
    redis: {
      client: {
        port: 6379, // Redis port
        host: "127.0.0.1", // Redis host
        password: "",
        db: 0,
      },
    },
    tron: {
      api: "https://nile.trongrid.io",
      apiKey: "18c19d1b-edac-4fae-80f8-411e74d8afb8",
      usdt: "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
      gasAddress: "THzTkg7ydSH4sF1TRM6ZhbVovVDc9ptrtd",
    },
    bsc: {
      chainId: 97,
      api: "https://api-testnet.bscscan.com/api",
      apiKey: "GQEIXEWICMNP4ECIVS1BQFC2S4N6YHWRAB",
      web3Url: "https://data-seed-prebsc-2-s2.binance.org:8545",
      usdt: "0x7ef95a0FEE0Dd31b22626fA2e10Ee6A223F8a684",
    },
    eth: {
      web3Url: "https://goerli.infura.io/v3/169bb315d1f14cb78d20a93ed667c786",
      apiKey: "17Y4MCAKUPWASA5ZA7STPZ8JKB4HHEWB7T",
      api: "https://api-goerli.etherscan.io/api",
      usdt: "0x509Ee0d083DdF8AC028f2a56731412edD63223B9",
    },
    minAmount: 1000000,
    contractPath: path.join(appInfo.baseDir, "contracts/wallet.json"),
    gasPrice: "10", //gwei
    pancakeRouter: "0x9ac64cc6e4415144c455bd8e4837fea55603e5c3",
  };
};

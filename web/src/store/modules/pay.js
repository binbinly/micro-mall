import Payment from '@/contracts/Payment.json'
import Web3 from '@/utils/web3'
import { orderPayNotify } from '@/api/order'
export default {
  state: {
    web3: null,
    networkId: 0,
    account: null
  },
  mutations: {
    initWeb3(state) {
      state.web3 = Web3()
      state.web3.eth.getAccounts().then(accounts => {
        state.account = accounts[0]
        console.log('accounts', accounts)
      })
      state.web3.eth.net.getId().then(networkId => {
        state.networkId = networkId
        console.log('networkId', networkId)
      })
    }
  },
  actions: {
    //创建支付合约
    createContract({ state }) {
      return new Promise((result, rej) => {
        new state.web3.eth.Contract(Payment.abi)
          .deploy({ data: Payment.bytecode })
          .send({ from: state.account })
          .then(res => {
            result(res)
          })
          .catch(err => {
            console.error('err', err)
            rej(err)
          })
      })
    },
    //获取地址余额
    getBalance({ state }, address) {
      return new Promise((result, rej) => {
        state.web3.eth
          .getBalance(address)
          .then(balance => {
            result(state.web3.utils.fromWei(balance, 'ether'))
          })
          .catch(err => {
            console.error('err', err)
            rej(err)
          })
      })
    },
    doPay({ state }, { address, amount, order_no, trade_no, eth }) {
      const value = state.web3.utils.toWei(eth.toString(), 'ether')
      const contract = new state.web3.eth.Contract(Payment.abi, address)
      return new Promise((result, rej) => {
        contract.methods
          .pay(trade_no, order_no, amount)
          .send({ from: state.account, value, gasLimit: 210000 })
          .then(res => {
            result(res)
          })
          .catch(err => {
            console.error('pay err', err)
            rej(err)
          })
      })
    },
    //提款
    doWithdraw({ state }, { address, amount }) {
      const value = state.web3.utils.toWei(amount, 'ether')
      const contract = new state.web3.eth.Contract(Payment.abi, address)
      return new Promise((result, rej) => {
        contract.methods
          .withdraw(value)
          .send({ from: state.account })
          .then(res => {
            result(res)
          })
          .catch(err => {
            console.error('withdraw err', err)
            rej(err)
          })
      })
    },
    //交易笔数
    tradeCount({ state }, address) {
      const contract = new state.web3.eth.Contract(Payment.abi, address)
      return new Promise((result, rej) => {
        contract.methods
          .tradesOf(state.account)
          .call()
          .then(res => {
            result(res)
          })
          .catch(err => {
            console.error('tradesOf err', err)
            rej(err)
          })
      })
    },
    //交易记录
    tradeLog({ state }, { address, offset, limit }) {
      const contract = new state.web3.eth.Contract(Payment.abi, address)

      Promise.all(
        Array(limit)
          .fill()
          .map(
            (el, index) =>
              new Promise((result, rej) => {
                contract.methods
                  .tradesOfAt(state.account, index + offset)
                  .call()
                  .then(trade => {
                    const amount = state.web3.utils.fromWei(trade[2], 'ether')
                    result({
                      id: trade[0],
                      order_no: trade[1],
                      amount,
                      date: trade[3]
                    })
                  })
                  .catch(err => {
                    console.error('tradesOfAt err', index, err)
                    rej(err)
                  })
              })
          )
      )
    },
    onNotify({ state }, { address, order_no, trade_no, eth, pay_type }) {
      const contract = new state.web3.eth.Contract(Payment.abi, address)
      contract.events
        .Notify(
          {
            filter: {}
          },
          function (error, event) {
            console.log('notify event', error, event)
          }
        )
        .on('data', function (event) {
          console.log('notify data', event)
          orderPayNotify({ trans_hash: event.transactionHash, order_no, trade_no, pay_amount: eth * 100, pay_type })
        })
        .on('changed', function (event) {
          // remove event from local database
          console.log('changed', event)
        })
        .on('error', console.error)
    }
  }
}

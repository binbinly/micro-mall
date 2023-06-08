import { default as Web3 } from 'web3'

export default () => {
  let web3

  if (window && window.ethereum) {
    web3 = new Web3(window.ethereum)
    window.ethereum.enable()
    // web3.setProvider(new Web3.providers.WebsocketProvider('ws://127.0.0.1:8546'))
  } else if (window.web3) {
    web3 = new Web3(window.web3.currentProvider)
  } else {
    console.warn(
      'No web3 detected. Falling back to http://127.0.0.1:8545. You should remove this fallback when you deploy live'
    )
    // fallback - use your fallback strategy (local node / hosted node + in-dapp id mgmt / fail)
    web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'))
  }
  console.log('web3', web3)
  return web3
}

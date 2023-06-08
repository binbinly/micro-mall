const { ethers } = require('ethers');

const provider = new ethers.JsonRpcProvider("https://goerli.infura.io/v3/169bb315d1f14cb78d20a93ed667c786");
// 合约地址
const contractAddress = '0x509Ee0d083DdF8AC028f2a56731412edD63223B9'
// 构建ABI
const abi = [
  "function decimals() view returns (uint8)",
  "function symbol() view returns (string)",
  "function balanceOf(address a) view returns (uint)"
]
// 构建合约对象
const contract = new ethers.Contract(contractAddress, abi, provider);
// console.log('contract', contract)

run = async () => {
  // Begin listening for any Transfer event
  console.log('start', await contract.symbol(), await contract.decimals())
  // The symbol name for the token
}
run()
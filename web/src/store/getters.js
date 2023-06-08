const getters = {
  token: state => state.user.token,
  user: state => state.user.user,
  account: state => state.pay.account,
  // 获取默认地址
  defaultAddress: state => {
    return state.address.list.filter(v => v.is_default)
  },
  // 购物车是否为空
  cartIsEmpty: state => {
    return state.cart.list.length === 0
  },
  // 购物车商品数量
  cartCount: state => {
    if (state.cart.list.length <= 99) {
      return state.cart.list.length || ''
    }
    return '99+'
  },
  //以太坊网络
  ethNetwork: state => {
    switch (state.pay.networkId) {
      case 1:
        return '以太坊 Ethereum 主网络'
      case 3:
        return 'Ropsten 测试网络'
      case 42:
        return 'Kovan 测试网络'
      case 4:
        return 'Rinkeby 测试网络'
      case 5:
        return 'Goerli 测试网络'
      case 15:
        return 'Localhost 8545'
      default:
        return 'Unknown'
    }
  }
}
export default getters

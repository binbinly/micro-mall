import api from './index'

//购物车列表
export function getCartList() {
  return api.get(api.Cart.List, null, true)
}

//修改购物车商品数量
export function cartEditNum(data) {
  return api.post(api.Cart.EditNum, data)
}

//修改购物车商品
export function cartEdit(data) {
  return api.post(api.Cart.Edit, data)
}

//添加购物车
export function cartAdd(data) {
  return api.post(api.Cart.Add, data)
}

//删除购物车商品
export function cartDelete(data) {
  return api.post(api.Cart.Del, data)
}

//清空购物车
export function cartEmpty() {
  return api.get(api.Cart.Empty, null, true)
}

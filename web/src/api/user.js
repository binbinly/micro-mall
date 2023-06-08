import api from './index'

//用户注销
export function userLogout() {
  return api.get(api.Member.Logout, null, true)
}

//修改用户信息
export function userEdit(data) {
  return api.post(api.Member.Edit, data)
}

//用户收货地址
export function userAddressList() {
  return api.get(api.Address.List, null, true)
}

//添加用户收货地址
export function userAddressAdd(data) {
  return api.post(api.Address.Add, data)
}

//修改用户收货地址
export function userAddressEdit(data) {
  return api.post(api.Address.Edit, data)
}

//删除用户收货地址
export function userAddressDel(id) {
  return api.get(api.Address.Del, { id }, true)
}

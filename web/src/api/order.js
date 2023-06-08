import api from './index'

//提交订单
export function submitOrder(data) {
  return api.post(api.Order.Submit, data)
}

//商品直接提交订单
export function submitGoodsOrder(data) {
  return api.post(api.Order.GoodsSubmit, data)
}

//订单详情
export function orderDetail(id) {
  return api.get(api.Order.Detail + '/' + id, null, true)
}

//订单列表
export function getOrderList(status, page) {
  return api.get(api.Order.List + '/status/' + status + '/p/' + page, null, true)
}

//退款
export function orderRefund(data) {
  return api.post(api.Order.Refund, data)
}

//确认收货
export function orderReceipt(data) {
  return api.post(api.Order.Receipt, data)
}

//支付成功通知
export function orderPayNotify(data) {
  return api.post(api.Order.Notify, data)
}

//取消订单
export function orderCancel(data) {
  return api.post(api.Order.Cancel, data)
}

//评价
export function orderComment(data) {
  return api.post(api.Order.Comment, data)
}

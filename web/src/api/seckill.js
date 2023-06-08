import api from './index'

//当下所有秒杀场次
export function getSessions() {
  return api.get(api.Seckill.Sessions)
}

//场次下商品
export function getSkus(sessionId) {
  return api.get(api.Seckill.Skus + '/' + sessionId)
}

//获取秒杀商品信息
export function getSkuInfo(skuId) {
  return api.get(api.Seckill.Detail + '/' + skuId)
}

//秒杀商品
export function submitKillOrder(data) {
  return api.post(api.Seckill.Submit, data)
}

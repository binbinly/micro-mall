import api from './index'

//优惠券列表
export function getCouponList(sku_id) {
  return api.get(api.Coupon.List + '/' + sku_id, null, true)
}

//我的优惠券
export function getMyCoupon() {
  return api.get(api.Coupon.My, null, true)
}

//优惠券领取
export function couponDraw(coupon_id) {
  return api.post(api.Coupon.Draw, { coupon_id })
}

import request from '@/utils/request'

const api = {
  SendCode: '/send_code',
  Member: {
    Login: '/member/login',
    PhoneLogin: '/member/phone_login',
    Register: '/member/reg',
    Edit: '/member/edit',
    EditPwd: '/member/edit_pwd',
    Profile: '/member/profile',
    Logout: '/member/logout'
  },
  App: {
    Home: '/app/home',
    HomeCat: '/app/home_cat',
    Notice: '/app/notice',
    Search: '/app/search',
    PayList: '/app/pay_list'
  },
  Product: {
    List: '/product/list',
    Detail: '/product/detail',
    Search: '/product/search',
    Category: '/product/cat',
    Attr: '/product/attr'
  },
  Address: {
    List: '/address/list',
    Add: '/address/add',
    Edit: '/address/edit',
    Del: '/address/del'
  },
  Cart: {
    List: '/cart/list',
    Del: '/cart/del',
    EditNum: '/cart/edit_num',
    Edit: '/cart/edit',
    Add: '/cart/add',
    Empty: '/cart/clear'
  },
  Coupon: {
    List: '/coupon/list',
    My: '/coupon/my',
    Draw: '/coupon/draw'
  },
  Seckill: {
    Sessions: '/seckill/sessions',
    Skus: '/seckill/skus',
    Detail: '/seckill/detail',
    Submit: '/seckill/submit'
  },
  Order: {
    Submit: '/order/submit',
    GoodsSubmit: '/order/sku_submit',
    Detail: '/order/detail',
    List: '/order/list',
    Notify: '/order/notify',
    Refund: '/order/refund',
    Receipt: '/order/receipt',
    Cancel: '/order/cancel',
    Comment: '/order/comment'
  },
  post(url, data, params = null, auth = true, hideLoading = false, result = false) {
    return request({
      url,
      method: 'post',
      data,
      params,
      auth,
      hideLoading,
      result
    })
  },
  get(url, params = null, auth = false, hideLoading = true, result = false) {
    return request({
      url,
      params,
      auth,
      hideLoading,
      result
    })
  }
}

export default api

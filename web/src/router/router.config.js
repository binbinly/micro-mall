/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/',
    component: () => import('@/views/layouts/index'),
    redirect: '/home',
    meta: {
      title: '首页',
      keepAlive: false
    },
    children: [
      {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/home/index'),
        meta: { title: '首页', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/class',
        name: 'Class',
        component: () => import('@/views/class/index'),
        meta: { title: '分类', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/cart',
        name: 'Cart',
        component: () => import('@/views/cart/index'),
        meta: { title: '购物车', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/my',
        name: 'My',
        component: () => import('@/views/my/index'),
        meta: { title: '我的', keepAlive: false, tabbar: true, badge: '' }
      },
      {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/user/login'),
        meta: { title: '登录', keepAlive: false }
      },
      {
        path: '/setting',
        name: 'Setting',
        component: () => import('@/views/user/setting'),
        meta: { title: '用户设置', keepAlive: true, auth: true }
      },
      {
        path: '/about',
        name: 'About',
        component: () => import('@/views/home/about'),
        meta: { title: '关于', keepAlive: true }
      },
      {
        path: '/product_list',
        name: 'ProductList',
        component: () => import('@/views/product/list'),
        meta: { title: '商品列表', keepAlive: false }
      },
      {
        path: '/product_detail',
        name: 'ProductDetail',
        component: () => import('@/views/product/detail'),
        meta: { title: '商品详情', keepAlive: false }
      },
      {
        path: '/notice',
        name: 'Notice',
        component: () => import('@/views/home/notice'),
        meta: { title: '公告', keepAlive: true }
      },
      {
        path: '/seckill',
        name: 'Seckill',
        component: () => import('@/views/home/seckill'),
        meta: { title: '秒杀', keepAlive: true }
      },
      {
        path: '/search',
        name: 'Search',
        component: () => import('@/views/home/search'),
        meta: { title: '搜索', keepAlive: true }
      },
      {
        path: '/product_comment',
        name: 'ProductComment',
        component: () => import('@/views/product/comment'),
        meta: { title: '商品评论', keepAlive: true }
      },
      {
        path: '/order_confirm',
        name: 'OrderConfirm',
        component: () => import('@/views/order/confirm'),
        meta: { title: '订单确认', keepAlive: true, auth: true }
      },
      {
        path: '/user_address',
        name: 'UserAddress',
        component: () => import('@/views/user/address'),
        meta: { title: '收货地址', keepAlive: true, auth: true }
      },
      {
        path: '/address_edit',
        name: 'AddressEdit',
        component: () => import('@/views/user/address_edit'),
        meta: { title: '收货地址修改', keepAlive: true, auth: true }
      },
      {
        path: '/order_detail',
        name: 'OrderDetail',
        component: () => import('@/views/order/detail'),
        meta: { title: '订单详情', keepAlive: true, auth: true }
      },
      {
        path: '/order',
        name: 'Order',
        component: () => import('@/views/order/list'),
        meta: { title: '我的订单', keepAlive: true, auth: true }
      },
      {
        path: '/order_pay',
        name: 'OrderPay',
        component: () => import('@/views/order/pay'),
        meta: { title: '订单支付', keepAlive: true, auth: true }
      },
      {
        path: '/user_wallet',
        name: 'UserWallet',
        component: () => import('@/views/user/wallet'),
        meta: { title: '我的钱包', keepAlive: true, auth: true }
      },
      {
        path: '/user_contract',
        name: 'UserContract',
        component: () => import('@/views/user/contract'),
        meta: { title: '我的合约', keepAlive: true, auth: true }
      },
      {
        path: '/order_refund',
        name: 'OrderRefund',
        component: () => import('@/views/order/refund'),
        meta: { title: '退款', keepAlive: true, auth: true }
      },
      {
        path: '/order_comment',
        name: 'OrderComment',
        component: () => import('@/views/order/comment'),
        meta: { title: '评价', keepAlive: true, auth: true }
      },
      {
        path: '/order_logistics',
        name: 'OrderLogistics',
        component: () => import('@/views/order/logistics'),
        meta: { title: '物流信息', keepAlive: true, auth: true }
      }
    ]
  }
]

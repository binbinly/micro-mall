<template>
  <div>
    <van-nav-bar :title="detail.title" left-arrow fixed placeholder @click-left="back" />

    <!-- 商品详情大图 -->
    <swiper-image :images="detail.banners" height="375" prediv></swiper-image>
    <!-- 基础详情 -->
    <base-info :detail="detail"></base-info>
    <van-notice-bar v-if="killInfo.id" wrapable :scrollable="false" color="#FFF" background="#ee0a24">
      <span v-if="killInfo.open">该商品正在秒杀</span>
      <template v-else>
        <span>请提前设置好默认收货地址！秒杀价：￥{{killInfo.price}}，请等待</span>
        <van-count-down :time="killTime" style="color:#FFF;" />
      </template>
    </van-notice-bar>
    <!-- 滚动商品特性 -->
    <product-attrs :baseAttrs="detail.attrs"></product-attrs>
    <!-- 属性选择 -->
    <div class="p-1">
      <div class="rounded border bg-light-secondary">
        <div @click="coupon.show=true">
          <div class="d-flex p-1">
            <span class="mr-2 text-muted">优惠券</span>
            <span class="main-text-color">马上领取</span>
          </div>
        </div>
        <div @click="serviceShow = true">
          <div class="d-flex a-center p-1">
            <div class="text-muted font d-flex a-center mr-2">
              <span class="iconfont icon-finish main-text-color"></span>
              小米自营
            </div>
            <div class="text-muted font d-flex a-center mr-2">
              <span class="iconfont icon-finish main-text-color"></span>
              小米发货
            </div>
            <div class="text-muted font d-flex a-center mr-2">
              <span class="iconfont icon-finish main-text-color"></span>
              七天无理由退货
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="d-flex a-center j-center p-1 text-primary font" hover-class="bg-light-secondary" @click="open('/product_comment')">
      查看评论 <div class="iconfont icon-you"></div>
    </div>

    <!-- 商品详情 -->
    <div class="w-100">
      <div class="bg-light-secondary text-center font-md" style="padding:5px 0;">详情</div>
      <img v-for="(img,index) in detail.mains" v-lazy="img" :key="index" class="w-100" @click="previewImage(img)" />
    </div>
    <!-- 热门推荐 -->
    <card headTitle="热门推荐" :headTitleWeight="false" :headBorderBottom="false">
      <van-list class="pb-3" v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad" error-text="加载失败，请重试" :immediate-check="false">
        <div class="row j-sb">
          <common-list v-for="(item,index) in list" :key="index" :item="item" :index="index" type="redirectTo" />
        </div>
      </van-list>
    </card>

    <!-- 底部操作条 -->
    <van-goods-action>
      <van-goods-action-icon icon="chat-o" text="客服" @click="event" />
      <van-goods-action-icon icon="cart-o" text="购物车" @click="open('/cart')" :badge="cartCount" />
      <van-goods-action-icon icon="shop-o" text="店铺" @click="event" />
      <van-goods-action-button type="warning" text="加入购物车" @click="skuShow=true" />
      <van-goods-action-button type="danger" text="立即购买" @click="skuShow=true" />
    </van-goods-action>
    <!-- 服务说明 -->
    <van-popup v-model="serviceShow" position="bottom" round>
      <div class="d-flex a-center j-center font-md border-bottom border-light-secondary" style="height: 50px;">
        服务说明
      </div>
      <div class="p-1" style="height: 200px;">
        <div class="d-flex a-center">
          <span class="iconfont icon-finish main-text-color mr-1"></span>
          小米自营
        </div>
        <span class="text-light-muted font pl-2">不管满多少，就是不包邮</span>
      </div>
      <van-button class="w-100" color="#FD6801" @click="serviceShow = false">确定</van-button>
    </van-popup>

    <!-- 优惠券领取 -->
    <van-popup v-model="coupon.show" position="bottom" round>
      <div class="d-flex a-center j-center font-md border-bottom border-light-secondary" style="height: 50px;">
        领取优惠券
      </div>
      <van-list v-model="coupon.loading" :finished="coupon.finished" finished-text="没有更多了" @load="onCouponLoad" error-text="加载失败，请重试"
                :immediate-check="true" style="height: 420px;overflow: auto">
        <coupon v-for="(item,index) in coupon.list" :key="index" :item="item" :index="index" @draw="onDrawCoupon" />
      </van-list>
      <van-button class="w-100" color="#FD6801" @click="coupon.show = false">确定</van-button>
    </van-popup>

    <van-submit-bar v-if="killInfo.open" :price="killInfo.price*100" button-text="我要秒杀" @submit="killSubmit" />
    <!-- SKU -->
    <van-sku v-else v-model="skuShow" :sku="sku" :goods="goods" :goods-id="detail.id" :quota="0" :quota-used="0" :hide-stock="false"
             @buy-clicked="onBuyClicked" @add-cart="onAddCartClicked" />
  </div>
</template>

<script>
import base from '@/mixin/base.js';
import coupon from '@/components/common/coupon.vue';
import swiperImage from "@/components/index/swiper-image.vue"
import baseInfo from "@/components/detail/base-info.vue"
import productAttrs from "@/components/detail/product-attrs.vue"
import card from "@/components/common/card.vue"
import commonList from "@/components/common/common-list.vue"
import price from "@/components/common/price.vue"
import { mapActions, mapGetters } from "vuex"
import { skuDetail, skuList } from '@/api/product'
import { getCouponList, couponDraw } from '@/api/coupon'
import { getSkuInfo, submitKillOrder } from '@/api/seckill'
import { ImagePreview, Toast } from 'vant'
export default {
  mixins: [base],
  components: {
    swiperImage,
    baseInfo,
    productAttrs,
    card,
    commonList,
    price,
    coupon
  },
  data() {
    return {
      killTime: 0,
      serviceShow: false,
      skuShow: false,
      loading: false,
      finished: false,
      page: 1,
      killInfo: {
        id: 0,
        open: false,
      },
      detail: {
        id: 0,
        cat_id: 0,
        title: "",
        subtitle: "",
        cover: "",
        price: 0,
        is_many: false,
        stock: 100,
        attrs: [],
        skus: [],
        banners: [],
        mains: [],
        sale_attrs: []
      },
      goods: { picture: '' },
      sku: {
        tree: [],
        list: [],
        price: 0,
        stock_num: 0,
        collection_id: 0,
        none_sku: true
      },
      list: [],
      coupon: {
        show: false,
        list: [],
        loading: false,
        finished: false,
        page: 1
      }
    }
  },
  computed: {
    ...mapGetters([
      'cartCount',
      'defaultAddress'
    ])
  },
  mounted() {
    const id = parseInt(this.$route.query.id)
    console.log('id', id)
    if (!id) {
      return this.back()
    }
    const kill = parseInt(this.$route.query.kill)
    this.initData(id, kill)
  },
  methods: {
    ...mapActions([
      'doAddCart'
    ]),
    // 初始化页面
    initData(id, kill) {
      skuDetail(id).then(res => {
        this.detail = res
        //TODO 暂不考虑属性分组的显示方式
        this.detail.attrs = res.attrs[0].items
        this.goods.picture = res.cover
        this.sku.price = res.price
        this.sku.stock_num = res.stock
        this.sku.none_sku = res.is_many == 0
        this.sku.tree = res.sale_attrs.map(item => {
          return {
            k: item.attr_name,
            k_s: 'k_' + item.attr_id,
            v: item.values.map(value => {
              return {
                id: value.id,
                name: value.name
              }
            })
          }
        })
        this.sku.list = res.skus.map(item => {
          let sku = {
            id: item.sku_id,
            price: item.price * 100,
            stock_num: item.stock
          }
          item.attrs.forEach(attr => {
            sku['k_' + attr.attr_id] = attr.value_id
          })
          return sku
        })
        console.log('sku', this.sku)
      }).catch((err) => {
        console.log('err', err)
        this.backToast()
      })
      if (kill > 0) {
        //获取秒杀信息
        getSkuInfo(id).then((res) => {
          this.killInfo = res
          this.killTime = res.start_at * 1000 - (new Date().getTime())
        })
      }
    },
    onBuyClicked(data) {
      console.log('data', data)
      const goods = {
        goods_id: data.goodsId,
        sku_id: data.selectedSkuComb ? data.selectedSkuComb.id : 0,
        num: data.selectedNum,
        cover: this.detail.cover,
        goods_name: this.detail.title,
        price: this.detail.price,
        sku_name: this.chooseAttrs(data),
        type: 'goods'
      }
      this.$router.push({ path: '/order_confirm', query: goods })
    },
    onAddCartClicked(data) {
      console.log('data', data)
      if (!data.selectedSkuComb) {
        return Toast('请选择商品')
      }
      const sku_id = parseInt(data.selectedSkuComb.id)
      const num = data.selectedNum || 1
      //转化为元
      const price = data.selectedSkuComb.price / 100
      let sku_attrs = []
      this.detail.skus.forEach(sku => {
        if (sku.sku_id == sku_id) {
          sku.attrs.forEach(attrs => {
            sku_attrs.push(attrs.value_name)
          })
        }
      })
      const item = {
        sku_id,
        title: this.detail.title,
        cover: this.detail.cover,
        num,
        price,
        sku_attr: sku_attrs.join(',')
      }
      this.skuShow = false
      this.doAddCart(item)
    },
    chooseAttrs(data) {
      if (!data.selectedSkuComb) {
        return []
      }
      let attrs = []
      this.sku.tree.forEach(item => {
        let choose_id = 0
        for (const key in data.selectedSkuComb) {
          if (item.k_s == key) {
            choose_id = data.selectedSkuComb[key]
          }
        }
        console.log('choose_id', choose_id)
        if (choose_id > 0) {
          item.v.forEach(v => {
            if (v.id == choose_id) {
              attrs.push(v.name)
            }
          })
        }
      })
      return attrs.join(', ')
    },
    onLoad() {
      skuList(this.detail.cat_id, this.page).then(res => {
        this.page++
        this.loading = false
        this.list = [...this.list, ...res]
        if (res.length < 20) {
          this.finished = true
        }
      })
    },
    killSubmit() {
      //默认地址
      if (this.defaultAddress.length > 0) {
        const address = this.defaultAddress[0]
        submitKillOrder({
          address_id: parseInt(address.id),
          num: 1,
          sku_id: parseInt(this.detail.id),
          key: this.killInfo.key
        }).then((res) => {
          Dialog.alert({
            message: '秒杀成功啦，订单号：' + res + '，请尽快完成支付哦！',
          }).then(() => {
            this.open('/order');
          });
        })
      } else {
        return this.toast('请设置默认收货地址')
      }
    },
    previewImage(url) {
      const index = this.detail.mains.indexOf(url) || 0;
      ImagePreview({
        images: this.detail.mains,
        startPosition: index,
        closeable: true
      })
    },
    onCouponLoad() {
      getCouponList(this.detail.id).then(res => {
        this.coupon.page++
        this.coupon.loading = false
        this.coupon.list = [...this.coupon.list, ...res]
        if (res.length < 20) {
          this.coupon.finished = true
        }
      })
    },
    onDrawCoupon(item) {
      couponDraw(item.id).then(() => {
        item.status = 2
        Toast.success('领取成功')
      })
    }
  }
}
</script>

<style>
</style>

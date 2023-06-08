<template>
  <div>
    <van-row class="main-bg-color" style="height:45px;">
      <van-col span="3" class="text-center">
        <van-icon name="arrow-left" class="text-white my-1" size="20" @click="back" />
      </van-col>
      <van-col span="18" class="text-center font-md text-white my-1">订单结算</van-col>
      <van-col span="3" />
    </van-row>
    <!-- 地址选择 -->
    <div class="main-bg-color text-white d-flex a-center px-3" style="height: 80px;" @click="openAddress">
      <div class="mb-2">
        <template v-if="address">
          <div class="font-weight font-md d-flex a-center">
            {{address.name}} {{address.phone}}
            <div class="border border-white rounded px-1 font ml-2" v-if="address.is_default">默认</div>
          </div>
          <div class="font">
            {{address.province}}{{address.city}}{{address.county}}&nbsp;&nbsp;{{address.detail}}
          </div>
        </template>
        <template v-else>
          <div class="font-weight font-md d-flex a-center">
            请选择收货地址
          </div>
        </template>
      </div>
      <div class="iconfont icon-you ml-auto mb-2"></div>
    </div>

    <!-- 列表 -->
    <div style="border-top-left-radius: 12px;border-top-right-radius: 12px;margin-top: -12px;overflow: hidden;">
      <order-goods :items="items"></order-goods>
    </div>
    <van-field v-model="message" rows="2" autosize label="订单备注" type="textarea" maxlength="100" placeholder="请输入订单备注" show-word-limit />
    <van-cell title="小计">
      <template #default>
        <span class="main-text-color">￥{{total/100}}</span>
      </template>
    </van-cell>
    <van-divider />
    <!-- 优惠券单元格 -->
    <van-coupon-cell :coupons="coupons" :chosen-coupon="chosenCoupon" @click="showList = true" />
    <!-- 优惠券列表 -->
    <van-popup v-model="showList" round position="bottom" style="height: 90%; padding-top: 4px;">
      <van-coupon-list :coupons="coupons" :chosen-coupon="chosenCoupon" :disabled-coupons="disabledCoupons" @change="onChange"
                       :show-exchange-bar="false" @exchange="onExchange" />
    </van-popup>
    <van-cell title="运费" value="包邮" />

    <!-- 合计 -->
    <van-submit-bar :price="total" button-text="确认订单" @submit="orderConfirm" button-color="#FD6801" />

  </div>
</template>

<script>
import price from "@/components/common/price.vue"
import orderGoods from "@/components/order/order-goods.vue"
import { mapGetters, mapState, mapMutations } from "vuex"
import { getMyCoupon } from '@/api/coupon'
import { submitOrder, submitGoodsOrder } from '@/api/order'
import { Toast } from 'vant'
import ev from '@/utils/event'
export default {
  components: {
    price,
    orderGoods
  },
  data() {
    return {
      type: '',
      total: 0,
      message: '',
      loading: false,
      address: false,
      items: [],
      indexs: [],
      showList: false,
      chosenCoupon: -1,
      coupons: [],
      disabledCoupons: [],
    }
  },
  computed: {
    ...mapState({
      list: state => state.cart.list
    }),
    ...mapGetters([
      'defaultAddress'
    ])
  },
  activated() {
    this.type = this.$route.query.type || ''
    this.items = []
    this.total = 0
    this.indexs = []
    if (this.type == 'goods') {//直接购买商品
      const { goods_id, sku_id, goods_name, sku_name, cover, num, price } = this.$route.query
      if (!goods_id) {
        Toast('请选择商品')
        return this.back()
      }
      this.items.push({
        goods_id: parseInt(goods_id),
        sku_id: parseInt(sku_id),
        num: parseInt(num),
        cover,
        goods_name,
        price: parseInt(price),
        sku_name,
      })
      this.total = this.items[0].price * this.items[0].num * 100
    } else { //购物车确认商品
      this.indexs = this.$route.query.ids ? this.$route.query.ids.split(',') : false
      if (!this.indexs) {
        Toast('请选择商品')
        return this.back()
      }
      for (const key in this.list) {
        if (this.indexs.indexOf(key) !== -1) {
          this.items.push(this.list[key])
          this.total += this.list[key].price * this.list[key].num * 100
        }
      }
    }
  },
  mounted() {
    //设置默认地址
    if (this.defaultAddress.length > 0) {
      this.address = this.defaultAddress[0]
    }
    //获取我的优惠券
    getMyCoupon().then(res => {
      res.forEach(item => {
        let coupon = {
          id: item.id,
          condition: item.name,
          value: item.amount * 100,
          name: item.name,
          startAt: item.start_at,
          endAt: item.end_at,
          valueDesc: item.amount,
          unitDesc: '元',
          description: item.note
        }
        if (this.total < item.min_point * 100) {
          this.disabledCoupons.push(coupon)
        } else {
          this.coupons.push(coupon)
        }
      })
    })
    console.log('address', this.address)
    // 开启监听选择收货地址事件
    ev.$on('chooseAddr', this.onChooseAddr)
  },
  destroyed() {
    // 关闭监听选择收货地址事件
    ev.$off('chooseAddr', this.onChooseAddr)
  },
  methods: {
    ...mapMutations(['delBatchGoods']),
    back() {
      this.$router.back()
    },
    openAddress() {
      this.$router.push({ path: '/user_address', query: { type: 'choose' } })
    },
    open(path) {
      this.$router.push({ path })
    },
    onChooseAddr(item) {
      console.log('choose', item)
      this.address = item
    },
    onChange(index) {
      this.showList = false;
      this.chosenCoupon = index;
    },
    onExchange(code) {
      this.coupons.push(coupon);
    },
    orderConfirm() {
      if (!this.address || !this.address.id) {
        return Toast('请选择收货地址')
      }
      if (this.type == 'goods') {
        return this.submitGoodsOrder()
      }
      this.submitCartOrder()
    },
    submitCartOrder() {
      let address_id = this.address.id
      let coupon_id = this.chosenCoupon === -1 ? 0 : this.coupons[this.chosenCoupon].id
      const sku_ids = this.items.map(v => {
        return v.sku_id
      })
      submitOrder({ coupon_id, address_id, note: this.message, sku_ids }).then(id => {
        console.log('id', id)
        this.delBatchGoods(this.indexs)
        this.$router.push({ path: '/order_detail', query: { id } })
      })
    },
    submitGoodsOrder() {
      submitGoodsOrder({
        sku_id: this.items[0].sku_id,
        num: this.items[0].num || 1,
        address_id: this.address.id,
        coupon_id: this.chosenCoupon === -1 ? 0 : this.coupons[this.chosenCoupon].id,
        note: this.message
      }).then(id => {
        console.log('id', id)
        this.$router.push({ path: '/order_detail', query: { id } })
      })
    }
  }
}
</script>

<style>
.van-button--danger {
  background-color: #fd6801;
}
</style>

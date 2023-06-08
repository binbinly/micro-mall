<template>
  <div>
    <van-sticky>
      <van-row class="main-bg-color" style="height:45px;">
        <van-col span="3" class="text-center">
          <van-icon name="arrow-left" class="text-white my-1" size="20" @click="back" />
        </van-col>
        <van-col span="18" class="text-center font-md text-white my-1">订单详情</van-col>
        <van-col span="3" />
      </van-row>
    </van-sticky>

    <div class="main-bg-color text-white p-4 d-flex a-end j-sb" style="height: 60px;">
      <div class="mb-2">
        <div class="font-lg">{{order.status|formatOrderStatus}}</div>
        <div class="font">{{statusText}}</div>
      </div>
      <div class="iconfont icon-daishouhuo line-h mb-2" style="font-size: 50px;"></div>
    </div>

    <div class="p-1 bg-white">
      <div class="font-md">
        <span class="font-md text-dark mr-1">{{order.address.name}}</span>
        {{order.address.phone}}
      </div>
      <div class="text-light-muted font-sm"> {{order.address.area}}</div>
      <div class="text-light-muted font-sm"> {{order.address.detail}}</div>
    </div>
    <van-divider />
    <order-goods :items="order.items"></order-goods>
    <van-divider />
    <van-cell title="商品总价" :value="'￥'+order.total_amount" />
    <van-cell title="运费" value="包邮" />
    <van-cell title="优惠金额" :value="'-￥'+order.coupon_amount" />
    <van-cell title="订单金额" :value="'￥'+order.amount" />
    <van-divider />
    <van-cell-group title="订单信息">
      <van-cell title="创建时间" :value="order.create_at|formatTime" />
      <van-cell title="订单号" :value="order.order_no" />
      <van-cell title="订单备注" :value="order.note" />
    </van-cell-group>
    <van-cell-group title="退款信息" v-if="order.status == 5">

    </van-cell-group>

    <div v-if="order.status == 1 || order.status == 2 || order.status == 3"
         class="bg-white position-fixed bottom-0 left-0 right-0 d-flex a-center j-end px-1" style="height: 45px;">
      <action-button :order="order" @cancel="onCancel" />
    </div>

  </div>
</template>

<script>
import { orderDetail } from '@/api/order'
import orderGoods from "@/components/order/order-goods.vue"
import actionButton from "@/components/order/action-button.vue"
import price from '@/components/common/price.vue';
import card from '@/components/common/card.vue';
export default {
  components: {
    price,
    card,
    orderGoods,
    actionButton
  },
  data() {
    return {
      id: 0,
      order: {
        id: 0,
        order_no: '',
        note: '',
        total_amount: 0,
        amount: 0,
        coupon_amount: 0,
        pay_type: 0,
        pay_amount: 0,
        pay_at: 0,
        status: 0,
        create_at: 0,
        address: {},
        items: []
      },
    }
  },
  computed: {
    statusText() {
      switch (this.order.status) {
        case 1:
          return '取消'
        case 2:
          return '等待商家发货'
        case 3:
          return '宝贝已发出，请耐心等待'
        case 4:
          return '货物已收到'
        case 5:
          return '订单已完成'
        case 6:
          return '等待商家审核'
        case 7:
          return '查收退款金额到账'
      }
    }
  },
  activated() {
    this.id = this.$route.query.id ? parseInt(this.$route.query.id) : 0
    if (!this.id) {
      Toast('订单错误')
      return this.back()
    }
    this.initData()
  },
  mounted() {
    //绑定滚动事件
    window.addEventListener('scroll', this.scroll, true);
  },
  destroyed() {
    window.removeEventListener('scroll', this.scroll); // 离开页面清除（移除）滚轮滚动事件
  },
  methods: {
    back() {
      this.$router.back()
    },
    scroll(e) {
      //可滚动总高度
      // const clientHeight = document.documentElement.clientHeight || document.body.clientHeight;
      //距离顶部高度
      this.scrollTop = document.documentElement.scrollTop
      //TODO
    },
    initData() {
      orderDetail(this.id).then(res => {
        this.order = res
      })
    },
    onCancel() {
      setTimeout(() => {
        this.back()
      }, 2000)
    }
  }
}
</script>

<style>
</style>

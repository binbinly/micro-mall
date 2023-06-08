<template>
  <div>
    <van-nav-bar title="订单支付" left-arrow @click-left="back" />

    <div class="d-flex flex-column a-center j-center py-2 my-2">
      <span class="text-light-muted font-md">支付金额</span>
      <price priceSize="font-lg" unitSize="font-md" class="p-1" :text="amount"></price>
    </div>
    <van-cell title="当前网络" :value="ethNetwork" />
    <van-cell title="收款地址" v-if="address" value="ETH" :label="address" />

    <van-radio-group v-model="choose">
      <van-cell-group title="选择支付">
        <van-cell v-for="(item,index) in list" :title="item.name" clickable @click="choose = index">
          <template #right-icon>
            <van-radio :name="index" checked-color="#FD6801" />
          </template>
          <template #default>
            <span class="mr-1">0.1 ETH</span>
          </template>
        </van-cell>
      </van-cell-group>
    </van-radio-group>

    <div class="p-1">
      <van-button type="danger" color="#FD6801" :disabled="disabled" block @click="onPay">{{title}}</van-button>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex"
import { orderPayNotify } from '@/api/order'
import { getPayList } from '@/api/app'
import price from "@/components/common/price.vue"
import { Toast, Notify } from 'vant'
export default {
  components: {
    price,
  },
  data() {
    return {
      choose: 0,
      order_no: '',
      amount: 0,
      list: [],
      disabled: false,
      title: '去支付'
    }
  },
  computed: {
    ...mapGetters([
      'ethNetwork'
    ]),
    address() {
      return this.list[this.choose] ? this.list[this.choose].address : ''
    }
  },
  activated() {
    const { order_no, amount } = this.$route.query
    if (!order_no || !amount) {
      Toast('参数缺失')
      return this.back()
    }
    this.order_no = order_no
    this.amount = parseInt(amount)
    this.initData()
  },
  methods: {
    ...mapActions(['createContract', 'doPay', 'tradeCount', 'onNotify']),
    back() {
      this.$router.back()
    },
    initData() {
      getPayList().then(res => {
        console.log('res', res)
        this.list = res
      })
    },
    onPay() {
      this.disabled = true
      this.title = '支付中...'
      const address = this.list[this.choose].address
      //交易号
      const trade_no = new Date().getTime() + '' + parseInt(Math.random() * 90000 + 10000, 10)
      //TODO 支付测试，固定 0.1 ETH
      const eth = 0.1

      //开启监听
      const pay_type = parseInt(this.list[this.choose].id)
      this.onNotify({ address, order_no: this.order_no, trade_no, eth, pay_type, })
      this.doPay({ address, order_no: this.order_no, amount: this.amount, trade_no, eth }).then(res => {
        console.log('res', res)
        orderPayNotify({ trans_hash: res.transactionHash, order_no: this.order_no, trade_no, pay_amount: eth * 100, pay_type })
        this.title = '支付完成'
        this.disabled = false
        this.title = '去支付'
        setTimeout(() => {
          this.back()
        }, 3000)
        Notify({ type: 'success', message: '支付成功' });
      }).catch(() => {
        Notify({ type: 'danger', message: '支付失败' });
        this.disabled = false
        this.title = '去支付'
      })
    }
  }
}
</script>

<style>
</style>

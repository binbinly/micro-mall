<template>
  <div>
    <van-nav-bar title="订单退款" left-arrow @click-left="back" />
    <van-form @submit="onSubmit">
      <van-field v-model="content" rows="2" autosize label="退款理由" type="textarea" maxlength="120" placeholder="请输入退款理由" show-word-limit
                 :rules="[{ required: true }]" />
      <div style="margin: 16px;">
        <van-button round block type="danger" color="#FD6801" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import { Toast, Notify } from 'vant'
import { orderRefund } from '@/api/order'
export default {
  data() {
    return {
      order_id: 0,
      content: ''
    }
  },
  activated() {
    this.order_id = parseInt(this.$route.query.order_id);
    if (!this.order_id) {
      Toast('参数缺失')
      return this.back()
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
    onSubmit() {
      orderRefund({ order_id: this.order_id, content: this.content }).then(() => {
        Notify({ type: 'success', message: '申请成功，等待商家审核' });
        setTimeout(() => {
          this.back();
        }, 1000)
      })
    }
  }
}
</script>

<style>
</style>

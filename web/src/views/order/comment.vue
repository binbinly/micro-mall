<template>
  <div>
    <van-nav-bar title="评价" left-arrow @click-left="back" />
    <van-form @submit="onSubmit">
      <van-field name="rate" label="评分">
        <template #input>
          <van-rate v-model="star" />
        </template>
      </van-field>
      <van-field v-model="content" rows="2" autosize label="内容" type="textarea" maxlength="120" placeholder="请填写评价内容" show-word-limit
                 :rules="[{ required: true }]" />
      <div style="margin: 16px;">
        <van-button round block type="danger" color="#FD6801" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import { Toast, Notify } from 'vant'
import { orderComment } from '@/api/order'
export default {
  data() {
    return {
      order_id: '',
      star: 0,
      content: '',
      ids: []
    }
  },
  activated() {
    const { order_id, ids } = this.$route.query;
    if (!order_id || !ids) {
      Toast('参数缺失')
      return this.back()
    }
    this.order_id = parseInt(order_id)
    this.ids = ids.split(',').map(id => parseInt(id))
  },
  methods: {
    back() {
      this.$router.back()
    },
    onSubmit() {
      orderComment({ order_id: this.order_id, content: this.content, sku_ids: this.ids, star: this.star }).then(() => {
        Notify({ type: 'success', message: '感觉您的宝贵评价' });
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

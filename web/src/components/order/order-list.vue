<template>
  <div class="bg-white" @click.stop="openDetail">
    <!-- 头部 -->
    <div class="d-flex a-center p-1 border-bottom border-light-secondary">
      <span class="text-light-muted font-sm">{{item.time|formatTime}}</span>
      <span class="main-text-color ml-auto font-sm">{{item.status|formatOrderStatus}}</span>
    </div>
    <!-- 商品 -->
    <order-goods :items="item.items"></order-goods>
    <!-- 底部 -->
    <div class="d-flex a-center p-1">
      <span class="text-dark font-sm ml-auto">
        共&nbsp;{{item.items.length}}&nbsp;件商品，合计：￥{{item.amount}}
      </span>
    </div>
    <div class="d-flex j-end a-center px-1 pb-1">
      <action-button :order="item" :index="index" @cancel="onCancel" />
    </div>
  </div>
</template>

<script>
import orderGoods from './order-goods';
import actionButton from './action-button';
export default {
  components: {
    orderGoods,
    actionButton
  },
  props: {
    item: Object,
    index: Number
  },
  methods: {
    openDetail() {
      this.$router.push({ path: '/order_detail', query: { id: this.item.id } })
    },
    onCancel() {
      this.$emit('cancel', this.index)
    }
  }
}
</script>

<style>
</style>

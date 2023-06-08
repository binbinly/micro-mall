<template>
  <div class="px-1" style="width:44%" @click="openDetail">
    <van-image style="display:flex" :src="item.cover" fit="cover" />
    <div class="p-1 pt-1">
      <div><span class="text-inline-3 font-sm">{{item.title}}</span></div>
      <div class="d-flex py-1">
        <price :text="item.price"></price>
        <div v-if="item.original_price" class="ml-1 a-self-end line-h"><span
                class="font-sm text-light-muted line-through">￥{{item.original_price}}</span></div>
        <div v-else class="ml-1 a-self-end line-h"><span class="font-sm text-light-muted">{{item.sale_count}}</span></div>
      </div>
    </div>
  </div>
</template>

<script>
import price from "@/components/common/price.vue";
export default {
  components: {
    price
  },
  props: {
    item: Object,
    index: [Number, String]
  },
  methods: {
    openDetail() {
      const query = { id: this.item.id }
      if (this.item.start_at) {
        query['kill'] = 1
      }
      this.$router.push({ path: '/product_detail', query })
    }
  }
}
</script>

<style>
</style>

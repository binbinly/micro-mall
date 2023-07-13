<template>
  <div>
    <div v-for="(item,index) in items" :key="index" @click.stop="openGoods(item.sku_id)">
      <div class="d-flex a-center pt-1 border-bottom bg-white border-light-secondary">
        <van-image class="pl-1 border border-light-secondary rounded flex-shrink" :src="item.cover" fit="cover" width="75" />

        <div class="flex-1 d-flex flex-column pl-1" style="width:280px;">
          <div class="text-dark font-sm text-inline">{{item.title}}</div>
          <!-- 规格属性 -->
          <div class="d-flex text-light-muted font" style="padding:5px 0;">
            {{item.attr_value || '默认'}}
          </div>
          <div class="mt-auto d-flex j-sb">
            <price :text="item.price" class="a-center" />
            <span class="mr-1 font-sm">x{{item.num}}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import price from "@/components/common/price.vue"
export default {
  components: {
    price
  },
  props: {
    items: Array
  },
  methods: {
    openGoods(id) {
      this.$router.push({ path: '/product_detail', query: { id } })
    }
  }
}
</script>

<style>
</style>

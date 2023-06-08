<template>
  <div class="p-1">
    <div class="font-md">{{detail.title}}</div>
    <div class="font text-light-muted mb-2 line-h-sm">
      {{detail.subtitle}}
    </div>
    <price priceSize="font-lg" unitSize="font" :text="detail.price"></price>
  </div>
</template>

<script>
import price from "@/components/common/price.vue"
export default {
  components: {
    price
  },
  props: ['detail'],
}
</script>

<style>
</style>

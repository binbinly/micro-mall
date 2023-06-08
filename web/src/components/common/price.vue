<template>
  <div class="d-flex line-h">
    <span :class="unitSize+' '+color">￥</span>
    <span :class="priceSize+' '+color">
      <slot />{{text}}
    </span>
  </div>
</template>

<script>
export default {
  props: {
    text: {
      type: [String, Number],
      default: ""
    },
    priceSize: {
      type: String,
      default: "font-md"
    },
    unitSize: {
      type: String,
      default: "font-sm"
    },
    color: {
      type: String,
      default: "main-text-color"
    }
  }
}
</script>

<style>
</style>

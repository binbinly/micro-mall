<template>
  <div>
    <van-swipe :autoplay="5000" :duration="200" indicator-color="white" loop>
      <van-swipe-item v-for="(image, index) in images" :key="index">
        <van-image :src="image" :height="height" width="100%" fit="cover" lazy-load />
      </van-swipe-item>
    </van-swipe>
  </div>
</template>

<script>
export default {
  props: {
    images: [Array, Object],
    height: {
      type: String,
      default: "150"
    },
  }
}
</script>

<style>
</style>

<template>
  <div>
    <van-nav-bar title="物流信息" left-arrow @click-left="back" />
    <van-empty description="暂无物流信息" />
  </div>
</template>

<script>
export default {
  data() {
    return {
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
  }
}
</script>

<style>
</style>

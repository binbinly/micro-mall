<template>
  <div>
    <van-nav-bar title="关于商城" left-arrow @click-left="back" />

    <div class="d-flex flex-column j-center a-center py-2 my-3">
      <van-image :src="require('@/assets/images/about.jpg')" width="100" fit="cover" />
      <div class="text-light-muted font-md">
        xxx商城 1.0.0
      </div>
    </div>

    <van-cell-group title="商城信息">
      <van-cell title="版本更新" is-link />
      <van-cell title="商城协议" is-link />
      <van-cell title="历史记录" is-link />
      <van-cell title="帮助与反馈" is-link />
    </van-cell-group>
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
    }
  }
}
</script>

<style>
</style>

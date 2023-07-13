<template>
  <div>
    <van-nav-bar title="用户设置" left-arrow @click-left="back" />

    <van-cell-group v-if="token" title="账号管理">
      <van-cell title="个人资料" is-link />
      <van-cell title="收货地址" is-link to="/user_address" />
    </van-cell-group>

    <van-cell-group title="关于">
      <van-cell v-for="(item,index) in list" :key="index" :title="item.title" is-link @click="onPush(item.path)" />
    </van-cell-group>

    <div class="p-1" v-if="token">
      <van-button type="danger" block @click="onLogout">退出登录</van-button>
    </div>
  </div>
</template>

<script>
import { mapMutations, mapGetters } from 'vuex'
import { userLogout } from '@/api/user'
import { Toast, Dialog } from 'vant'
export default {
  data() {
    return {
      list: [
        { title: "关于商城", path: "/about" },
        { title: "意见反馈", path: "" },
        { title: "协议规则", path: "" },
        { title: "资质证件", path: "" },
        { title: "用户协议", path: "" },
        { title: "隐私政策", path: "" },
      ]
    }
  },
  computed: {
    ...mapGetters(['token'])
  },
  methods: {
    ...mapMutations(['logout', 'clearCart']),
    back() {
      this.$router.back()
    },
    onPush(path) {
      if (path) {
        return this.$router.push({ path })
      }
      Toast('待开发')
    },
    // 退出登录
    onLogout() {
      Dialog.confirm({
        title: '提示',
        message: '是否退出当前账号？',
      })
        .then(() => {
          userLogout().then(() => {
            this.clear()
          }).catch(err => {
            this.clear()
          })
        })
        .catch(() => {
          // on cancel
        });
    },
    clear() {
      // 退出登录
      this.logout();
      // 清空购物车
      this.clearCart()
      Toast.success('退出成功')
      this.back()
    }
  }
}
</script>

<style>
</style>

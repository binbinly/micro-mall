<template>
  <div>
    <van-nav-bar title="我的钱包" left-arrow @click-left="back" />

    <div class="d-flex flex-column a-center j-center py-2 my-2">
      <span class="text-light-muted font-md">余额</span>
      <span class="font-lg p-1 main-text-color">{{balance}} ETH</span>
    </div>

    <van-cell title="当前网络" :value="ethNetwork" />
    <van-cell title="我的地址" value="ETH" :label="account" />

  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex"
export default {
  data() {
    return {
      balance: 0,
      address: '',
      list: [],
      count: 0,
    }
  },
  computed: {
    ...mapGetters([
      'ethNetwork',
      'account'
    ])
  },
  mounted() {
    this.initData()
  },
  methods: {
    ...mapActions(['getBalance', 'tradeCount', 'tradeLog']),
    back() {
      this.$router.back()
    },
    initData() {
      this.getBalance(this.account).then(val => {
        this.balance = val
      })
      // this.tradeCount(this.account).then(res => {
      //   console.log('res', res)
      // })
    },
  }
}
</script>

<style>
</style>

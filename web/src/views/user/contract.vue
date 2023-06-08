<template>
  <div>
    <van-nav-bar title="我的合约" left-arrow @click-left="back" />

    <div class="d-flex flex-column a-center j-center py-2 my-2">
      <span class="text-light-muted font-md">合约余额</span>
      <span class="font-lg p-1 main-text-color">{{balance}} ETH</span>
    </div>

    <van-cell title="当前网络" :value="ethNetwork" />
    <van-cell title="合约地址" value="ETH" :label="address" />
    <div class="p-1" v-if="address">
      <van-form @submit="onWithdraw">
        <van-field v-model="amount" type="digit" label="提取数额" :rules="[{ required: true, message: '请输入' }]" />
        <div class="pt-1">
          <van-button round block type="danger" native-type="submit">我要提款</van-button>
        </div>
      </van-form>
    </div>
    <div class="p-1" v-else>
      <van-button type="danger" color="#FD6801" block @click="onCreate">创建我的支付合约</van-button>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex"
import { setStorage, getStorage } from '@/utils/index'
import { Toast, Notify } from 'vant'
export default {
  data() {
    return {
      balance: 0,
      address: '',
      amount: 1
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
    ...mapActions(['createContract', 'getBalance', 'doWithdraw']),
    back() {
      this.$router.back()
    },
    initData() {
      this.address = getStorage('contract_address')
      if (!this.address) {
        return
      }
      this.getBalance(this.address).then(val => {
        console.log('banance: ' + val)
        this.balance = val
      })
    },
    onCreate() {
      this.createContract().then(contract => {
        console.log('my contract address', contract.options.address)
        setStorage('contract_address', contract.options.address)
        this.address = contract.options.address
        Notify({ type: 'success', message: '创建成功' });
      })
    },
    onWithdraw() {
      if (this.balance < this.amount) {
        return Toast('余额不足哦')
      }
      this.doWithdraw(this.account, this.amount).then(() => {
        Notify({ type: 'success', message: '提款成功' });
      }).catch(err => {
        Notify({ type: 'danger', message: '提款失败了' });
      })
    }
  }
}
</script>

<style>
</style>

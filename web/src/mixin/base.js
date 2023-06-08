import { Toast } from 'vant'
export default {
  methods: {
    back() {
      this.$router.back()
    },
    open(path, params = null) {
      this.$router.push({ path, params })
    },
    event() {
      Toast('待开发')
    },
    // 返回并提示
    backToast(msg = '非法参数') {
      this.toast(msg)
      setTimeout(() => {
        this.$router.back()
      }, 1000)
    },
    toast(msg = '非法参数') {
      Toast({ message: msg, position: 'bottom' })
    }
  }
}

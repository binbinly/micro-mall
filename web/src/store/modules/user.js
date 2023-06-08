import { setStorage, getStorage, removeStorage } from '@/utils/index.js'

// 用户
export default {
  state: {
    // token
    token: null,
    // 用户信息
    user: null
  },
  mutations: {
    // 初始化登录状态
    initUser(state) {
      const user = getStorage('user')
      if (user) {
        state.user = JSON.parse(user)
      }
      state.token = getStorage('token')
    },
    // 登录
    login(state, data) {
      data.member.name = data.member.nickname || data.member.username
      state.user = data.member
      state.token = data.token
      // 持久化存储
      // 存储到本地存储中
      setStorage('token', data.token)
      setStorage('user', JSON.stringify(data.member))
    },
    // 退出登录
    logout(state) {
      state.user = null
      state.token = null
      removeStorage('token')
      removeStorage('user')
    }
  },
  actions: {
    // 初始化登录状态
    initUser({ commit }) {
      const user = getStorage('user')
      if (user) {
        state.user = JSON.parse(user)
      }
      state.token = getStorage('token')
    }
  }
}

import axios from 'axios'
import $store from '@/store'
import $router from '@/router/index.js'
import { Toast } from 'vant'
// 根据环境不同引入不同api地址
import { baseApi } from '@/config'

// create an axios instance
const service = axios.create({
  baseURL: baseApi, // url = base api url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 8000, // request timeout
  headers: { 'Content-Type': 'application/json;charset=UTF-8' }
})

// request拦截器 request interceptor
service.interceptors.request.use(
  config => {
    // 不传递默认开启loading
    if (!config.hideLoading) {
      // loading
      Toast.loading({
        forbidClick: true
      })
    }
    if (config.auth === true) {
      const token = $store.getters.token
      if (!token) {
        $router.replace({ path: '/login' })
        return Promise.reject(new Error('请登录'))
      }
      config.headers['Token'] = token || ''
    }
    return config
  },
  error => {
    // do something with request error
    console.log('req err:', error)
    return Promise.reject(error)
  }
)
// response拦截器
service.interceptors.response.use(
  response => {
    Toast.clear()
    if (response.status === 200) {
      const result = response.data
      if (result.code === 0) {
        if (response.config.result) {
          return Promise.resolve(result)
        }
        return Promise.resolve(result.data)
      } else if (result.code == 104) {
        Toast('令牌已过期，请重新登录')
        $store.commit('logout')
        $router.push({ path: '/login' })
        return Promise.reject(new Error(result.message))
      }
      Toast.fail(result.message)
      return Promise.reject(new Error(result.message || 'Error'))
    } else {
      Toast.fail('网络开小差了')
      console.log('response err', response)
      return Promise.reject(new Error(response.statusText))
    }
  },
  error => {
    Toast.clear()
    if (error.message === 'Network Error') {
      Toast.fail('服务器连接异常，请检查服务器！')
      return Promise.reject(error)
    }
    console.log('rsp err:', error) // for debug

    Toast.fail(error.message)
    return Promise.reject(error)
  }
)

export default service

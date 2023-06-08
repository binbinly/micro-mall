<template>
  <div>
    <div class="d-flex a-center j-sb py-2 px-3 text-light-muted">
      <div class="iconfont icon-shanchu1" @click="back"></div>
    </div>

    <div class="pt-5">
      <div class="text-center" style="padding-bottom: 20px;font-size: 25px;">{{titleName}}</div>
      <van-form @submit="onSubmit">
        <template v-if="status == 0">
          <van-field v-model="username" name="username" label="用户名" placeholder="请输入用户名" :rules="[{ required: true }]" />
          <van-field v-model="password" type="password" name="password" label="密码" placeholder="密码长度6-20位'"
                     :rules="[{ required: true, validator: pwdValid, message: '密码长度6-20位' }]" />
        </template>
        <template v-else-if="status == 2">
          <van-field v-model="phone" name="phone" label="手机号" placeholder="请输入手机号" :rules="[{ required: true, validator: phoneValid }]" />
          <van-field v-model="username" name="username" label="用户名" placeholder="请输入用户名" :rules="[{ required: true }]" />
          <van-field v-model="password" type="password" name="password" label="密码" placeholder="请输入密码"
                     :rules="[{ required: true, validator: pwdValid, message: '密码长度6-20位' }]" />
          <van-field v-model="confirm_password" type="password" name="confirm_password" label="确认密码" placeholder="请确认密码"
                     :rules="[{ required: true, validator: confirmPwdValid, message: '密码长度6-20位' }]" />
          <div class="d-flex a-stretch mr-1">
            <van-field v-model="code" name="code" label="短信验证码" placeholder="请输入短信验证码" :rules="[{ required: true }]" />
            <div class="van-field__button" style="margin-top:5px;">
              <van-button size="small" type="primary" native-type="button" @click="getCode">{{ codeTime === 0 ? '发送验证码' : codeTime }}</van-button>
            </div>
          </div>
        </template>
        <template v-else>
          <van-field v-model="phone" name="phone" label="手机号" placeholder="请输入手机号" :rules="[{ required: true, validator: phoneValid }]" />
          <div class="d-flex a-stretch mr-1">
            <van-field v-model="code" name="code" label="短信验证码" placeholder="请输入短信验证码" :rules="[{ required: true }]" />
            <div class="van-field__button" style="margin-top:5px;">
              <van-button size="small" type="primary" native-type="button" @click="getCode">{{ codeTime === 0 ? '发送验证码' : codeTime }}</van-button>
            </div>
          </div>
        </template>
        <div style="margin: 16px;">
          <van-button class="text-white" :class="disabled ? 'main-bg-hover-color':'main-bg-color'" block native-type="submit">
            提交
          </van-button>
        </div>
      </van-form>
    </div>
    <div class="d-flex a-center j-center pt-2 pb-2">
      <div class="text-primary font-sm" @click="changeStatus(0)">
        {{status == 1?'账号密码登录':'验证码登陆'}}
      </div>
      <span class="text-muted mx-2">|</span>
      <div class="text-primary font-sm" @click="changeStatus(2)">还没有账号？</div>
    </div>
    <van-divider>社交账号登录</van-divider>
    <div class="d-flex a-center j-center text-muted mt-2">
      注册即代表同意<span class="text-primary">《社区协议》</span>
    </div>

  </div>
</template>

<script>
const statusName = ['账号密码登录', '手机验证码登录', '新用户注册']
import { mapMutations, mapActions } from 'vuex';
import authApi from '@/api/auth.js'
import base from '@/mixin/base.js';
export default {
  mixins: [base],
  data() {
    return {
      status: 0,
      username: "",
      password: "",
      confirm_password: "",
      phone: "",
      code: "",
      codeTime: 0,
      loading: false
    }
  },
  computed: {
    disabled() {
      if ((this.username === '' || this.password === '') && (this.phone === '' || this.code === '')) {
        return true
      }
      return false
    },
    titleName() {
      return statusName[this.status]
    }
  },
  methods: {
    ...mapMutations(['login']),
    ...mapActions(['initCart']),
    // 初始化表单
    initForm() {
      this.username = ''
      this.password = ''
      this.confirm_password = ''
      this.phone = ''
      this.code = ''
    },
    changeStatus(s) {
      // 初始化表单
      this.initForm()
      if (s == 0) {
        this.status = this.status == 0 ? 1 : 0
      } else {
        this.status = s
      }
    },
    // 获取验证码
    getCode() {
      // 防止重复获取
      if (this.codeTime > 0) return;
      // 验证手机号
      if (!this.phoneValid(this.phone)) return;
      authApi.sendCode({
        phone: this.phone
      }).then(res => {
        this.toast('验证码：' + res)
        // 倒计时
        this.codeTime = 60
        let timer = setInterval(() => {
          if (this.codeTime >= 1) {
            this.codeTime--
          } else {
            this.codeTime = 0
            clearInterval(timer)
          }
        }, 1000)
      })
    },
    phoneValid(val) {
      var mPattern = /^1[345678]\d{9}$/;
      return mPattern.test(this.phone)
    },
    pwdValid(val) {
      return val.length >= 6 && val.length <= 20
    },
    confirmPwdValid(val) {
      return val === this.password
    },
    register() {
      authApi.register({
        phone: this.phone,
        code: this.code,
        username: this.username,
        password: this.password,
        confirm_password: this.confirm_password,
      }).then(res => {
        this.toast('注册成功')
        this.status = 0
        // 初始化表单
        this.initForm()
      })
    },
    loginDo(res) {
      // 修改vuex的state,持久化存储
      this.login(res)
      // 刷新购物车
      this.initCart()
      this.backToast('登录成功')
    },
    // 提交
    onSubmit() {
      if (this.status == 1) { // 手机验证码登录
        authApi.loginPhone({
          phone: this.phone,
          code: this.code
        }).then(res => {
          this.loginDo(res)
        })
      } else if (this.status == 0) { // 账号密码登录
        authApi.login({
          username: this.username,
          password: this.password
        }).then(res => {
          this.loginDo(res)
        })
      } else { //注册
        this.register()
      }
    }
  }
}
</script>

<style>
</style>

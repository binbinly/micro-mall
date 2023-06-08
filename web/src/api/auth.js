import api from './index'

const authApi = {
  login(data) {
    return api.post(api.Member.Login, data, null, false)
  },
  loginPhone(data) {
    return api.post(api.Member.PhoneLogin, data, null, false)
  },
  register(data) {
    return api.post(api.Member.Register, data, null, false)
  },
  sendCode(data) {
    return api.post(api.SendCode, data, null, false)
  }
}

export default authApi

import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import user from './modules/user'
import address from './modules/address'
import cart from './modules/cart'
import pay from './modules/pay'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    user,
    address,
    cart,
    pay
  },
  getters
})

export default store

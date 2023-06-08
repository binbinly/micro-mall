import { userAddressList, userAddressAdd, userAddressEdit, userAddressDel } from '@/api/user'
import { areaList } from '@/api/app'
export default {
  state: {
    list: [],
    isInit: false, //是否已初始化
    area: null
  },
  mutations: {
    // 覆盖收货地址
    initAddress(state) {
      if (state.isInit) {
        return
      }
      userAddressList().then(res => {
        state.list = res
        state.isInit = true
      })
    },
    // 创建收货地址
    pushList(state, item) {
      state.list.unshift(item)
    },
    // 删除收货地址
    delList(state, index) {
      state.list.splice(index, 1)
    },
    // 修改收货地址
    editList(state, { index, data }) {
      for (let key in data) {
        state.list[index][key] = data[key]
      }
    },
    // 取消默认地址
    removeDefault(state) {
      state.list.forEach(v => {
        if (v.is_default) {
          v.is_default = 0
        }
      })
    },
    initArea(state) {
      if (state.area) {
        return
      }
      areaList().then(res => {
        state.area = {
          province_list: res.province,
          city_list: res.city,
          county_list: res.county
        }
      })
    }
  },
  actions: {
    //添加地址
    addAddress({ commit }, data) {
      return new Promise((res, rej) => {
        userAddressAdd(data)
          .then(id => {
            if (data.is_default) {
              commit('removeDefault')
            }
            data.id = id
            commit('pushList', data)
            res(id)
          })
          .catch(err => {
            rej(err)
          })
      })
    },
    //修改地址
    editAddress({ commit }, { index, data }) {
      userAddressEdit(data).then(() => {
        if (data.is_default) {
          commit('removeDefault')
        }
        commit('editList', { index, data })
      })
    },
    delAddress({ commit }, { index, id }) {
      userAddressDel(id).then(() => {
        commit('delList', index)
      })
    }
  }
}

<template>
  <div>
    <van-nav-bar :title="navTitle" left-text="返回" left-arrow @click-left="back" />

    <van-address-edit :area-list="areaList" :address-info="address" :show-delete="index>-1" show-set-default
                      :area-columns-placeholder="['请选择', '请选择', '请选择']" @save="onSave" @delete="onDelete" />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex"
import { areaList } from '@vant/area-data';
import { Toast } from 'vant'
export default {
  data() {
    return {
      navTitle: '添加地址',
      areaList,
      index: -1,
      address: {},
    }
  },
  activated() {
    this.index = this.$route.query.index ? parseInt(this.$route.query.index) : -1
    if (this.index > -1) {
      this.navTitle = '修改地址'
      this.address = {
        id: this.list[this.index].id,
        name: this.list[this.index].name,
        tel: this.list[this.index].phone,
        province: this.list[this.index].province,
        city: this.list[this.index].city,
        county: this.list[this.index].county,
        addressDetail: this.list[this.index].detail,
        areaCode: this.list[this.index].area_code + '',
        isDefault: this.list[this.index].is_default > 0,
      }
    } else {
      this.navTitle = '添加地址'
      this.address = {}
    }
  },
  computed: {
    ...mapState({
      list: state => state.address.list
    }),
  },
  methods: {
    ...mapActions(['addAddress', 'editAddress', 'delAddress']),
    back() {
      this.$router.back()
    },
    onSave(data) {
      let addr = {
        name: data.name,
        phone: data.tel,
        province: data.province,
        city: data.city,
        county: data.county,
        area_code: parseInt(data.areaCode),
        detail: data.addressDetail,
        is_default: data.isDefault ? 1 : 0
      }
      if (this.index > -1) {
        addr.id = this.list[this.index].id
        return this.editAddress({ index: this.index, data: addr }).then(() => {
          this.back()
        })
      }
      this.addAddress(addr).then(() => {
        this.back()
      })
    },
    onDelete() {
      if (this.index === -1) {
        return Toast('非法操作')
      }
      this.delAddress({ index: this.index, id: this.list[this.index].id }).then(() => {
        this.back()
      })
    },
  }
}
</script>

<style>
.van-button--danger,
.van-switch--on {
  background-color: #fd6801;
}
</style>

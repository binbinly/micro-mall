<template>
  <div>
    <div @click="show=true">
      <div class="scroll-row-item" style="width: 24%;height: 55px;" v-for="(item,index) in attrs" :key="index">
        <div class="d-flex flex-column a-center j-center">
          <div class="iconfont line-h-sm icon-cpu"></div>
          <div class="line-h-sm">{{item.name}}</div>
          <div class="text-light-muted font line-h-sm">{{item.value}}</div>
        </div>
      </div>
    </div>

    <van-popup v-model="show" position="bottom" round>
      <div class="d-flex a-center j-center font-md border-bottom border-light-secondary" style="height: 50px;">
        参数配置
      </div>
      <div style="height: 400px;overflow: auto;">
        <van-cell :title="item.name" :value="item.value" v-for="(item,index) in baseAttrs" :key="index"/>
      </div>
      <van-button class="w-100" type="default" color="#FD6801" @click="show = false">确定</van-button>
    </van-popup>
  </div>
</template>

<script>
export default {
  props: ['baseAttrs'],
  data() {
    return {
      show: false
    }
  },
  computed: {
    attrs() {
      let list = this.baseAttrs.map(item => {
        item.name = item.name.substring(0, 5)
        item.value = item.value.substring(0, 5)
        return item
      })
      return list.slice(0, 8)
    }
  }
}
</script>

<style>
</style>

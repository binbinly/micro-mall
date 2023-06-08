<template>
  <div class="row j-center m-1">
    <template v-for="(item,index) in resdata">
      <div class="d-flex flex-column j-center a-center p-1" @click="event">
        <van-image width="32" height="32" :src="item.icon" fit="cover" lazy-load />
        <span class="font-small mt-1">{{item.title}}</span>
      </div>
    </template>

  </div>
</template>

<script>
import { Toast } from 'vant'
export default {
  props: {
    resdata: [Array, Object]
  },
  methods: {
    event() {
      Toast('待开发')
    }
  }
}
</script>

<style>
</style>

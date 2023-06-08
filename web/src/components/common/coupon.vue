<template>
  <div class="bg-white rounded border d-flex a-center border-light-secondary mb-1">
    <div class="flex-1 d-flex flex-column j-center px-1">
      <div class="font-md">{{item.name}}</div>
      <div class="font text-light-muted">
        {{item.start_at|formatTime}}~{{item.end_at|formatTime}}
      </div>
      <div class="font text-light-muted">{{item.note}}</div>
    </div>
    <div class="text-white d-flex flex-column a-center j-center main-bg-color mr-1" style="width: 110px;height: 80px;">
      <div style="font-size: 20px;" class="mb-1">
        {{item.amount}}<span class="font">元</span></div>

      <van-button v-if="item.status==0" color="linear-gradient(to right, #ff6034, #ee0a24)" size="small" @click="onDraw">马上领取</van-button>
      <span v-else>已领取</span>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    item: Object,
    index: Number
  },
  filters: {
    formatTime(shorttime) {
      shorttime = shorttime.toString().length < 13 ? shorttime * 1000 : shorttime;
      let date = new Date(shorttime)
      let parseNumber = (num) => {
        return num < 10 ? "0" + num : num;
      }
      return date.getFullYear() + '-' + parseNumber(date.getMonth() + 1) + '-' + parseNumber(date.getDate()) + ' ' + parseNumber(date.getHours()) + ':' + parseNumber(date.getMinutes())
    }
  },
  methods: {
    onDraw() {
      this.$emit('draw', this.item)
    }
  }
}
</script>

<style>
</style>

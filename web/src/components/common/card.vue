<template>
  <div class="card" :style="cardStyle">
    <!-- head -->
    <div v-if="showhead" class="p-1 main-border-color d-flex a-center j-sb" :class="getHeadClass">
      <slot name="title">
        <span v-if="headTitle" class="font" :class="headTitleWeight?'font-weight':''">{{headTitle}}</span>
      </slot>
      <slot name="right"></slot>
    </div>
    <!-- body -->
    <div :class="getBodyClass" :style="bodyStyle">
      <van-image v-if="bodyCover" :src="bodyCover" fit="cover" lazy-load />
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    bodyStyle: String,
    // 头部标题
    headTitle: String,
    // 封面图
    bodyCover: String,
    // 是否显示头部
    showhead: {
      type: Boolean,
      default: true
    },
    // 是否显示下边线
    headBorderBottom: {
      type: Boolean,
      default: true
    },
    // 标题是否加粗
    headTitleWeight: {
      type: Boolean,
      default: true
    },
    // 是否给body加padding
    bodyPadding: {
      type: Boolean,
      default: false
    },
    cardStyle: {
      type: String,
      default: ""
    }
  },
  computed: {
    getHeadClass() {
      let BorderBottom = this.headBorderBottom ? 'border-bottom' : ''
      return `${BorderBottom}`
    },
    getBodyClass() {
      let padding = this.bodyPadding ? 'p-1' : ''
      return `${padding}`
    }
  }
}
</script>

<style>
</style>

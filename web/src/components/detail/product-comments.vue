<template>
  <div class="p-2">

    <div class="scroll-row-item rounded bg-light-secondary mr-2 p-2" v-for="(item,index) in comments" :key="index" style="width: 155px;height: 95px;">

      <div class="d-flex a-center mb-1">
        <van-image :src="item.avatar" fit="cover" width="35" round />
        <div class="ml-2">
          <span class="font-md line-h">
            {{item.username}}</span>
          <span class="d-block font text-light-muted line-h">{{item.create_time | formatTime}}</span>
        </div>
        <!-- <div class="iconfont icon-dianzan text-light-muted font pl-1 ml-auto">{{item.goods_num}}</div> -->
      </div>

      <span class="d-block font-md mb-1">{{item.context}}</span>

      <div class="row">
        <div class="span24-8 px-1" v-for="(img,i) in item.imglist" :key="i">
          <van-image :src="img" fit="cover" width="58" round />
        </div>
      </div>

    </div>

    <div class="d-flex a-center j-center py-2 mt-2 text-primary" hover-class="bg-light-secondary" @click="open">
      更多评论 <div class="iconfont icon-you"></div>
    </div>

  </div>
</template>

<script>
export default {
  props: ['comments', 'goods_id'],
  methods: {
    open() {
      this.$router.push({ path: '/goods_comment', query: { id: this.goods_id } })
    }
  }
}
</script>

<style>
</style>

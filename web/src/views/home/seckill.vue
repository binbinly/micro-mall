<template>
  <div>
    <van-nav-bar title="每日秒杀" left-arrow @click-left="back" />
    <van-tabs v-model="active" @change="onChangeTab" color="#FD6801" title-active-color="#FD6801" offset-top="45" sticky animated swipeable>
      <van-tab v-for="(item,index) in tabBars" :key="index" :title="item.name">
        <template v-if="item.skus.length > 0">
          <van-list v-model="item.loading" :finished="item.finished" finished-text="没有更多了" @load="onLoad" error-text="加载失败，请重试">
            <div class="row j-sb">
              <common-list v-for="(item2,index2) in item.skus" :key="index2" :item="item2" :index="index2" />
            </div>
          </van-list>
        </template>

        <van-empty v-else style="height:500px;" description="没有数据哦" />

      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import commonList from "@/components/common/common-list.vue";
import { getSessions, getSkus } from '@/api/seckill.js';
export default {
  components: {
    commonList
  },
  data() {
    return {
      active: 0,
      tabBars: [],
    }
  },
  mounted() {
    this.initData()
  },
  methods: {
    back() {
      this.$router.back()
    },
    // 初始化事件
    initData() {
      // 获取顶部选项卡
      getSessions().then((res) => {
        this.tabBars = res
        this.tabBars.forEach((item) => {
          item['loading'] = false
          item['finished'] = false
        })
      })
    },
    onLoad() {
      const index = this.active
      getSkus(this.tabBars[index].id).then((res) => {
        this.tabBars[index].loading = false
        this.tabBars[index].skus = res
        // 数据全部加载完成
        this.tabBars[index].finished = true;
      }).catch((err) => {
        this.tabBars[index].finished = true;
      })
    },
    // 监听滑动列表
    onChangeTab(name, item) {
      if (this.tabBars[this.active].skus.length > 0) {
        return
      }
      this.onLoad()
    },
  }
}
</script>

<style>
</style>

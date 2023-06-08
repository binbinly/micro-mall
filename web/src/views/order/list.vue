<template>
  <div>
    <van-nav-bar title="我的订单" left-arrow @click-left="back" fixed placeholder />

    <!-- 订单列表 -->
    <van-tabs v-model="active" @change="onChange" color="#FD6801" title-active-color="#FD6801" offset-top="45" sticky animated swipeable>
      <van-tab v-for="(item,index) in tabBars" :title="item.name">
        <van-pull-refresh v-model="item.refreshing" @refresh="onRefresh">
          <template v-if="item.empty === false">
            <van-list v-model="item.loading" :finished="item.finished" finished-text="没有更多了" :immediate-check="false" @load="onLoad"
                      error-text="加载失败，请重试">
              <template v-for="(item2,index2) in item.list">
                <order-list :item="item2" :index="index2" @cancel="onCancel"></order-list>
              </template>
            </van-list>
          </template>
          <van-empty v-else style="height:500px;" description="没有数据哦" />
        </van-pull-refresh>
      </van-tab>
    </van-tabs>

    <div class="text-center main-text-color font-md font-weight mt-2">为你推荐</div>
    <van-divider>买的人还买了</van-divider>
    <!-- 为你推荐 -->
    <div class="row j-sb bg-white">
      <common-list v-for="(item,index) in hotList" :key="index" :item="item" :index="index"></common-list>
    </div>

  </div>
</template>

<script>
import commonList from "@/components/common/common-list.vue"
import orderList from "@/components/order/order-list.vue"
import { skuList } from '@/api/product'
import { getOrderList } from '@/api/order'
export default {
  components: {
    commonList,
    orderList
  },
  data() {
    return {
      hotList: [],
      list: [],
      tabBars: [
        { name: '全部', status: 0, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
        { name: '待付款', status: 1, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
        { name: '待发货', status: 2, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
        { name: '已发货', status: 3, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
        { name: '已收货', status: 4, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
        { name: '已完成', status: 5, page: 1, empty: false, loading: false, finished: false, refreshing: false, list: [] },
      ],
      active: 0,
    }
  },
  activated() {
    this.active = this.$route.query.status ? parseInt(this.$route.query.status) : 0
    this.initData()
  },
  methods: {
    back() {
      this.$router.back()
    },
    // 获取数据
    initData() {
      // 获取热门推荐
      skuList(0, 1).then(res => {
        this.hotList = res
      })
      this.onLoad()
    },
    onChange(index) {
      this.active = index
      this.onLoad()
    },
    onLoad() {
      this.getList()
    },
    onRefresh() {
      this.tabBars[this.active].finished = false
      this.tabBars[this.active].loading = true
      this.tabBars[this.active].empty = false
      this.tabBars[this.active].page = 1
      this.tabBars[this.active].list = []
      this.getList()
    },
    getList() {
      const index = this.active
      const status = this.tabBars[index].status
      const p = this.tabBars[index].page
      if (this.tabBars[index].finished) {
        return
      }
      getOrderList(status, p).then(res => {
        this.tabBars[index].refreshing = false
        this.tabBars[index].loading = false
        if (this.tabBars[index].list.length == 0 && res.length == 0) {
          this.tabBars[index].empty = true
        } else {
          this.tabBars[index].list = [...this.tabBars[index].list, ...res]
          this.tabBars[index].page++
        }
        if (res.length < 20) {
          this.tabBars[index].finished = true
        }
      }).catch(() => {
        this.tabBars[index].finished = true
      })
    },
    onCancel(index) {
      this.tabBars[this.active].list.splice(index, 1)
    }
  }
}
</script>

<style>
</style>

<template>
  <div>
    <van-sticky>
      <van-search show-action placeholder="请输入搜索关键词" @click="openSearch">
        <template #action>
          <van-icon name="gift-o" size="24" color="#ee0a24" @click="open('/seckill')" />
        </template>
        <template #left>
          <span class="iconfont icon-xiaoxi font-lg" @click="openNotice"></span>
        </template>
      </van-search>
    </van-sticky>
    <van-tabs v-model="active" @change="onChangeTab" color="#FD6801" title-active-color="#FD6801" offset-top="45" sticky animated swipeable>
      <van-tab v-for="(item,index) in tabBars" :title="item.name">
        <template v-if="item.list.length > 0">
          <template v-for="(list,listIndex) in item.list">
            <!-- 轮播图组件 -->
            <swiper-image v-if="list.type === 1" :images="list.images.list" />

            <!-- 首页图标 -->
            <index-nav v-else-if="list.type === 2" :resdata="list.navs.list" />

            <!-- 三图广告 -->
            <three-adv v-else-if="list.type === 3" :resdata="list.images.list" />

            <!-- 大图广告位 -->
            <card v-else-if="list.type === 4" :headTitle="list.ads.title" :bodyCover="list.ads.cover" />

            <!-- 公共列表组件-->
            <template v-else-if="list.type === 5">
              <van-list v-model="goodsList[index].loading" :finished="goodsList[index].finished" finished-text="没有更多了"
                        @load="onLoad(list.product.router)" error-text="加载失败，请重试">
                <div class="row j-sb">
                  <common-list v-for="(item2,index2) in goodsList[index].list" :item="item2" :index="index2" />
                </div>
              </van-list>
            </template>
          </template>
        </template>

        <van-empty v-else style="height:500px;" description="没有数据哦" />

      </van-tab>
    </van-tabs>

  </div>
</template>

<script>
import base from '@/mixin/base.js';
import swiperImage from "@/components/index/swiper-image.vue";
import indexNav from "@/components/index/index-nav.vue";
import threeAdv from "@/components/index/three-adv.vue";
import card from "@/components/common/card.vue";
import commonList from "@/components/common/common-list.vue";
import { homeData, homeCatData } from '@/api/app.js';
import api from '@/api/index'
export default {
  mixins: [base],
  components: {
    swiperImage,
    indexNav,
    threeAdv,
    card,
    commonList
  },
  data() {
    return {
      active: 0,
      tabBars: [],
      goodsList: [],
    }
  },
  mounted() {
    this.initData()
  },
  methods: {
    openSearch() {
      this.$router.push({ path: '/search' })
    },
    openNotice() {
      this.$router.push({ path: '/notice' })
    },
    // 初始化事件
    initData() {
      // 获取顶部选项卡
      homeData().then((res) => {
        this.tabBars = res
        this.tabBars.forEach(() => {
          this.goodsList.push({
            page: 1,
            list: [],
            loading: false,
            finished: false,
          })
        })
      })
    },
    onLoad(router) {
      this.getList(router)
    },
    getList(router) {
      console.log('getlist', this.active, this.tabBars[this.active])
      const index = this.active
      const page = this.goodsList[index].page || 1
      api.get(router + '/cat/' + this.tabBars[this.active].id + '/p/' + page).then((res) => {
        this.goodsList[index].loading = false
        if (res.length > 0) {
          this.goodsList[index].page++
          this.goodsList[index].list = [...this.goodsList[index].list, ...res];
        }
        // 数据全部加载完成
        if (res.length < 20) {
          this.goodsList[index].finished = true;
        }
      }).catch((err) => {
        this.goodsList[index].finished = true;
      })
    },
    // 监听滑动列表
    onChangeTab(name, item) {
      const cat = this.tabBars[name].id;
      if (this.tabBars[name].list.length > 0) {
        return
      }
      homeCatData(cat).then(res => {
        this.tabBars[name].list = res
      })
    },
  }
}
</script>

<style>
</style>

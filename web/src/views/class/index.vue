<template>
  <div class="d-flex flex-column">
    <van-search show-action placeholder="请输入搜索关键词" @click="openSearch">
      <template #action>
        <span class="iconfont icon-xiaoxi" @click="openNotice"></span>
      </template>
    </van-search>

    <div class="d-flex" style="height:560px;overflow:auto;">
      <van-sidebar v-model="activeIndex" @change="onChange">
        <van-sidebar-item v-for="(item,index) in cats" :title="item.name" />
      </van-sidebar>
      <div style="width:80%;overflow:auto;" id="cat-list">
        <template v-for="(item,index) in cats">
          <div :id="'child-'+item.id">
            <template v-if="item.child.length > 0" v-for="(item2,index2) in item.child">
              <div class="m-1">
                <span class="d-block font-sm main-text-color">{{item2.name}}</span>
              </div>
              <div class="row">
                <template v-if="item2.child.length > 0" v-for="(item3,index3) in item2.child">
                  <span class="d-block m-1" @click="openGoodsList(item3.id,item3.name)">{{item3.name}}</span>
                </template>
              </div>
            </template>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
import { categoryTree } from '@/api/product'
export default {
  data() {
    return {
      // 当前选中的分类
      activeIndex: 0,
      cats: []
    }
  },
  mounted() {
    this.initData()
  },
  methods: {
    openNotice() {
      this.$router.push({ path: '/notice' })
    },
    openSearch() {
      this.$router.push({ path: '/search' })
    },
    openGoodsList(cat, title) {
      this.$router.push({ path: '/product_list', query: { cat, title } })
    },
    onChange(index) {
      console.log('index', index)
      this.goAnchor('#child-' + this.cats[index].id)
    },
    initData() {
      categoryTree().then(res => {
        this.cats = res
      })
    },
    // 打开详情页
    openDetail(item) {
      this.$router.push({ path: '/goods_detail', query: { id: item.goods_id } })
    },
    clickRight() {
      this.$router.push({ path: '/msg_list' })
    },
    goAnchor(selector) {
      var anchor = this.$el.querySelector(selector)
      this.$el.querySelector("#cat-list").scrollTop = anchor.offsetTop - 60
    }
  }
}
</script>

<style>
</style>

<template>
  <div>
    <van-search v-model="keyword" show-action placeholder="请输入搜索关键词" @search="onSearch" @cancel="onCancel" />

    <template v-for="(list,index) in data">
      <!-- 大图广告位 -->
      <card v-if="list.type === 4" :headTitle="list.ads.title" :bodyCover="list.ads.cover" />
    </template>
    <!-- 多色按钮 -->
    <div class="px-1 mb-2">
      <van-tag class="p-1" style="margin:5px;" :color="item.color" size="medium" v-for="(item,index) in hot" @click="quickSearch(item.name)">
        {{item.name}}
      </van-tag>
    </div>

    <template v-if="historyList.length > 0">
      <!-- 搜索记录 -->
      <card headTitle="搜索记录">
        <div slot="right" class="font text-primary" @click="clearHistory">清除搜索记录</div>
        <van-cell :title="item" v-for="(item,index) in historyList" @click="quickSearch(item)" />
      </card>
    </template>

  </div>
</template>

<script>
import card from "@/components/common/card.vue"
import { Dialog, Toast } from 'vant';
import { getStorage, setStorage, removeStorage } from '@/utils/index'
import { searchData } from '@/api/app'
export default {
  components: {
    card
  },
  data() {
    return {
      historyList: [],
      keyword: "",
      hot: [],
      data: [],
    }
  },
  mounted() {
    this.initData()
  },
  methods: {
    onCancel() {
      this.$router.back()
    },
    initData() {
      // 加载历史记录
      let history = getStorage('searchHistory')
      this.historyList = history ? JSON.parse(history) : []
      searchData().then((res) => {
        console.log(res)
        this.hot = res.hot.map(name => {
          let color = '#' + Math.floor(Math.random() * 16777215).toString(16);
          return { name, color }
        })
        this.data = res.data
      })
    },
    quickSearch(name) {
      this.keyword = name;
      this.onSearch()
    },
    onSearch() {
      if (this.keyword === '') {
        return Toast('关键词不能为空')
      }
      this.addHistory()
      this.$router.push({ path: '/product_list', query: { title: this.keyword, type: 'search' } })
    },
    // 加入搜索记录
    addHistory() {
      let index = this.historyList.indexOf(this.keyword)
      if (index === -1) {
        this.historyList.unshift(this.keyword)
      } else {
        this.historyList.unshift(this.historyList.splice(index, 1)[0])
      }
      // 移除最后一条
      if (this.historyList.length > 6) {
        this.historyList.splice(this.historyList.length - 1, 1)
      }
      setStorage('searchHistory', JSON.stringify(this.historyList))
    },
    // 清除搜索记录
    clearHistory() {
      Dialog.confirm({
        title: '提示',
        message: '是否要清除搜索历史？',
      })
        .then(() => {
          this.historyList = []
          removeStorage('searchHistory')
        })
        .catch(() => {
          // on cancel
        });
    }
  }
}
</script>

<style>
</style>

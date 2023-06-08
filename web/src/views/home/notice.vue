<template>
  <div>
    <van-nav-bar title="公告" left-arrow @click-left="back" />
    <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad" error-text="加载失败，请重试">
      <van-cell :title="item.title" is-link v-for="(item,index) in list" :key="index" @click="open(item)">
        <!-- 使用 title 插槽来自定义标题 -->
        <template #label>
          <div class="font text-light-muted text-inline-3">{{item.content}}</div>
          <div class="font text-light-muted">{{item.created_at|formatTime}}</div>
        </template>
      </van-cell>
    </van-list>
    <van-overlay :show="show" @click="show = false">
      <div class="wrapper">
        <div class="bg-white rounded p-1 m-3">
          <van-cell :title="current_title">
            <!-- 使用 title 插槽来自定义标题 -->
            <template #label>
              <div class="font text-light-muted">{{current_content}}</div>
            </template>
          </van-cell>
        </div>
      </div>
    </van-overlay>
  </div>
</template>

<script>
import { noticeList } from '@/api/app'
export default {
  data() {
    return {
      show: false,
      current_title: '',
      current_content: '',
      page: 1,
      loading: false,
      finished: false,
      list: []
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
    onLoad() {
      this.getData()
    },
    getData() {
      noticeList(this.page).then(res => {
        console.log('res', res)
        this.list = this.page == 1 ? res : [...this.list, ...res]
        this.page++
        this.loading = false
        // 数据全部加载完成
        if (res.length < 20) {
          this.finished = true;
        }
      }).catch(err => {
        this.finished = true;
        if (this.page > 1) {
          this.page--
        }
      })
    },
    open(item) {
      this.show = true;
      this.current_title = item.title;
      this.current_content = item.content;
    }
  }
}
</script>

<style>
.wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>

<template>
  <div>
    <van-nav-bar title="商品评价" left-arrow @click-left="back" />
    <div class="p-1">
      <div class="d-flex a-center py-1 border-bottom border-light-secondary">
        <span class="text-muted">用户评价({{total}})</span>
        <span class="main-text-color ml-auto mr-1">{{ (good_rate * 100) + '%' }}</span>
        <span class="text-muted">满意</span>
      </div>
      <div class="d-flex flex-wrap pt-1">
        <div class="px-1 py border rounded mr-1 mb-1 cate-item font" v-for="(item,index) in cateList" :key="index"
             :class="cateIndex === index ? 'active' : ''" @click="cateChange(index)">
          {{item.name}}
        </div>
      </div>
    </div>

    <div class="p-2 d-flex a-start border-bottom border-light-secondary" v-for="(item,index) in list" :key="index">
      <image :src="item.user.avatar" mode="widthFix" style="width: 45px;height: 45px;" class="rounded flex-shrink"></image>
      <div class="pl-2 flex-1">
        <div class="d-flex a-center">
          <span class="font-md text-primary font-weight">{{item.user.nickname}}</span>
          <div class="iconfont icon-service main-text-color ml-auto">
            <span class="pl-1 font">{{item.rating | rating}}</span>
          </div>
        </div>
        <div class="line-h-md font-md">
          {{item.review.data}}
        </div>
        <div class="row" v-if="item.review.image.length > 0">
          <div class="span24-8 px pb" v-for="(img,imgIndex) in item.review.image" :key="imgIndex">
            <image :src="img" mode="widthFix" style="height: 50px;"></image>
          </div>
        </div>
        <div class="d-flex a-center">
          <span class="text-light-muted">{{item.review_time|formatTime}}</span>
          <!-- <div class="d-flex text-light-muted ml-auto a-center mr-2">
						{{item.goods_num}} <span class="iconfont icon-dianzan text-muted ml"></span>
					</div> -->
          <!-- <div class="d-flex text-light-muted a-center">
						<span class="iconfont icon-pinglun text-muted ml"></span>
					</div> -->
        </div>
        <!-- 客服回复 -->
        <div class="bg-light-secondary rounded p-2 d-flex flex-wrap a-center mt-1" v-for="(item2,index2) in item.extra" :key="index2">
          <template v-if="!item2.isuser">
            <span class="main-text-color">官方回复：</span>
            {{item2.data}}
            <!-- <div class="iconfont icon-dianzan text-light-muted line-h-md font ml-2">
							<span class="text-muted pl-1">赞客服 {{item2.good_num}}</span>
						</div> -->
          </template>
        </div>
      </div>
    </div>
    <!-- 上拉加载 -->
    <div class="d-flex a-center j-center text-light-muted font-md py-3">
      {{loadtext}}
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: 0,
      cateIndex: 0,
      cateList: [
        { name: "全部", value: "" },
        { name: "好评", value: "/good" },
        { name: "中评", value: "/middle" },
        { name: "差评", value: "/bad" },
      ],
      total: 0,
      good_rate: 0,
      list: [],
      page: 1,
      // 1.上拉加载更多 2.加载中... 3.没有更多了
      loadtext: "上拉加载更多"
    }
  },
  onLoad(e) {
    this.id = e.id
    this.__init()
  },
  onReachBottom() {
    if (this.loadtext !== '上拉加载更多') return;

    this.page++
    this.loadtext = '加载中...'
    this.getData()
  },
  filters: {
    rating(value) {
      if (value === 3) {
        return '中评'
      } else if (value > 3) {
        return '好评'
      } else {
        return '差评'
      }
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
    // 初始化
    __init() {
      this.getData()
    },
    // 加载数据
    getData(callback = false, refresh = false) {
      let value = this.cateList[this.cateIndex].value
      this.$H.get(`/goods/${this.id}/comments${value}?page=${this.page}`).then(res => {
        this.total = res.total
        this.good_rate = res.good_rate
        this.list = !refresh ? [...this.list, ...res.list] : [...res.list]

        this.loadtext = res.list.length < 10 ? '没有更多数据了' : '上拉加载更多'
        if (typeof callback === 'function') {
          callback(true)
        }
      })
    },
    cateChange(index) {
      this.cateIndex = index
      this.getData(false, true)
    }
  }
}
</script>

<style>
.cate-item {
  background: #ffebec;
  color: #7b6d6c;
  border-color: #f4e0e1;
}
.cate-item.active {
  background: #ff6b01 !important;
  color: #ffffff !important;
  border-color: #de7232 !important;
}
</style>

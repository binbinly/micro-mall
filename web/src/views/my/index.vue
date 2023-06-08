<template>
  <div>
    <!-- 头部 -->
    <div class="position-relative d-flex a-center animated fadeIn faster" style="height: 160px;">
      <!-- 消息列表 -->
      <div class="iconfont icon-xiaoxi position-absolute text-white" @click="open('/notice')"
           style="font-size: 25px;top: 38px;right: 10px;z-index: 100;">
      </div>

      <van-image :src="require('@/assets/images/bg.jpg')" width="100%" fit="cover" />
      <div class="d-flex a-center position-absolute left-0 right-0" style="bottom: 25px;">

        <van-image :src="token ? user.avatar : ''|formatAvatar" height="70" fit="cover" round class="rounded-circle border-light ml-2" />
        <div v-if="user" class="ml-2 text-white font-md">
          <div>{{user.name}}</div>
          <div>{{user.phone}}</div>
        </div>
        <div class="ml-2 text-white font-md" @click="open('/login')" v-else>登录/注册</div>
      </div>
    </div>
    <!-- 图标分类 -->
    <card class="bg-white">
      <div slot="title" class="d-flex a-center j-sb w-100">
        <span class="font-md font-weight">我的订单</span>
        <div class="text-secondary font ml-auto" @click="open('/order')">
          全部订单 <span class="iconfont icon-you font"></span>
        </div>
      </div>
      <div class="d-flex a-center">
        <div class="flex-1 d-flex flex-column a-center j-center py-3" v-for="(item,index) in orders" :key="index" @click="openOrder(item)">
          <span class="iconfont font-lg line-h" :class="item.icon"></span>
          <span style="margin-top:2px;">{{item.name}}</span>
        </div>
      </div>
    </card>
    <van-divider />
    <van-cell title="我的钱包" is-link to="/user_wallet">
      <template #icon>
        <span class="iconfont icon-VIP" style="color:#FDBF2E;"></span>
      </template>
    </van-cell>
    <van-cell title="我的合约" is-link to="/user_contract">
      <template #icon>
        <span class="iconfont icon-huangguan" style="color:#FCBE2D;"></span>
      </template>
    </van-cell>
    <van-cell title="服务中心" is-link @click="event">
      <template #icon>
        <span class="iconfont icon-service" style="color:#FA6C5E;"></span>
      </template>
    </van-cell>
    <van-cell title="我的账单" is-link @click="event">
      <template #icon>
        <span class="iconfont icon-home" style="color:#FE8B42;"></span>
      </template>
    </van-cell>
    <van-cell title="更多功能" is-link @click="event">
      <template #icon>
        <span class="iconfont icon-gengduo" style="color:#9ED45A;"></span>
      </template>
    </van-cell>
    <van-divider />
    <van-cell title="更多设置" is-link to="/setting">
      <template #icon>
        <span class="iconfont icon-icon_set_up" style="color:#808C98;"></span>
      </template>
    </van-cell>
  </div>
</template>

<script>
import card from "@/components/common/card.vue"
import { mapGetters } from 'vuex';
import { Toast } from 'vant'
export default {
  components: {
    card,
  },
  data() {
    return {
      orders: [{
        name: "待付款",
        icon: "icon-wallet_icon",
        status: 1
      }, {
        name: "待收货",
        icon: "icon-daishouhuo",
        status: 2
      }, {
        name: "待评价",
        icon: "icon-pinglun",
        status: 4
      }, {
        name: "已完成",
        icon: "icon-buoumaotubiao46",
        status: 5
      }]
    }
  },
  computed: {
    ...mapGetters(['token', 'user'])
  },
  methods: {
    event() {
      Toast('待开发')
    },
    open(path) {
      this.$router.push({ path })
    },
    openOrder(item) {
      this.$router.push({ path: '/order', query: { status: item.status } })
    }
  }
}
</script>

<style>
.iconfont {
  padding-right: 5px;
}
</style>

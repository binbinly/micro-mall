<template>
  <div class="app-container">
    <div class="layout-content">
      <keep-alive v-if="$route.meta.keepAlive">
        <router-view></router-view>
      </keep-alive>
      <router-view v-else></router-view>
    </div>
    <div class="layout-footer" v-if="$route.meta.tabbar">
      <TabBar :data="tabbars" @change="handleChange" />
    </div>
  </div>
</template>

<script>
import TabBar from '@/components/TabBar'
export default {
  name: 'AppLayout',
  data() {
    return {
      tabbars: [
        {
          title: '首页',
          to: {
            name: 'Home'
          },
          icon: require('@/assets/tabbar/index.png'),
          iconS: require('@/assets/tabbar/index-selected.png')
        },
        {
          title: '分类',
          to: {
            name: 'Class'
          },
          icon: require('@/assets/tabbar/class.png'),
          iconS: require('@/assets/tabbar/class-selected.png')
        },
        {
          title: '购物车',
          to: {
            name: 'Cart'
          },
          icon: require('@/assets/tabbar/cart.png'),
          iconS: require('@/assets/tabbar/cart-selected.png')
        },
        {
          title: '我的',
          to: {
            name: 'My'
          },
          icon: require('@/assets/tabbar/my.png'),
          iconS: require('@/assets/tabbar/my-selected.png')
        }
      ]
    }
  },
  components: {
    TabBar
  },
  created() {
    // 初始化登录状态
    this.$store.commit('initUser')
    if (this.$store.getters.user) {
      // 初始化我的购物车
      this.$store.dispatch('initCart')
      // 初始化web3对象
      this.$store.commit('initWeb3')
    }

  },
  methods: {
    handleChange(v) {
      console.log('tab value:', v)
    }
  }
}
</script>

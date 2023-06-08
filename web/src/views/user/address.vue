<template>
  <div>
    <van-nav-bar title="收货地址" left-text="返回" left-arrow @click-left="back">
      <template #right>
        <van-icon name="plus" size="20" @click="onAdd" />
      </template>
    </van-nav-bar>

    <div v-for="(item,index) in list">
      <van-swipe-cell>
        <van-cell value="内容" is-link center clickable value-class="none" @click.stop="onChoose(item)">
          <!-- 使用 title 插槽来自定义标题 -->
          <template #title>
            <div>
              <span class="custom-title font-md">{{item.name}}&nbsp;&nbsp;{{item.phone}}</span>
              <van-tag v-if="item.is_default" class="ml-1" type="danger" round>默认</van-tag>
            </div>
            <div class="font">{{item.province}}{{item.city}}{{item.county}}&nbsp;{{item.detail}}</div>
          </template>
          <template #right-icon>
            <van-icon name="edit" size="20" @click.stop="onEdit(index)" />
          </template>
        </van-cell>
        <template #right>
          <van-button square type="danger" @click="onDelete(index)" text="删除" style="height:100%" />
        </template>
      </van-swipe-cell>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex"
import ev from '@/utils/event'
import { Dialog, Toast } from 'vant'
export default {
  data() {
    return {
      isChoose: false,
    }
  },
  computed: {
    ...mapState({
      list: state => state.address.list
    }),
  },
  created() {
    // 初始化收货地址
    this.$store.commit('initAddress')
  },
  activated() {
    this.isChoose = false
    if (this.$route.query.type === 'choose') {
      this.isChoose = true
    }
  },
  methods: {
    ...mapActions(['delAddress']),
    back() {
      this.$router.back()
    },
    onAdd() {
      this.$router.push({ path: '/address_edit' })
    },
    onEdit(index) {
      this.$router.push({ path: '/address_edit', query: { index } })
    },
    onDelete(index) {
      Dialog.confirm({
        title: '提示',
        message: '是否退出当前账号？',
      })
        .then(() => {
          this.delAddress({ index: index, id: this.list[index].id }).then(() => {
            Toast('删除成功')
          })
        })
        .catch(() => {
          // on cancel
        });
    },
    // 选择收货地址
    onChoose(item) {
      if (this.isChoose) {
        // 通知订单提交页修改收货地址
        ev.$emit('chooseAddr', item)
        this.back()
      }
    }
  }
}
</script>

<style>
.none {
  display: none;
}
</style>

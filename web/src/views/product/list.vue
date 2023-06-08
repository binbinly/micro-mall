<template>
  <div>
    <van-nav-bar :title="navTitle" left-arrow fixed placeholder @click-left="back" />

    <!-- 排序|筛选 -->
    <div class="d-flex border-bottom a-center position-fixed left-0 right-0 bg-white" style="height: 45px;z-index: 100;">
      <div class="flex-1 d-flex a-center j-center font-sm" v-for="(item,index) in screen.list" @click="changeScreen(index)">
        <span :class="screen.currentIndex === index ? 'main-text-color' : 'text-muted'">{{item.name}}</span>
        <div>
          <div class="iconfont icon-paixu-shengxu line-h0" :class="item.status==='asc'?'main-text-color':'text-light-muted'">
          </div>
          <div class="iconfont icon-paixu-jiangxu line-h0" :class="item.status==='desc'?'main-text-color':'text-light-muted'">
          </div>
        </div>
      </div>
      <div class="flex-1 d-flex a-center j-center main-text-color font-sm" @click="show = true">
        筛选
      </div>
    </div>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="padding-top:60px;">
      <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad" error-text="加载失败，请重试">
        <div class="row j-sb bg-white">
          <common-list v-for="(item,index) in list" :item="item" :index="index"></common-list>
        </div>
      </van-list>
    </van-pull-refresh>

    <van-popup v-model="show" position="right" :style="{ height: '100%' }">
      <div style="width:300px;height: 620px;overflow: auto;">
        <card headTitle="价格范围" :headBorderBottom="false" :headTitleWeight="false">
          <!-- 单选按钮组 -->
          <zcm-radio-group :label="label" :selected.sync='label.selected'></zcm-radio-group>
        </card>
        <card headTitle="分类" :headBorderBottom="false" :headTitleWeight="false">
          <!-- 单选按钮组 -->
          <zcm-radio-group :label="catList" :selected.sync='catList.selected'></zcm-radio-group>
        </card>
        <card headTitle="品牌" :headBorderBottom="false" :headTitleWeight="false">
          <!-- 单选按钮组 -->
          <zcm-radio-group :label="brandList" :selected.sync='brandList.selected'></zcm-radio-group>
        </card>
        <card v-for="(item,index) in attrs" :headTitle="item.name" :headBorderBottom="false" :headTitleWeight="false">
          <!-- 单选按钮组 -->
          <zcm-attr-group :label="item" :selected.sync='item.selected'>
          </zcm-attr-group>
        </card>
      </div>
      <!-- 按钮 -->
      <div class="d-flex position-fixed bottom-0 right-0 w-100 border-top border-light-secondary">
        <van-button class="flex-1" color="#FD6801" type="primary" @click="confirm">确定</van-button>
        <van-button class="flex-1" type="default" @click="reset">重置</van-button>
      </div>
    </van-popup>

  </div>
</template>

<script>
import zcmRadioGroup from "@/components/common/radio-group.vue"
import zcmAttrGroup from "@/components/common/attr-group.vue"
import card from "@/components/common/card.vue"
import commonList from "@/components/common/common-list.vue"
import { skuSearch } from '@/api/product'
export default {
  components: {
    commonList,
    card,
    zcmRadioGroup,
    zcmAttrGroup
  },
  data() {
    return {
      show: false,
      navTitle: '商品列表',
      cat: 0,
      keyword: '',
      list: [],
      page: 1,
      loading: false,
      finished: false,
      refreshing: false,
      catList: {
        selected: 0,
        list: []
      },
      attrs: [],
      brandList: {
        selected: 0,
        list: []
      },
      select: {
        brand: 0,
        attrs: {}
      },
      screen: {
        currentIndex: 0,
        list: [
          { name: "综合", status: 'desc', key: "0" },
          { name: "销量", status: '', key: "1" },
          { name: "价格", status: '', key: "2" },
        ]
      },
      label: {
        selected: 0,
        list: [
          { name: "不限", value: '' },
          { name: "0-50", value: "0_50" },
          { name: "50-100", value: "50_100" },
          { name: "100-500", value: "100_500" },
          { name: "500-1000", value: "500_1000" },
          { name: ">1000", value: "1000_0" },
        ]
      },
    }
  },
  filters: {
    toString(value) {
      return value.toString();
    }
  },
  mounted() {
    this.navTitle = this.$route.query.title || '商品列表';
    if (this.$route.query.type === 'search') {
      this.keyword = this.navTitle
    }
    if (this.$route.query.cat) {
      this.cat = parseInt(this.$route.query.cat)
    }
  },
  methods: {
    back() {
      this.$router.back()
    },
    reset() {
      this.label.selected = 0
      this.show = false
      this.select = { brand: 0, attrs: {} }
      this.catList.selected = 0
      this.brandList.selected = ''
      this.attrs = []
      this.onRefresh()
    },
    confirm() {
      this.show = false
      this.onRefresh()
    },
    changeScreen(index) {
      // 判断当前点击是否已经是激活状态
      let oldIndex = this.screen.currentIndex
      let oldItem = this.screen.list[oldIndex]
      if (oldIndex === index) {
        oldItem.status = oldItem.status === 'asc' ? 'desc' : 'asc'
        // 加载数据
        this.onRefresh()
        return
      }
      let newIitem = this.screen.list[index]
      // 移除旧激活状态
      oldItem.status = ''
      this.screen.currentIndex = index
      // 增加新激活状态
      newIitem.status = 'desc'

      // 加载数据
      this.onRefresh()
    },
    onLoad() {
      this.getList()
    },
    onRefresh() {
      this.finished = false
      this.loading = true
      this.page = 1
      this.list = []
      this.onLoad()
    },
    getList() {
      // TODO 客户端暂未做多选
      const field = this.screen.list[this.screen.currentIndex].key
      const order = this.screen.list[this.screen.currentIndex].status
      const price = this.label.list[this.label.selected].value
      const cat = this.catList.list[this.catList.selected] ? this.catList.list[this.catList.selected].id : 0
      const brand = this.brandList.list[this.brandList.selected] ? this.brandList.list[this.brandList.selected].id : 0

      let attrs = []
      this.attrs.forEach(item => {
        if (item.selected != '') {
          //设置已选属性
          this.select.attrs[item.id] = item.selected
          attrs.push({ id: item.id, values: [item.selected] })
        }
      })
      // console.log('search', field, order, price, cat, brand, attrs)
      let params = { page: this.page }
      if (attrs.length > 0) {
        params['attrs'] = attrs
      }
      if (field) {
        params['field'] = parseInt(field)
      }
      params['order'] = order == 'desc' ? 1 : 0
      if (price != "") {
        const [price_s, price_e] = price.split('_')
        params['price_s'] = price_s ? parseInt(price_s) : 0
        params['price_e'] = price_e ? parseInt(price_e) : 0
      }
      if (this.cat > 0) {
        params['cat_id'] = this.cat
      }
      if (cat > 0) {
        params['cat_id'] = parseInt(cat)
      }
      if (brand > 0) {//品牌多选
        // 设置已选品牌
        this.select.brand = brand
        params['brand_id'] = [parseInt(brand)]
      }
      if (this.keyword != "") {
        params['keyword'] = this.keyword
      }
      skuSearch(params).then(res => {
        if (this.refreshing) {
          this.refreshing = false
        }
        this.page++
        this.loading = false
        this.list = [...this.list, ...res.data]
        if (res.data.length < 20) {
          this.finished = true
        }
        if (res.brands.length > 0) {
          res.brands.unshift({ id: 0, 'name': '不限' })
        }
        if (res.cats.length > 0) {
          res.cats.unshift({ id: 0, 'name': '不限' })
        }
        //定位所选分类
        if (this.cat > 0) {
          for (let i = 0; i < res.cats.length; i++) {
            if (res.cats[i].id == this.cat) {
              this.catList.selected = i
              break
            }
          }
        }
        this.catList.list = res.cats
        //定位已选品牌
        if (this.select.brand > 0) {
          for (let i = 0; i < res.brands.length; i++) {
            if (res.brands[i].id == this.select.brand) {
              this.brandList.selected = i
              break
            }
          }
        }
        this.brandList.list = res.brands
        //定位已选属性
        this.attrs = res.attrs.map(item => {
          item['selected'] = ''
          if (this.select.attrs[item.id]) {
            item['selected'] = this.select.attrs[item.id]
          }
          return item
        })
      }).catch(err => {
        console.log('err', err)
        this.finished = true
      })
    },
  }
}
</script>

<style>
</style>

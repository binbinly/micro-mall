<template>
  <div class="animated fadeIn faster" style="padding-bottom:75px;">

    <van-nav-bar title="购物车" fixed placeholder>
      <template #right>
        <van-icon name="delete-o" size="20" @click="doEmptyCart" />
      </template>
    </van-nav-bar>

    <!-- 购物车为空 -->
    <div class="py-5 d-flex a-center j-center bg-white" v-if="cartIsEmpty">
      <div class="iconfont icon-gouwuche text-light-muted" style="font-size: 25px;"></div>
      <span class="text-light-muted mx-1">购物车还是为空</span>
      <van-button type="default" to="/product_list">去逛逛</van-button>
    </div>

    <!-- 购物车商品列表 -->
    <div class="bg-white" v-else>
      <van-checkbox-group v-model="result" @change="onChange" ref="checkboxGroup">
        <!-- 列表 -->
        <div v-for="(item,index) in list" :key="index">
          <van-swipe-cell>
            <div class="d-flex a-center py-1 border-bottom border-light-secondary">
              <label class="radio d-flex a-center j-center flex-shrink" style="width: 40px;height: 40px;">
                <van-checkbox :name="index" checked-color="#FD6801"></van-checkbox>
              </label>

              <van-image class="border border-light-secondary rounded flex-shrink" :src="item.cover" fit="cover" width="75" />

              <div class="flex-1 d-flex flex-column pl-1" style="width:240px;">
                <div class="text-dark font-sm text-inline" @click="openGoods(item)">{{item.title}}</div>
                <!-- 规格属性 -->
                <div class="d-flex text-light-muted" style="padding:5px 0;" @click.stop="showPopup(index,item)">
                  {{item.sku_attr}}
                </div>
                <div class="mt-auto d-flex j-sb">
                  <price :text="item.price" class="a-center" />
                  <van-stepper v-model="item.num" @plus="onChangeNum(index,item)" @minus="onChangeNum(index,item)" @blur="onChangeNum(index,item)"
                               async-change class="mr-1" />
                </div>
              </div>
            </div>
            <template #right>
              <van-button square type="danger" text="删除" class="swipe-button" @click="doDelCart(index)" />
            </template>
          </van-swipe-cell>
        </div>
      </van-checkbox-group>
    </div>

    <div class="text-center main-text-color font-md font-weight mt-2">为你推荐</div>
    <van-divider>买的人还买了</van-divider>
    <!-- 为你推荐 -->
    <div class="row j-sb bg-white">
      <common-list v-for="(item,index) in hotList" :key="index" :item="item" :index="index"></common-list>
    </div>
    <!-- 合计 -->
    <van-submit-bar :price="total" button-text="结算" @submit="orderConfirm" button-color="#FD6801" style="bottom:50px;">
      <van-checkbox v-model="checked" @change="onChangeAll" checked-color="#FD6801">全选</van-checkbox>
    </van-submit-bar>
    <!-- SKU -->
    <van-sku v-model="skuShow" :sku="sku" :goods="goods" :goods-id="sku_id" :quota="0" :quota-used="0" :initial-sku="initialSku" ref="cartSku">
      <!-- 自定义 sku actions -->
      <template #sku-actions="props">
        <div class="van-sku-actions">
          <van-button square size="large" type="warning" @click="onSkuClicked">
            确定
          </van-button>
        </div>
      </template>
    </van-sku>
  </div>
</template>

<script>
import price from "@/components/common/price.vue"
import { mapState, mapGetters, mapActions } from "vuex"
import commonList from "@/components/common/common-list.vue"
import { Toast } from 'vant'
import { skuList, skuDetailSaleAttr } from '@/api/product'
export default {
  components: {
    price,
    commonList
  },
  data() {
    return {
      result: [],
      total: 0,
      checked: false,
      isedit: false,
      hotList: [],
      goods: { picture: '' },
      sku: {
        tree: [],
        list: [],
        price: 0,
        stock_num: 0,
        collection_id: 0,
        none_sku: true
      },
      initialSku: {},
      sku_id: 0,
      skuShow: false,
      current_index: 0, //当前操作索引
    }
  },
  computed: {
    ...mapState({
      list: state => state.cart.list,
    }),
    ...mapGetters([
      'cartIsEmpty'
    ])
  },
  mounted() {
    this.initData()
  },
  methods: {
    ...mapActions([
      'doEditCart',
      'doDelCart',
      'doEmptyCart',
      'doEditCart'
    ]),
    onChangeAll(checked) {//全选
      if (checked) {
        return this.$refs.checkboxGroup.toggleAll(true);
      }
      this.$refs.checkboxGroup.toggleAll(false);
    },
    onChange(names) {//购物车列表选中变更监听
      this.total = 0
      names.forEach(index => {
        this.total += this.list[index].price * this.list[index].num * 100
      })
    },
    onChangeNum(index, item) {//增减商品数量
      if (this.result.indexOf(index) !== -1) {//选中，计算总计
        this.total = 0
        this.result.forEach(index => {
          this.total += this.list[index].price * this.list[index].num * 100
        })
      }
      this.doEditCart({ index, item, num: item.num })
    },
    onSkuClicked() {//修改sku
      const data = this.$refs.cartSku.getSkuData()
      if (!data.selectedSkuComb) {
        return Toast('请选择商品')
      }
      const index = this.current_index;
      this.skuShow = false;
      this.current_index = 0;
      let item = this.list[index];
      const sku_id = parseInt(data.selectedSkuComb.id)
      if (sku_id == item.sku_id && data.selectedNum == item.num) {//数据未修改
        return
      }
      let params = { index, item, num: data.selectedNum }
      if (sku_id != item.sku_id) {
        params['sku_id'] = sku_id
      }
      this.doEditCart(params)
    },
    openGoods(item) {
      this.$router.push({ path: "/goods_detail", query: { id: item.sku_id } })
    },
    // 订单结算
    orderConfirm() {
      if (this.result.length === 0) {
        return Toast('请选择要结算的商品')
      }
      this.$router.push({ path: "/order_confirm", query: { ids: this.result.join(',') } })
    },
    showPopup(index, item) {
      console.log('show', index, item);
      this.current_index = index
      this.sku_id = item.sku_id
      this.goods.picture = item.cover
      this.sku.price = item.price
      skuDetailSaleAttr(this.sku_id).then(res => {
        this.sku.none_sku = res.is_many == 0
        this.initialSku = {
          selectedNum: item.num,
          selectedProp: {}
        }
        this.sku.tree = res.sale_attrs.map(item => {
          this.initialSku.selectedProp['k_' + item.attr_id] = []
          return {
            k: item.attr_name,
            k_s: 'k_' + item.attr_id,
            v: item.values.map(value => {
              this.initialSku.selectedProp['k_' + item.attr_id].push(value.id)
              return {
                id: value.id,
                name: value.name
              }
            })
          }
        })
        this.sku.list = res.skus.map(item => {
          let sku = {
            id: item.sku_id,
            price: item.price * 100,
            stock_num: item.stock
          }
          item.attrs.forEach(attr => {
            sku['k_' + attr.attr_id] = attr.value_id
            if (item.sku_id == this.list[index].sku_id) {
              this.initialSku['k_' + attr.attr_id] = attr.value_id
            }
          })
          return sku
        })
        this.skuShow = true
        console.log('sku', this.sku)
        console.log('goods', this.goods)
        console.log('initialSku', this.initialSku)
      })
    },
    // 获取数据
    initData() {
      // 获取热门推荐
      skuList(0, 1).then(res => {
        this.hotList = res
      })
    }
  }
}
</script>

<style>
.swipe-button {
  height: 100%;
}
</style>

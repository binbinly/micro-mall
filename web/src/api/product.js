import api from './index'

//商品列表
export function skuList(cat_id, p) {
  return api.get(api.Product.List + '/cat/' + cat_id + '/p/' + p, false)
}

export function skuSearch(data) {
  return api.post(api.Product.Search, data, null, false, false, true)
}

//商品详情
export function skuDetail(id) {
  return api.get(api.Product.Detail + '/' + id)
}

//商品销售属性
export function skuDetailSaleAttr(id) {
  return api.get(api.Product.Attr + '/' + id)
}

//商品所有分类
export function categoryTree() {
  return api.get(api.Product.Category)
}

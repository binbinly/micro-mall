import api from './index'

//首页数据
export function homeData() {
  return api.get(api.App.Home)
}

//首页分类数据
export function homeCatData(cat_id = 0) {
  return api.get(api.App.HomeCat+'/'+cat_id)
}

//通知列表
export function noticeList(page) {
  return api.get(api.App.Notice + '/' + page)
}

//搜索页面配置信息
export function searchData() {
  return api.get(api.App.Search, null, false, true, true)
}

//支付配置列表
export function getPayList() {
  return api.get(api.App.PayList)
}

import $C from '@/config/index.js'
import $Time from '@/utils/time.js'

/**
 *格式化时间
 *yyyy-MM-dd hh:mm:ss
 */
export function formatDate(time, fmt) {
  if (time === undefined || '') {
    return
  }
  const date = new Date(time)
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
  }
  const o = {
    'M+': date.getMonth() + 1,
    'd+': date.getDate(),
    'h+': date.getHours(),
    'm+': date.getMinutes(),
    's+': date.getSeconds()
  }
  for (const k in o) {
    if (new RegExp(`(${k})`).test(fmt)) {
      const str = o[k] + ''
      fmt = fmt.replace(RegExp.$1, RegExp.$1.length === 1 ? str : padLeftZero(str))
    }
  }
  return fmt
}

function padLeftZero(str) {
  return ('00' + str).substr(str.length)
}
/*
 * 隐藏用户手机号中间四位
 */
export function hidePhone(phone) {
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}

export function formatTime(value) {
  return $Time.gettime(value)
}

/**
 *格式化图片地址
 */
export function formatImage(path) {
  return $C.dfsUrl + path
}

export function formatAvatar(url) {
  if (url == '') {
    return require('@/assets/images/avatar.jpg')
  } else {
    return url
  }
}

export function formatOrderStatus(status) {
  const all = { 1: '待支付', 2: '待发货', 3: '已发货', 4: '已收货', 5: '已完成', 6: '待退款', 7: '已退款' }
  return all[status] || '未知'
}

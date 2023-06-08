package repository

import "fmt"

// _cartKeyPrefix 购物车键前缀
var _cartKeyPrefix = "cart:%v"

// CartModel 购物车模型，其不存储mysql，使用redis存储
type CartModel struct {
	SkuID   int64  `json:"sku_id"`   // 商品id
	Title   string `json:"title"`    //商品标题
	Price   int    `json:"price"`    //商品价格
	Cover   string `json:"cover"`    //图片
	SkuAttr string `json:"sku_attr"` //sku属性
	Num     int    `json:"num"`      //数量
	UTime   int64  `json:"u_time"`   //更新时间
}

// Sku 商品
type Sku struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Price     int64  `json:"price"`
	AttrValue string `json:"attr_value"`
}

// BuildCartKey 用户缓存键
func BuildCartKey(userID int64) string {
	return fmt.Sprintf(_cartKeyPrefix, userID)
}

// CartSort 按照 sort 从大到小排序
type CartSort []*CartModel

func (a CartSort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a CartSort) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a CartSort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].UTime < a[i].UTime
}

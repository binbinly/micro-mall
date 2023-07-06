package resource

import (
	"github.com/binbinly/pkg/util"

	"cart/config"
	"cart/repository"
	pb "pkg/proto/common"
)

// CartResource 转换购物车接口输出
func CartResource(cart *repository.CartModel, external bool) *pb.Cart {
	if cart == nil {
		return &pb.Cart{}
	}

	item := &pb.Cart{
		SkuId:   int32(cart.SkuID),
		Title:   cart.Title,
		Price:   float64(cart.Price),
		Cover:   cart.Cover,
		SkuAttr: cart.SkuAttr,
		Num:     int32(cart.Num),
	}
	if external { //对外格式化数据
		item.Price = util.ParseAmount(cart.Price)
		item.Cover = util.FormatResUrl(config.Cfg.DFS.Endpoint, cart.Cover)
	}
	return item
}

// CartListResource 转换购物车列表结构
func CartListResource(list []*repository.CartModel, external bool) []*pb.Cart {
	if len(list) == 0 {
		return make([]*pb.Cart, 0)
	}

	res := make([]*pb.Cart, 0, len(list))
	for _, cart := range list {
		res = append(res, CartResource(cart, external))
	}
	return res
}

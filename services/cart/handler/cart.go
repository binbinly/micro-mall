package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/client"

	"cart/logic"
	"cart/repository"
	"cart/resource"
	"pkg/constvar"
	"pkg/handler"
	pb "pkg/proto/cart"
	cpb "pkg/proto/common"
	"pkg/proto/product"
)

// Auth 购物车服身份验证
func Auth(method string) bool {
	return true
}

// Cart 购物车处理器
type Cart struct {
	Logic   logic.Logic
	Product product.ProductService
}

func New(l logic.Logic, c client.Client) *Cart {
	return &Cart{
		Logic:   l,
		Product: product.NewProductService(constvar.ServiceProduct, c),
	}
}

// AddCart 添加购物车
func (c *Cart) AddCart(ctx context.Context, req *pb.AddReq, empty *empty.Empty) error {
	sku, err := c.skuInfo(ctx, req.SkuId)
	if err != nil {
		return err
	}

	if err = c.Logic.AddCart(ctx, handler.GetUserID(ctx), int(req.Num), sku); err != nil {
		return err
	}
	return nil
}

// EditCart 更新购物车
func (c *Cart) EditCart(ctx context.Context, req *pb.EditReq, empty *empty.Empty) error {
	sku, err := c.skuInfo(ctx, req.NewSkuId)
	if err != nil {
		return err
	}

	if err = c.Logic.EditCart(ctx, handler.GetUserID(ctx), req.OldSkuId, sku, int(req.Num)); err != nil {
		return err
	}
	return nil
}

// EditCartNum 更新购物车数量
func (c *Cart) EditCartNum(ctx context.Context, req *pb.AddReq, empty *empty.Empty) error {
	sku, err := c.skuInfo(ctx, req.SkuId)
	if err != nil {
		return err
	}

	if err = c.Logic.EditCartNum(ctx, handler.GetUserID(ctx), sku, int(req.Num)); err != nil {
		return err
	}
	return nil
}

// DelCart 删除购物车
func (c *Cart) DelCart(ctx context.Context, req *cpb.SkuIDReq, empty *empty.Empty) error {
	if err := c.Logic.DelCart(ctx, handler.GetUserID(ctx), []int64{req.Id}); err != nil {
		return err
	}
	return nil
}

// ClearCart 清空购物车
func (c *Cart) ClearCart(ctx context.Context, empty *empty.Empty, empty2 *empty.Empty) error {
	if err := c.Logic.ClearCart(ctx, handler.GetUserID(ctx)); err != nil {
		return err
	}
	return nil
}

// MyCart 我的购物车
func (c *Cart) MyCart(ctx context.Context, empty *empty.Empty, reply *pb.CartsReply) error {
	carts, err := c.Logic.CartList(ctx, handler.GetUserID(ctx))
	if err != nil {
		return err
	}

	reply.Data = resource.CartListResource(carts, true)
	return nil
}

// BatchGetCarts 批量获取购物车信息
func (c *Cart) BatchGetCarts(ctx context.Context, req *pb.SkusReq, reply *pb.CartsReply) error {
	carts, err := c.Logic.BatchGetCarts(ctx, handler.GetUserID(ctx), req.SkuIds)
	if err != nil {
		return err
	}

	reply.Data = resource.CartListResource(carts, false)
	return nil
}

// BatchDelCarts 批量删除购物车
func (c *Cart) BatchDelCarts(ctx context.Context, req *pb.SkusReq, empty *empty.Empty) error {
	if err := c.Logic.DelCart(ctx, handler.GetUserID(ctx), req.SkuIds); err != nil {
		return err
	}
	return nil
}

// skuInfo 前往产品服务获取sku信息
func (c *Cart) skuInfo(ctx context.Context, skuID int64) (*repository.Sku, error) {
	sku, err := c.Product.GetSkuByID(ctx, &cpb.SkuIDReq{Id: skuID})
	if err != nil {
		return nil, err
	}

	return &repository.Sku{
		ID:        sku.Sku.Id,
		Title:     sku.Sku.Title,
		Cover:     sku.Sku.Cover,
		Price:     sku.Sku.Price,
		AttrValue: sku.Sku.AttrValue,
	}, nil
}

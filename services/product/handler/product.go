package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/client"
	"pkg/constvar"
	"pkg/errno"
	"pkg/handler"
	cpb "pkg/proto/common"
	pb "pkg/proto/product"
	"pkg/proto/seckill"
	"pkg/proto/warehouse"
	wpb "pkg/proto/warehouse"
	"product/logic"
	"product/model"
	"product/resource"
	"reflect"
)

// Auth 产品服身份验证
func Auth(method string) bool {
	return false
}

// Product 产品服务处理器
type Product struct {
	logic            logic.Logic
	warehouseService warehouse.WarehouseService
	seckillService   seckill.SeckillService
}

func New(l logic.Logic, c client.Client) *Product {
	return &Product{
		logic:            l,
		warehouseService: warehouse.NewWarehouseService(constvar.ServiceWarehouse, c),
		seckillService:   seckill.NewSeckillService(constvar.ServiceSeckill, c),
	}
}

// CategoryTree 产品分类树
func (p *Product) CategoryTree(ctx context.Context, empty *empty.Empty, reply *pb.CategoryReply) error {
	tree, err := p.logic.CategoryTree(ctx)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	reply.Data = resource.CategoryResource(tree)
	return nil
}

// SkuSearch 搜索
func (p *Product) SkuSearch(ctx context.Context, req *pb.SearchReq, reply *pb.SearchReply) error {
	attrs := make(map[int64][]string)
	if len(req.Attrs) > 0 {
		for _, attr := range req.Attrs {
			attrs[attr.Id] = attr.Values
		}
	}
	data, err := p.logic.Search(ctx, req.Keyword, req.CatId, req.Field, req.Order,
		req.PriceS, req.PriceE, req.HasStock, req.BrandId, attrs, req.Page)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(data).Elem())
	return nil
}

// SkuList sku商品列表
func (p *Product) SkuList(ctx context.Context, req *pb.SkuListReq, reply *pb.SkuListReply) error {
	list, err := p.logic.SkuList(ctx, req.CatId, req.Page)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	reply.Data = resource.SkuListResource(list)
	return nil
}

// SkuDetail sku商品详情
func (p *Product) SkuDetail(ctx context.Context, req *cpb.SkuIDReq, reply *pb.SkuReply) error {
	sku, err := p.logic.SkuDetail(ctx, req.Id)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	sku.Stocks, err = p.skusStock(ctx, sku.Skus)
	if err != nil {
		return err
	}

	reply.Data = resource.SkuResource(sku)
	return nil
}

// GetSkuSaleAttrs 获取sku的销售属性
func (p *Product) GetSkuSaleAttrs(ctx context.Context, req *cpb.SkuIDReq, reply *pb.SkuSaleAttrReply) error {
	sku, err := p.logic.GetSkuSaleAttrs(ctx, req.Id)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	sku.Stocks, err = p.skusStock(ctx, sku.Skus)
	if err != nil {
		return err
	}

	reply.Data = resource.SkuSaleAttrsResource(sku)
	return nil
}

// GetSkuByID 获取sku信息
func (p *Product) GetSkuByID(ctx context.Context, req *cpb.SkuIDReq, internal *pb.SkuInfoInternal) error {
	sku, err := p.logic.GetSkuInfo(ctx, req.Id)
	if err != nil {
		return errno.ProductReplyErr(err)
	}

	internal.Sku = resource.SkuBaseResource(sku)
	return nil
}

// SpuComment 商品评价
func (p *Product) SpuComment(ctx context.Context, req *pb.CommentReq, empty *empty.Empty) error {
	if err := p.logic.SpuComment(ctx, req.SkuIds, handler.GetUserID(ctx), req.OrderId, int8(req.Star), req.Content, req.Resources); err != nil {
		return errno.ProductReplyErr(err)
	}

	return nil
}

func (p *Product) skusStock(ctx context.Context, skus []*model.SkuModel) (map[int64]int32, error) {
	skuIds := make([]int64, 0, len(skus))
	for _, s := range skus {
		skuIds = append(skuIds, s.ID)
	}
	stocks, err := p.warehouseService.BatchSkusStock(ctx, &wpb.SkusReq{
		Ids: skuIds,
	})
	if err != nil {
		return nil, err
	}

	return stocks.SkuNums, nil
}

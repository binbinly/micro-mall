package resource

import (
	"github.com/binbinly/pkg/util"
	cpb "pkg/proto/common"
	pb "pkg/proto/product"

	"product/config"
	"product/es"
	"product/model"
)

// SkuListResource 转换sku列表
func SkuListResource(list []*es.ProductEs) (res []*cpb.SkuEs) {
	if len(list) == 0 {
		return []*cpb.SkuEs{}
	}
	for _, product := range list {
		res = append(res, &cpb.SkuEs{
			Id:        int32(product.SkuID),
			Title:     product.SkuTitle,
			Price:     util.ParseAmount(product.SkuPrice),
			Cover:     util.FormatResUrl(config.Cfg.DFS.Endpoint, product.SkuImg),
			SaleCount: product.SaleCount,
			HasStock:  product.HasStock,
		})
	}
	return
}

// SearchResResource 转换搜索结果集
func SearchResResource(res *es.SearchRes, cats map[int64]string, brands map[int64]*model.Brand) *pb.SearchReply {
	reply := new(pb.SearchReply)
	for _, product := range res.Products {
		reply.Data = append(reply.Data, &cpb.SkuEs{
			Id:        int32(product.SkuID),
			Title:     product.SkuTitle,
			Price:     util.ParseAmount(product.SkuPrice),
			Cover:     util.FormatResUrl(config.Cfg.DFS.Endpoint, product.SkuImg),
			SaleCount: product.SaleCount,
			HasStock:  product.HasStock,
		})
	}
	for id, name := range cats {
		reply.Cats = append(reply.Cats, &cpb.CatEs{
			Id:   int32(id),
			Name: name,
		})
	}
	for id, brand := range brands {
		reply.Brands = append(reply.Brands, &cpb.BrandEs{
			Id:   int32(id),
			Name: brand.Name,
			Logo: util.FormatResUrl(config.Cfg.DFS.Endpoint, brand.Logo),
		})
	}
	reply.Attrs = parseAttrs(res)

	return reply
}

// parseAttrs 解析规格属性
func parseAttrs(res *es.SearchRes) (attrs []*cpb.AttrEs) {
	for _, bucket := range res.Attrs.AttrIDAgg.Buckets {
		attr := &cpb.AttrEs{
			Id: int32(bucket.Key),
		}
		for _, subBucket := range bucket.AttrNameAgg.Buckets {
			attr.Name = subBucket.Key
		}
		for _, subBucket := range bucket.AttrValueAgg.Buckets {
			attr.Values = append(attr.Values, subBucket.Key)
		}
		attrs = append(attrs, attr)
	}
	return
}

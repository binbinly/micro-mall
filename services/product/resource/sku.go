package resource

import (
	"encoding/json"

	"github.com/binbinly/pkg/util"

	cpb "pkg/proto/common"
	"product/config"
	"product/model"
)

// SkuBaseResource 转换sku基础信息
func SkuBaseResource(sku *model.SkuModel) *cpb.SkuBase {
	return &cpb.SkuBase{
		Id:        sku.ID,
		SpuId:     sku.SpuID,
		CatId:     sku.CatID,
		BrandId:   sku.BrandID,
		Title:     sku.Title,
		Desc:      sku.Desc,
		Cover:     sku.Cover,
		Subtitle:  sku.Subtitle,
		Price:     int64(sku.Price),
		SaleCount: int64(sku.SaleCount),
		AttrValue: sku.AttrValue,
	}
}

// SkuResource 转换sku数据输出
func SkuResource(input *model.Sku) *cpb.Sku {
	sku := new(cpb.Sku)
	sku.Id = int32(input.Sku.ID)
	sku.CatId = int32(input.Sku.CatID)
	sku.SpuId = int32(input.Sku.SpuID)
	sku.BrandId = int32(input.Sku.BrandID)
	sku.Title = input.Sku.Title
	sku.Subtitle = input.Sku.Subtitle
	sku.Desc = input.Sku.Desc
	sku.Price = util.ParseAmount(input.Sku.Price)
	sku.Cover = util.FormatResUrl(config.Cfg.DFS.Endpoint, input.Sku.Cover)
	sku.SaleCount = int32(input.Sku.SaleCount)
	sku.IsMany = input.Spu.IsMany == 1
	sku.Stock = input.Stocks[input.Sku.ID]
	for _, image := range input.SkuImages {
		sku.Banners = append(sku.Banners, util.FormatResUrl(config.Cfg.DFS.Endpoint, image.Img))
	}
	if input.SpuDesc == nil || input.SpuDesc.Content == "" {
		sku.Mains = []string{}
	} else {
		var imgs []string
		_ = json.Unmarshal([]byte(input.SpuDesc.Content), &imgs)
		sku.Mains = buildResUrls(imgs)
	}
	sku.Attrs = transferAttrs(input.SpuAttrs, input.AttrGroups)
	sku.Skus, sku.SaleAttrs = transferSkuAttrs(input.SkuAttrs, input.Skus, input.Stocks)
	return sku
}

// SkuSaleAttrsResource 转换sku销售属性
func SkuSaleAttrsResource(input *model.Sku) *cpb.SkuSaleAttr {
	sku := new(cpb.SkuSaleAttr)
	sku.Id = int32(input.Sku.ID)
	sku.IsMany = input.Spu.IsMany == 1
	sku.Skus, sku.SaleAttrs = transferSkuAttrs(input.SkuAttrs, input.Skus, input.Stocks)
	return sku
}

// transferAttrs 转换规格属性分组及其下属性
func transferAttrs(spuAttrs []*model.Attrs, attrGroups []*model.AttrGroupModel) (as []*cpb.Attrs) {
	attrMap := make(map[int64]*cpb.Attrs)
	for _, attr := range spuAttrs {
		// TODO 响应数据暂时不考虑属性分组
		attr.GroupID = 0
		if attrs, ok := attrMap[attr.GroupID]; ok { //分组已存在，追加其下的属性
			attrs.Items = append(attrs.Items, &cpb.Attr{
				Id:    int32(attr.AttrID),
				Name:  attr.AttrName,
				Value: attr.AttrValue,
			})
		} else { //添加新分组
			attrMap[attr.GroupID] = &cpb.Attrs{
				GroupId:   int32(attr.GroupID),
				GroupName: getGroupName(attrGroups, attr.GroupID),
				Items: []*cpb.Attr{
					{
						Id:    int32(attr.AttrID),
						Name:  attr.AttrName,
						Value: attr.AttrValue,
					},
				},
			}
		}
	}
	for _, attrs := range attrMap {
		as = append(as, attrs)
	}
	return
}

// transferSkuAttrs 转换销售属性
func transferSkuAttrs(skuAttrs []*model.SkuAttrModel, skus []*model.SkuModel, stocks map[int64]int32) (resSkus []*cpb.Skus, resAttrs []*cpb.SaleAttrs) {
	attrsMap := make(map[int64][]*cpb.SkuAttr)
	skuAttrsMap := make(map[int64]*cpb.SaleAttrs)
	for _, skuAttr := range skuAttrs {
		if attr, ok := attrsMap[skuAttr.SkuID]; ok {
			attr = append(attr, &cpb.SkuAttr{
				AttrId:    int32(skuAttr.AttrID),
				ValueId:   int32(skuAttr.ID),
				AttrName:  skuAttr.AttrName,
				ValueName: skuAttr.AttrValue,
			})
		} else {
			attrsMap[skuAttr.SkuID] = []*cpb.SkuAttr{
				{
					AttrId:    int32(skuAttr.AttrID),
					ValueId:   int32(skuAttr.ID),
					AttrName:  skuAttr.AttrName,
					ValueName: skuAttr.AttrValue,
				},
			}
		}
		if attrs, ok := skuAttrsMap[skuAttr.AttrID]; ok {
			attrs.Values = append(attrs.Values, &cpb.SkuValue{
				Id:   int32(skuAttr.ID),
				Name: skuAttr.AttrValue,
			})
		} else {
			skuAttrsMap[skuAttr.AttrID] = &cpb.SaleAttrs{
				AttrId:   int32(skuAttr.AttrID),
				AttrName: skuAttr.AttrName,
				Values: []*cpb.SkuValue{
					{
						Id:   int32(skuAttr.ID),
						Name: skuAttr.AttrValue,
					},
				},
			}
		}
	}
	for _, attrs := range skuAttrsMap {
		resAttrs = append(resAttrs, attrs)
	}
	for _, sku := range skus {
		attrs, ok := attrsMap[sku.ID]
		if !ok {
			attrs = make([]*cpb.SkuAttr, 0)
		}
		resSkus = append(resSkus, &cpb.Skus{
			SkuId: int32(sku.ID),
			Price: util.ParseAmount(sku.Price),
			Stock: stocks[sku.ID],
			Attrs: attrs,
		})
	}
	return
}

// getGroupName 获取分组名
func getGroupName(groups []*model.AttrGroupModel, id int64) string {
	for _, group := range groups {
		if group.ID == id {
			return group.Name
		}
	}
	return ""
}

// buildResUrls 构建资源图片路径列表
func buildResUrls(urls []string) (res []string) {
	for _, url := range urls {
		res = append(res, util.FormatResUrl(config.Cfg.DFS.Endpoint, url))
	}
	return
}

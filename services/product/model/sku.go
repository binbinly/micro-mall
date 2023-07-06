package model

import (
	"pkg/dbs"
)

// SkuModel sku信息
type SkuModel struct {
	dbs.PriID
	dbs.Spu
	dbs.Cat
	BrandID     int64  `json:"brand_id" gorm:"column:brand_id;not null;type:int;comment:品牌id"`
	Name        string `json:"name" gorm:"column:name;not null;type:varchar(60);comment:sku名称"`
	Desc        string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	Cover       string `json:"cover" gorm:"column:cover;not null;type:varchar(128);comment:默认图片"`
	Title       string `json:"title" gorm:"column:title;not null;type:varchar(255);comment:标题"`
	Subtitle    string `json:"subtitle" gorm:"column:subtitle;not null;type:varchar(255);comment:副标题"`
	Price       int    `json:"price" gorm:"column:price;not null;type:int;comment:价格"`
	SaleCount   int    `json:"sale_count" gorm:"column:sale_count;not null;type:int;default:0;comment:销量"`
	ReviewCount int    `json:"review_count" gorm:"column:review_count;not null;type:int;default:0;comment:评论数"`
	AttrValue   string `json:"attr_value" gorm:"column:attr_value;not null;type:varchar(120);default:'';comment:销售属性值"`
}

// TableName 表名
func (u *SkuModel) TableName() string {
	return "pms_sku"
}

// SkuEs es中sku结构
type SkuEs struct {
	ID        int64  `json:"id"`
	SpuID     int64  `json:"spuId"`
	CatID     int64  `json:"catId"`
	BrandID   int64  `json:"brandId"`
	SkuTitle  string `json:"skuTitle"`
	SkuPrice  int32  `json:"skuPrice"`
	SkuImg    string `json:"skuImg"`
	SaleCount int32  `json:"saleCount"`
	HasStock  bool   `json:"hasStock"`
	HotScore  int32  `json:"hotScore"`
	CatName   string `json:"catName"`
	BrandName string `json:"brandName"`
	BrandImg  string `json:"brandImg"`
}

// Sku 详细结构
type Sku struct {
	Sku        *SkuModel
	Spu        *SpuModel
	SpuDesc    *SpuDescModel
	Skus       []*SkuModel
	SkuImages  []*SkuImageModel
	AttrGroups []*AttrGroupModel
	SkuAttrs   []*SkuAttrModel
	SpuAttrs   []*Attrs
	Stocks     map[int64]int32
}

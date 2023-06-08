package es

const mapping = `
{
    "mappings":{
		"properties":{
			"skuId":{
				"type":"long"
			},
			"spuId":{
				"type":"long"
			},
			"skuTitle":{
				"type":"text",
				"analyzer":"ik_smart"
			},
			"skuPrice":{
				"type":"integer"
			},
			"skuImg":{
				"type":"keyword",
				"index":false,
				"doc_values":false
			},
			"saleCount":{
				"type":"integer"
			},
			"hasStock":{
				"type":"boolean"
			},
			"hotScore":{
				"type":"integer"
			},
			"brandId":{
				"type":"long"
			},
			"catId":{
				"type":"long"
			},
			"catName":{
				"type":"keyword",
				"index":false,
				"doc_values":false
			},
			"brandName":{
				"type":"keyword",
				"index":false,
				"doc_values":false
			},
			"brandImg":{
				"type":"keyword",
				"index":false,
				"doc_values":false
			},
			"attrs":{
				"type":"nested",
				"properties":{
					"attrId":{
						"type":"long"
					},
					"attrName":{
						"type":"keyword",
						"index":false,
						"doc_values":true
					},
					"attrValue":{
						"type":"keyword"
					}
				}
			}
		}  
    }
}`

// ProductEs es中product数据结构
type ProductEs struct {
	SkuID     int64    `json:"skuId"`
	SpuID     int64    `json:"spuId"`
	CatID     int64    `json:"catId"`
	BrandID   int64    `json:"brandId"`
	SkuTitle  string   `json:"skuTitle"`
	HighTitle string   `json:"high_title"`
	SkuPrice  int      `json:"skuPrice"`
	SkuImg    string   `json:"skuImg"`
	SaleCount int32    `json:"saleCount"`
	HasStock  bool     `json:"hasStock"`
	HotScore  int32    `json:"hotScore"`
	CatName   string   `json:"catName"`
	BrandName string   `json:"brandName"`
	BrandImg  string   `json:"brandImg"`
	Attrs     []AttrEs `json:"attrs"`
}

// SearchRes es搜索结果
type SearchRes struct {
	Products []ProductEs `json:"products"`
	Cats     Agg         `json:"cats"`
	Brands   Agg         `json:"brands"`
	Attrs    AttrAgg     `json:"attrs"`
}

// AttrEs 属性结构
type AttrEs struct {
	AttrId    int64  `json:"attrId"`
	AttrName  string `json:"attrName"`
	AttrValue string `json:"attrValue"`
}

// Agg 聚合结构
type Agg struct {
	Buckets []Bucket `json:"buckets"`
}

// Bucket bucket
type Bucket struct {
	Key int `json:"key"`
}

// AttrAgg 属性聚合结构
type AttrAgg struct {
	AttrIDAgg AttrIDAgg `json:"attr_id_agg"`
}

// AttrIDAgg 属性id聚合结构
type AttrIDAgg struct {
	Buckets []AttrBucket `json:"buckets"`
}

// AttrBucket attr bucket
type AttrBucket struct {
	Key          int    `json:"key"`
	AttrNameAgg  SubAgg `json:"attr_name_agg"`
	AttrValueAgg SubAgg `json:"attr_value_agg"`
}

// SubAgg 子聚合结构
type SubAgg struct {
	Buckets []SubBucket `json:"buckets"`
}

// SubBucket bucket
type SubBucket struct {
	Key string `json:"key"`
}

package es

import "github.com/olivere/elastic/v7"

// Query es query子句封装
type Query struct {
	query *elastic.BoolQuery
}

// NewQuery 实例化
func NewQuery() *Query {
	return &Query{query: elastic.NewBoolQuery()}
}

// Keyword 关键字匹配
func (q *Query) Keyword(keyword string) *Query {
	q.query.Must(elastic.NewMatchQuery("skuTitle", keyword))
	return q
}

// FilterCatID 过滤分类
func (q *Query) FilterCatID(catID int64) *Query {
	q.query.Filter(elastic.NewTermQuery("catId", catID))
	return q
}

// FilterBrandIds 过滤品牌
func (q *Query) FilterBrandIds(brandIds []int64) *Query {
	if len(brandIds) == 0 {
		return q
	}
	var ids []interface{}
	for _, id := range brandIds {
		ids = append(ids, id)
	}
	q.query.Filter(elastic.NewTermsQuery("brandId", ids...))
	return q
}

// FilterAttrs 过滤属性
func (q *Query) FilterAttrs(attrID int64, attrValues []string) *Query {
	if len(attrValues) == 0 {
		return q
	}
	var values []interface{}
	for _, value := range attrValues {
		values = append(values, value)
	}
	attrQ := elastic.NewBoolQuery()
	attrQ.Must(elastic.NewTermQuery("attrs.attrId", attrID),
		elastic.NewTermsQuery("attrs.attrValue", values...))
	q.query.Filter(elastic.NewNestedQuery("attrs", attrQ))
	return q
}

// FilterHasStock 过滤是否有库存
func (q *Query) FilterHasStock(hasStock bool) *Query {
	q.query.Filter(elastic.NewTermQuery("hasStock", hasStock))
	return q
}

// FilterSkuPrice 过滤价格区间
func (q *Query) FilterSkuPrice(startPrice, endPrice int32) *Query {
	if startPrice == 0 && endPrice == 0 {
		return q
	}
	if endPrice > 0 && startPrice > endPrice {
		return q
	}
	query := elastic.NewRangeQuery("skuPrice").Gte(startPrice)
	if endPrice > 0 {
		query.Lte(endPrice)
	}
	q.query.Filter(query)
	return q
}

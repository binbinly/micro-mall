package es

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

const (
	// 产品索引
	_productIndex = "product"
	// 汇总数量
	_aggSize = 10
)

const (
	//综合排序
	sortDefaultField = iota
	//销量排序
	sortSaleCountField
	//价格排序
	sortPriceField
)

var (
	sortField = map[int32]string{
		sortDefaultField:   "hotScore",
		sortPriceField:     "skuPrice",
		sortSaleCountField: "saleCount",
	}
)

// Product 产品es索引
type Product struct {
	client *elastic.Client
}

// New 实例化
func New(cli *elastic.Client) *Product {
	return &Product{client: cli}
}

// Init 初始化索引
func (p *Product) Init(ctx context.Context) error {
	exists, err := p.client.IndexExists(_productIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		// Create a new index.
		_, err = p.client.CreateIndex(_productIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create 创建数据
func (p *Product) Create(ctx context.Context, product ProductEs) error {
	_, err := p.client.Index().Index(_productIndex).
		Id(strconv.FormatInt(product.SkuID, 10)).
		BodyJson(product).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.product] create")
	}
	return nil
}

// Query 查询
func (p *Product) Query(ctx context.Context, catID int64, from, size int) ([]*ProductEs, error) {
	var query elastic.Query
	if catID > 0 {
		query = elastic.NewTermQuery("catId", catID)
	} else {
		query = elastic.NewMatchAllQuery()
	}
	res, err := p.client.Search(_productIndex).Query(query).
		From(from).Size(size).Do(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.query]")
	}
	var pro ProductEs
	var products []*ProductEs
	for _, item := range res.Each(reflect.TypeOf(pro)) {
		t := item.(ProductEs)
		products = append(products, &t)
	}
	return products, nil
}

// Search 搜索
func (p *Product) Search(ctx context.Context, q *Query, from, size int, field, order int32) (*SearchRes, error) {
	fieldName := "hotScore"
	if f, ok := sortField[field]; ok {
		fieldName = f
	}
	res, err := p.client.Search(_productIndex).Query(q.query).Highlight(p.highlight()).
		Aggregation("cat_agg", elastic.NewTermsAggregation().Field("catId").Size(_aggSize)).
		Aggregation("brand_agg", elastic.NewTermsAggregation().Field("brandId").Size(_aggSize)).
		Aggregation("attr_agg", elastic.NewNestedAggregation().Path("attrs").
			SubAggregation("attr_id_agg", elastic.NewTermsAggregation().Field("attrs.attrId").Size(5).
				SubAggregation("attr_name_agg", elastic.NewTermsAggregation().Field("attrs.attrName").Size(5)).
				SubAggregation("attr_value_agg", elastic.NewTermsAggregation().Field("attrs.attrValue").Size(5)))).
		Sort(fieldName, order == 0).
		From(from).Size(size).Do(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.search]")
	}
	var pro ProductEs
	var products []ProductEs
	for i, item := range res.Each(reflect.TypeOf(pro)) {
		t := item.(ProductEs)
		if len(res.Hits.Hits[i].Highlight["skuTitle"]) > 0 {
			t.HighTitle = res.Hits.Hits[i].Highlight["skuTitle"][0]
		}
		products = append(products, t)
	}
	var catAgg Agg
	err = json.Unmarshal(res.Aggregations["cat_agg"], &catAgg)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.search] cat_agg json unmarshal")
	}
	var brandAgg Agg
	err = json.Unmarshal(res.Aggregations["brand_agg"], &brandAgg)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.search] brand_agg json unmarshal")
	}
	var attrAgg AttrAgg
	err = json.Unmarshal(res.Aggregations["attr_agg"], &attrAgg)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.search] attr_agg json unmarshal")
	}

	return &SearchRes{
		Products: products,
		Cats:     catAgg,
		Brands:   brandAgg,
		Attrs:    attrAgg,
	}, nil
}

// highlight 高亮设置
func (p *Product) highlight() *elastic.Highlight {
	high := elastic.NewHighlight()
	high.Field("skuTitle")
	high.PreTags("<span style='color:red;font-weight:bold;'>")
	high.PostTags("</span>")
	return high
}

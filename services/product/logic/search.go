package logic

import (
	"context"

	"github.com/pkg/errors"

	"pkg/constvar"
	"product/es"
	"product/model"
)

// SkuList sku商品列表
func (l *logic) SkuList(ctx context.Context, catID int64, page int32) ([]*es.ProductEs, error) {
	return l.productEs.Query(ctx, catID, constvar.GetPageOffset(page), constvar.DefaultLimit)
}

// Search 搜索
func (l *logic) Search(ctx context.Context, keyword string, catID int64, field, order, priceS, priceE int32,
	hasStock bool, brandIds []int64, attrs map[int64][]string, page int32) (*es.SearchRes, error) {
	query := es.NewQuery()
	if keyword != "" {
		query.Keyword(keyword)
	}
	if catID > 0 {
		query.FilterCatID(catID)
	}
	if len(brandIds) > 0 {
		query.FilterBrandIds(brandIds)
	}
	query.FilterSkuPrice(priceS, priceE)
	if len(attrs) > 0 {
		for attrID, values := range attrs {
			query.FilterAttrs(attrID, values)
		}
	}
	res, err := l.productEs.Search(ctx, query, constvar.GetPageOffset(page), constvar.DefaultLimit, field, order)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.search] get result")
	}

	return res, nil
}

// ParseCats 解析商品分类信息
func (l *logic) ParseCats(ctx context.Context, res *es.SearchRes) (map[int64]string, error) {
	//结果集中所有分类id
	var catIds []int64
	for _, bucket := range res.Cats.Buckets {
		catIds = append(catIds, int64(bucket.Key))
	}
	//获取分类名
	names, err := l.repo.GetCategoryNamesByIds(ctx, catIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.search] get cat names")
	}
	return names, nil
}

// ParseBrands 解析品牌信息
func (l *logic) ParseBrands(ctx context.Context, res *es.SearchRes) (map[int64]*model.Brand, error) {
	//结果集中所有品牌id
	var brandIds []int64
	for _, bucket := range res.Brands.Buckets {
		brandIds = append(brandIds, int64(bucket.Key))
	}
	//获取品牌信息
	brands, err := l.repo.GetBrandsByIds(ctx, brandIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.search] get brand info")
	}
	return brands, nil
}

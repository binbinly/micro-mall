package logic

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"product/model"
	"product/repository"

	"product/es"

	"github.com/binbinly/pkg/cache"
)

var _ Logic = (*logic)(nil)

// Logic 营销服务接口
type Logic interface {
	CategoryTree(ctx context.Context) ([]*model.CategoryModel, error)
	SkuList(ctx context.Context, catID int64, page int32) ([]*es.ProductEs, error)
	SkuDetail(ctx context.Context, id int64) (*model.Sku, error)
	GetSkuSaleAttrs(ctx context.Context, id int64) (*model.Sku, error)
	GetSkuInfo(ctx context.Context, id int64) (*model.SkuModel, error)
	SpuComment(ctx context.Context, skuIds []int64, memberID, orderID int64, star int8, content, resources string) error
	Search(ctx context.Context, keyword string, catID int64, field, order, priceS, priceE int32,
		hasStock bool, brandIds []int64, attrs map[int64][]string, page int32) (*es.SearchRes, error)

	Close() error
}

type logic struct {
	repo      repository.IRepo
	productEs *es.Product
}

// New init logic
func New(e *es.Product, db *gorm.DB, rdb *redis.Client) Logic {
	return &logic{
		repo:      repository.New(db, cache.NewRedisCache(rdb)),
		productEs: e,
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}

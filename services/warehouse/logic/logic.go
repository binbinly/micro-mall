package logic

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"warehouse/repository"
)

var _ Logic = (*logic)(nil)

// Logic 仓储服务接口
type Logic interface {
	GetSkuStock(ctx context.Context, skuID int64) (int, error)
	GetSkusStock(ctx context.Context, skuIds []int64) (map[int64]int32, error)
	SKuStockLock(ctx context.Context, orderID int64, orderNo, consignee, phone, address, note string,
		sku map[int64]int32) error
	SkuStockUnlock(ctx context.Context, orderID int64, finish bool) error

	Close() error
}

type logic struct {
	repo repository.IRepo
}

// New init logic
func New(db *gorm.DB, rdb *redis.Client) Logic {
	return &logic{
		repo: repository.New(db, cache.NewRedisCache(rdb)),
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}

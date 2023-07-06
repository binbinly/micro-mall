package logic

import (
	"context"

	"github.com/redis/go-redis/v9"

	"seckill/repository"
)

var _ Logic = (*logic)(nil)

// Logic 营销服务接口
type Logic interface {
	SeckillCheck(ctx context.Context, memberID, skuID, num int64, key string) (*repository.SkuModel, error)
	GetSessionAll(ctx context.Context) ([]*repository.SessionModel, error)
	GetSessionSkus(ctx context.Context, sessionID int64) ([]*repository.SkuModel, error)
	GetSkuInfo(ctx context.Context, skuID int64) (*repository.SkuModel, error)

	Close() error
}

type logic struct {
	repo repository.IRepo
	rdb  *redis.Client
}

// New init logic
func New(rdb *redis.Client) Logic {
	return &logic{
		rdb:  rdb,
		repo: repository.New(rdb),
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}

package logic

import (
	"context"
	
	"github.com/redis/go-redis/v9"

	"cart/repository"
)

var _ Logic = (*logic)(nil)

// Logic 营销服务接口
type Logic interface {
	AddCart(ctx context.Context, userID int64, num int, sku *repository.Sku) error
	EditCart(ctx context.Context, userID, oldSkuID int64, sku *repository.Sku, num int) error
	EditCartNum(ctx context.Context, userID int64, sku *repository.Sku, num int) error
	DelCart(ctx context.Context, userID int64, skuIds []int64) error
	ClearCart(ctx context.Context, userID int64) error
	CartList(ctx context.Context, userID int64) ([]*repository.CartModel, error)
	BatchGetCarts(ctx context.Context, userID int64, skuIds []int64) ([]*repository.CartModel, error)

	Close() error
}

// logic 营销服务
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

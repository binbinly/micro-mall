package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	AddCart(ctx context.Context, userID int64, cart *CartModel) error
	EditCart(ctx context.Context, userID int64, oldID int64, cart *CartModel) error
	GetCartByID(ctx context.Context, userID int64, skuID int64) (*CartModel, error)
	GetCartsByIds(ctx context.Context, userID int64, ids []int64) ([]*CartModel, error)
	DelCart(ctx context.Context, userID int64, ids []int64) error
	EmptyCart(ctx context.Context, userID int64) error
	CartList(ctx context.Context, userID int64) ([]*CartModel, error)

	Close() error
}

// Repo dbs struct
type Repo struct {
	redis *redis.Client
}

// New new a Dao and return
func New(redis *redis.Client) IRepo {
	return &Repo{
		redis: redis,
	}
}

// Close release dbs connection
func (r *Repo) Close() error {
	return r.redis.Close()
}

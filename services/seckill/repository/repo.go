package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var _ IRepo = (*Repo)(nil)

// SessionModel 秒杀场次
type SessionModel struct {
	ID      int64       `json:"id"`       // id
	Name    string      `json:"name"`     // 场次名
	StartAt int64       `json:"start_at"` // 开始时间
	EndAt   int64       `json:"end_at"`   // 结束时间
	Skus    []*SkuModel `json:"skus"`     // 秒杀商品
}

// SkuModel 秒杀商品
type SkuModel struct {
	ID            int64  `json:"id"`             // 商品id
	Price         int    `json:"price"`          // 秒杀价格
	Count         int    `json:"count"`          // 秒杀数量
	Limit         int64  `json:"limit"`          // 个人限购
	OriginalPrice int    `json:"original_price"` // 原价
	StartAt       int64  `json:"start_at"`       // 开始时间
	EndAt         int64  `json:"end_at"`         // 结束时间
	Title         string `json:"title"`          // 标题
	Cover         string `json:"cover"`          // 封面
	Key           string `json:"key"`            // 加密key
}

// IRepo 数据仓库接口
type IRepo interface {
	GetSessionAll(ctx context.Context) ([]*SessionModel, error)
	GetSkuByID(ctx context.Context, skuID int64) (*SkuModel, error)
	GetSkusBySessionID(ctx context.Context, sessionID int64) ([]*SkuModel, error)

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

package logic

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"market/model"
	"market/repository"
)

var _ Logic = (*logic)(nil)

type Logic interface {
	GetHomeCat(ctx context.Context) ([]*model.ConfigHomeCat, error)
	GetHomeCatData(ctx context.Context, cid int) ([]*model.AppSettingModel, error)
	GetNoticeList(ctx context.Context, page int32) ([]*model.AppNoticeModel, error)
	GetSearchData(ctx context.Context) ([]*model.AppSettingModel, []string, error)
	GetPayConfig(ctx context.Context) ([]*model.ConfigPayList, error)
	GetCouponList(ctx context.Context, memberID, skuID int64) ([]*model.CouponModel, []int64, error)
	GetMyCouponList(ctx context.Context, memberID int64) ([]*model.Coupon, error)
	CouponDraw(ctx context.Context, memberID, id int64) error
	CouponUsed(ctx context.Context, memberID, id, orderID int64) error
	GetCouponInfo(ctx context.Context, memberID, id int64) (*model.CouponMemberModel, *model.CouponModel, error)

	Close() error
}

type logic struct {
	repo repository.IRepo
	rdb  *redis.Client
}

// New init logic
func New(db *gorm.DB, rdb *redis.Client) Logic {
	return &logic{
		repo: repository.New(db, cache.NewRedisCache(rdb)),
		rdb:  rdb,
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}

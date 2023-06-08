package repository

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"

	"gorm.io/gorm"

	"order/model"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OrderModel) error
	GetOrderDetail(ctx context.Context, id, memberID int64) (*model.OrderModel, error)
	GetOrderByID(ctx context.Context, id int64) (*model.OrderModel, error)
	GetOrderByNo(ctx context.Context, orderNo string) (*model.OrderModel, error)
	GetOrderList(ctx context.Context, memberID int64, status, offset, limit int) (list []*model.OrderModel, err error)
	OrderSave(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (err error)
	OrderDelete(ctx context.Context, order *model.OrderModel) error
	CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OrderRefundModel) error
	BatchCreateOrderItem(ctx context.Context, tx *gorm.DB, items []*model.OrderItemModel) error
	CreateOrderBill(ctx context.Context, tx *gorm.DB, bill *model.OrderBillModel) error

	GetDB() *gorm.DB
	Close() error
}

// Repo struct
type Repo struct {
	repo.Repo
}

// New a Dao and return
func New(db *gorm.DB, c cache.Cache) IRepo {
	return &Repo{repo.Repo{
		DB:    db,
		Cache: c,
	}}
}

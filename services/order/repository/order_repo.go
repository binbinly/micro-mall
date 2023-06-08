package repository

import (
	"context"
	"pkg/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"order/model"
)

// CreateOrder 创建订单
func (r *Repo) CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OrderModel) error {
	if err := tx.WithContext(ctx).Create(order).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] create")
	}

	return nil
}

// GetOrderByID 获取订单信息
func (r *Repo) GetOrderByID(ctx context.Context, id int64) (*model.OrderModel, error) {
	order := new(model.OrderModel)
	if err := r.DB.WithContext(ctx).First(order, id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by id: %v", id)
	}

	return order, nil
}

// GetOrderDetail 获取订单详情
func (r *Repo) GetOrderDetail(ctx context.Context, id, memberID int64) (*model.OrderModel, error) {
	order := new(model.OrderModel)
	err := r.DB.WithContext(ctx).Where("id=? and member_id=?", id, memberID).Preload("Items").First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by id: %v", id)
	}

	return order, nil
}

// GetOrderByNo 订单号获取订单信息
func (r *Repo) GetOrderByNo(ctx context.Context, orderNo string) (*model.OrderModel, error) {
	order := new(model.OrderModel)
	err := r.DB.WithContext(ctx).Where("order_no=?", orderNo).First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by no: %v", orderNo)
	}

	return order, nil
}

// GetOrderList 用户订单列表
func (r *Repo) GetOrderList(ctx context.Context, memberID int64, status, offset, limit int) (list []*model.OrderModel, err error) {
	err = r.DB.WithContext(ctx).Where("member_id=?", memberID).Preload("Items").Order(mysql.DefaultOrder).
		Scopes(orderScopesStatus(status), mysql.OffsetPage(offset, limit)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.order] find")
	}

	return
}

// OrderSave 订单保存
func (r *Repo) OrderSave(ctx context.Context, tx *gorm.DB, order *model.OrderModel) (err error) {
	if tx == nil {
		err = r.DB.WithContext(ctx).Save(order).Error
	} else {
		err = tx.WithContext(ctx).Save(order).Error
	}
	if err != nil {
		return errors.Wrapf(err, "[repo.order] save")
	}
	return nil
}

// OrderDelete 订单删除
func (r *Repo) OrderDelete(ctx context.Context, order *model.OrderModel) error {
	if err := r.DB.WithContext(ctx).Delete(order).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] delete")
	}

	return nil
}

// orderScopesStatus 状态筛选
func orderScopesStatus(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status == 0 {
			return db
		}
		return db.Where("status=?", status)
	}
}

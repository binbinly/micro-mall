package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"order/model"
)

// CreateOrderRefund 创建订单退款记录
func (r *Repo) CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OrderRefundModel) error {
	if err := tx.WithContext(ctx).Create(&refund).Error; err != nil {
		return errors.Wrapf(err, "[repo.orderRefund] create")
	}

	return nil
}

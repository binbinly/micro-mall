package cmd

import (
	"order/config"
	"order/model"
	"pkg/dbs"
)

// Migrate 数据迁移
func Migrate() {
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			new(model.ConfigModel),
			new(model.OrderModel),
			new(model.OrderItemModel),
			new(model.OrderBillModel),
			new(model.OrderOperateModel),
			new(model.OrderReturnApplyModel),
			new(model.OrderReturnReasonModel),
			new(model.OrderRefundModel),
			new(model.PaymentModel),
		}
	})
}

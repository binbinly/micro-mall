package cmd

import (
	"pkg/dbs"
	"warehouse/config"
	"warehouse/model"
)

// Migrate 数据迁移
func Migrate() {
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			new(model.PurchaseModel),
			new(model.PurchaseDetailModel),
			new(model.WareModel),
			new(model.WareSkuModel),
			new(model.WareTaskDetailModel),
			new(model.WareTaskModel),
		}
	})
}

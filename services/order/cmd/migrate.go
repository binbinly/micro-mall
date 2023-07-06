package cmd

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"gorm.io/gorm"

	"order/config"
	"order/model"
)

// Migrate 数据迁移
func Migrate() {
	// 初始化数据库链接
	db := orm.NewDB(&config.Cfg.MySQL)
	// 数据库迁移
	log.Println("数据库迁移开始")
	if err := migrateUp(db); err != nil {
		log.Fatalf("db err:%v", err)
	}
	log.Println(`数据库基础数据初始化成功`)
}

// migrateUp 数据迁移
func migrateUp(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.ConfigModel),
			new(model.OrderModel),
			new(model.OrderItemModel),
			new(model.OrderBillModel),
			new(model.OrderOperateModel),
			new(model.OrderReturnApplyModel),
			new(model.OrderReturnReasonModel),
			new(model.OrderRefundModel),
			new(model.PaymentModel),
		)
		if err != nil {
			return err
		}
		return nil
	})
}

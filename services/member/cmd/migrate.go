package cmd

import (
	"fmt"
	"log"
	"pkg/mysql"

	"gorm.io/gorm"

	"member/config"
	"member/model"
)

// Migrate 数据迁移
func Migrate() {
	// 初始化数据库链接
	db := mysql.NewDB(&config.Cfg.MySQL)
	// 数据库迁移
	fmt.Println("数据库迁移开始")
	if err := migrateUp(db); err != nil {
		log.Fatalf("db err:%v", err)
	}
	fmt.Println(`数据库基础数据初始化成功`)
}

// migrateUp 数据迁移
func migrateUp(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.GrowthLogModel),
			new(model.IntegrationLogModel),
			new(model.MemberAddressModel),
			new(model.CollectSpuModel),
			new(model.CollectSubjectModel),
			new(model.MemberLevelModel),
			new(model.LoginLogModel),
			new(model.MemberModel),
			new(model.MemberStatModel),
		)
		if err != nil {
			return err
		}
		return nil
	})
}

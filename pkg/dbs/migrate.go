package dbs

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"gorm.io/gorm"
)

// Migrate 数据迁移
func Migrate(cfg *orm.Config, up func() []any) {
	// 初始化数据库链接
	db := orm.NewDB(cfg)
	// 数据库迁移
	log.Println("数据库迁移开始")
	if err := migrateUp(db, up); err != nil {
		log.Fatalf("db err:%v", err)
	}
	log.Println(`数据库基础数据初始化成功`)
}

// migrateUp 数据迁移
func migrateUp(db *gorm.DB, up func() []any) error {
	models := up()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Migrator().AutoMigrate(models...); err != nil {
			return err
		}
		return nil
	})
}

package cmd

import (
	"center/config"
	"center/model"
	"pkg/dbs"
)

// Migrate 数据迁移
func Migrate() {
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			model.UserModel{},
		}
	})
}

package mysql

import (
	"github.com/binbinly/pkg/storage/orm"
	"gorm.io/gorm"
)

// DB 数据库全局变量
var DB *gorm.DB

// NewDB new mysql db
func NewDB(cfg *orm.Config) *gorm.DB {
	DB = orm.NewMySQL(cfg)
	return DB
}

// NewBasicDB new mysql db
func NewBasicDB(host, user, pwd, name string) *gorm.DB {
	DB = orm.NewBasicMySQL(host, user, pwd, name)
	return DB
}

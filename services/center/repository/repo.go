package repository

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"gorm.io/gorm"

	"center/model"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	// UserCreate 创建用户
	UserCreate(ctx context.Context, user *model.UserModel) (id int64, err error)
	// UserUpdate 修改用户信息
	UserUpdate(ctx context.Context, id int64, userMap map[string]interface{}) error
	// UserUpdatePwd 修改用户密码
	UserUpdatePwd(ctx context.Context, user *model.UserModel) error
	// GetUserByID id获取用户信息
	GetUserByID(ctx context.Context, id int64) (*model.UserModel, error)
	// GetUserByUsername username获取用户信息
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	// GetUserByPhone phone获取用户信息
	GetUserByPhone(ctx context.Context, phone int64) (*model.UserModel, error)
	// UserExist 用户是否已存在
	UserExist(ctx context.Context, username string, phone int64) (bool, error)

	GetDB() *gorm.DB
	Close() error
}

// Repo dbs struct
type Repo struct {
	repo.Repo
}

// New new a Dao and return
func New(db *gorm.DB, c cache.Cache) IRepo {
	return &Repo{repo.Repo{
		DB:    db,
		Cache: c,
	}}
}

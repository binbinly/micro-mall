package logic

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"center/model"
	"center/repository"
)

var _ Logic = (*logic)(nil)

// Logic 中心服接口定义
type Logic interface {
	// UserRegister 用户注册
	UserRegister(ctx context.Context, username, password string, phone int64) (id int64, err error)
	// UsernameLogin 用户名登录
	UsernameLogin(ctx context.Context, username, password string) (*model.UserModel, string, error)
	// UserPhoneLogin 手机号登录
	UserPhoneLogin(ctx context.Context, phone int64) (*model.UserModel, string, error)
	// UserEditPwd 修改密码
	UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error
	// UserEdit 修改用户信息
	UserEdit(ctx context.Context, id int64, userMap map[string]any) error
	// UserInfoByID 获取用户详情
	UserInfoByID(ctx context.Context, id int64) (*model.UserModel, error)
	// UserLogout 用户登出
	UserLogout(ctx context.Context, id int64) error

	Close() error
}

type logic struct {
	rdb  *redis.Client
	repo repository.IRepo
}

// New init logic
func New(db *gorm.DB, rdb *redis.Client) Logic {
	return &logic{
		rdb:  rdb,
		repo: repository.New(db, cache.NewRedisCache(rdb)),
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}

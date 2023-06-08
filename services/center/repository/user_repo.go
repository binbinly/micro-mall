package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"center/model"
)

// UserCreate 创建用户
func (r *Repo) UserCreate(ctx context.Context, user *model.UserModel) (id int64, err error) {
	if err = r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.user] Create err")
	}
	r.DelCache(ctx, userCacheKey(user.ID))

	return user.ID, nil
}

// UserUpdate 更新用户信息
func (r *Repo) UserUpdate(ctx context.Context, id int64, um map[string]any) error {
	if err := r.DB.WithContext(ctx).Model(&model.UserModel{}).Where("id=?", id).Updates(um).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update")
	}
	r.DelCache(ctx, userCacheKey(id))

	return nil
}

// UserUpdatePwd 修改用户密码
func (r *Repo) UserUpdatePwd(ctx context.Context, user *model.UserModel) error {
	if err := r.DB.WithContext(ctx).Save(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update pwd")
	}
	return nil
}

// GetUserByID 获取用户
func (r *Repo) GetUserByID(ctx context.Context, id int64) (user *model.UserModel, err error) {
	if err = r.QueryCache(ctx, userCacheKey(id), &user, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).First(data, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] query cache")
	}

	return user, nil
}

// GetUserByUsername 根据账号获取用户
func (r *Repo) GetUserByUsername(ctx context.Context, username string) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by username")
	}
	return user, nil
}

// GetUserByPhone 根据手机号获取用户
func (r *Repo) GetUserByPhone(ctx context.Context, phone int64) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.DB.WithContext(ctx).Where("phone = ?", phone).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by phone")
	}
	return user, nil
}

// UserExist 用户是否已存在
func (r *Repo) UserExist(ctx context.Context, username string, phone int64) (bool, error) {
	var c int64
	if err := r.DB.WithContext(ctx).Model(&model.UserModel{}).
		Where("phone = ? or username=?", phone, username).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.user] username %v or phone %v does not exist", username, phone)
	}
	return c > 0, nil
}

func userCacheKey(id int64) string {
	return fmt.Sprintf("user:%d", id)
}

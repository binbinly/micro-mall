package logic

import (
	"context"
	"time"

	"github.com/binbinly/pkg/auth"
	"github.com/pkg/errors"

	"center/config"
	"center/model"
	"pkg/constvar"
	"pkg/errno"
)

// UserRegister 注册用户
func (l *logic) UserRegister(ctx context.Context, username, password string, phone int64) (id int64, err error) {
	u := &model.UserModel{
		Username: username,
		Password: password,
		Phone:    phone,
		Status:   model.StatusNormal,
	}
	exist, err := l.repo.UserExist(ctx, username, phone)
	if err != nil {
		return 0, errors.Wrapf(err, "[center.user] user exist")
	}
	if exist {
		return 0, errno.ErrUserExisted
	}
	id, err = l.repo.UserCreate(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "[center.user] create user")
	}
	return id, nil
}

// UsernameLogin 用户名密码登录
func (l *logic) UsernameLogin(ctx context.Context, username, password string) (*model.UserModel, string, error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	user, err := l.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] err from db by username: %l", username)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", errno.ErrUserNotFound
	}

	if user.Status != model.StatusNormal {
		return nil, "", errno.ErrUserFrozen
	}

	// Compare the login password with the user password.
	if err = user.Compare(password); err != nil {
		return nil, "", errno.ErrUserNotMatch
	}

	token, err := l.generateToken(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// UserPhoneLogin 邮箱登录
func (l *logic) UserPhoneLogin(ctx context.Context, phone int64) (*model.UserModel, string, error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	user, err := l.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[center.user] err from db by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", errno.ErrUserNotFound
	}

	if user.Status != model.StatusNormal {
		return nil, "", errno.ErrUserFrozen
	}

	token, err := l.generateToken(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// UserEdit update user info
func (l *logic) UserEdit(ctx context.Context, id int64, um map[string]any) error {
	if err := l.repo.UserUpdate(ctx, id, um); err != nil {
		return errors.Wrapf(err, "[center.user] update user by id: %d", id)
	}

	return nil
}

// UserEditPwd 修改用户密码
func (l *logic) UserEditPwd(ctx context.Context, id int64, oldPassword, password string) error {
	user, err := l.userinfo(ctx, id)
	if err != nil {
		return err
	}

	if err = user.Compare(oldPassword); err != nil {
		return errno.ErrUserPwd
	}

	user.Password = password
	if err = l.repo.UserUpdatePwd(ctx, user); err != nil {
		return errors.Wrapf(err, "[center.user] update user pwd by id:%v", id)
	}
	return nil
}

// UserInfoByID 获取用户信息
func (l *logic) UserInfoByID(ctx context.Context, id int64) (*model.UserModel, error) {
	return l.userinfo(ctx, id)
}

// UserLogout 用户登出
func (l *logic) UserLogout(ctx context.Context, id int64) error {
	if err := l.rdb.Del(ctx, constvar.BuildUserTokenKey(id)).Err(); err != nil {
		return err
	}

	return nil
}

// generateToken 生成token
func (l *logic) generateToken(ctx context.Context, uid int64) (string, error) {
	// 签名
	payload := map[string]any{"user_id": uid}
	token, err := auth.Sign(ctx, payload, config.Cfg.Secret.JwtSecret, config.Cfg.Secret.JwtTimeout)
	if err != nil {
		return "", errors.Wrapf(err, "[logic.user] gen token sign")
	}

	// 设置新令牌，用户单点登录
	if err = l.rdb.Set(ctx, constvar.BuildUserTokenKey(uid), token, time.Duration(config.Cfg.Secret.JwtTimeout)*time.Second).Err(); err != nil {
		return "", errors.Wrapf(err, "[logic.user] set token to redis")
	}
	return token, nil
}

// userinfo 获取用户模型
func (l *logic) userinfo(ctx context.Context, id int64) (*model.UserModel, error) {
	user, err := l.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[center.user] from repo by id: %d", id)
	}
	if user.ID == 0 {
		return nil, errno.ErrUserNotFound
	}
	if user.Status != model.StatusNormal {
		return nil, errno.ErrUserFrozen
	}
	return user, nil
}

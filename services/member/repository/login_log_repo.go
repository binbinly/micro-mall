package repository

import (
	"context"

	"github.com/pkg/errors"

	"member/model"
)

// CreateLoginLog 创建登录日志
func (r *Repo) CreateLoginLog(ctx context.Context, log *model.LoginLogModel) error {
	err := r.DB.WithContext(ctx).Create(log).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.loginLog] create")
	}

	return nil
}

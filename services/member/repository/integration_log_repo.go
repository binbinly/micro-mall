package repository

import (
	"context"

	"github.com/pkg/errors"

	"member/model"
)

// CreateIntegrationLog 创建积分变化记录
func (r *Repo) CreateIntegrationLog(ctx context.Context, log *model.IntegrationLogModel) error {
	err := r.DB.WithContext(ctx).Create(log).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.integratonLog] create")
	}

	return nil
}

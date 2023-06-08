package repository

import (
	"context"

	"github.com/pkg/errors"

	"member/model"
)

// CreateGrowthLog 创建成长值变化记录
func (r *Repo) CreateGrowthLog(ctx context.Context, log *model.GrowthLogModel) error {
	err := r.DB.WithContext(ctx).Create(log).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.growthLog] create")
	}

	return nil
}

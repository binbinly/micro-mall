package repository

import (
	"context"

	"github.com/pkg/errors"

	"product/model"
)

// CreateSpuComment 创建商品评价
func (r *Repo) CreateSpuComment(ctx context.Context, comments []*model.SpuCommentModel) error {
	err := r.DB.WithContext(ctx).Create(&comments).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.spuComment] batch create")
	}
	return nil
}

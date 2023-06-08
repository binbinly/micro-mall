package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"product/model"
)

// GetSpuDescBySpuID 获取spu介绍
func (r *Repo) GetSpuDescBySpuID(ctx context.Context, spuID int64) (desc *model.SpuDescModel, err error) {
	doKey := fmt.Sprintf("spuDesc:%d", spuID)
	if err = r.QueryCache(ctx, doKey, &desc, 0, func(data any) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Where("spu_id=?", spuID).First(data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.spuDesc] query cache")
	}

	return
}

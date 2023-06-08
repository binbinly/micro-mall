package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"product/model"
)

// GetSpuByID 获取spu信息
func (r *Repo) GetSpuByID(ctx context.Context, id int64) (spu *model.SpuModel, err error) {
	doKey := fmt.Sprintf("spu:%d", id)
	if err = r.QueryCache(ctx, doKey, &spu, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).First(data, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.spu] query cache")
	}

	return
}

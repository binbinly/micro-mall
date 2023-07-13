package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"pkg/dbs"
	"product/model"
)

// GetImagesBySkuID 获取sku图集
func (r *Repo) GetImagesBySkuID(ctx context.Context, skuID int64) (list []*model.SkuImageModel, err error) {
	doKey := fmt.Sprintf("sku_image:%d", skuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data any) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.SkuImageModel{}).Where("sku_id=?", skuID).
			Order(dbs.DefaultOrderSort).Find(data).Error; err != nil {
			return errors.Wrap(err, "[r.skuImage] query db")
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[r.skuImage] query cache")
	}
	return
}

package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"product/model"
)

// GetSkuByID 获取sku信息
func (r *Repo) GetSkuByID(ctx context.Context, id int64) (sku *model.SkuModel, err error) {
	if err = r.QueryCache(ctx, buildSkuCacheKey(id), &sku, 0, func(data any) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).First(data, id).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] query cache")
	}

	return
}

// GetSkusByIds 批量获取ksu信息
func (r *Repo) GetSkusByIds(ctx context.Context, ids []int64) (list []*model.SkuModel, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildSkuCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.SkuModel)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, func() interface{} {
		return &model.SkuModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] multi get sku cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		sku, ok := cacheMap[buildSkuCacheKey(id)]
		if !ok {
			sku, err = r.GetSkuByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.sku] get sku model err: %v", err)
				continue
			}
			if sku.ID == 0 {
				continue
			}
		}
		list = append(list, sku)
	}
	return list, nil
}

// GetSkusBySpuID 获取spu下所有sku
func (r *Repo) GetSkusBySpuID(ctx context.Context, spuID int64) (list []*model.SkuModel, err error) {
	doKey := fmt.Sprintf("skus:%d", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.SkuModel{}).Where("spu_id=?", spuID).Find(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.skuImage] query cache")
	}
	return
}

// buildSkuCacheKey 构建sku模型缓存键
func buildSkuCacheKey(id int64) string {
	return fmt.Sprintf("sku:%d", id)
}

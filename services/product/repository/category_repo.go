package repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"

	"pkg/dbs"
	"product/model"
)

// GetCategoryNameByID 获取分类名
func (r *Repo) GetCategoryNameByID(ctx context.Context, id int64) (name string, err error) {
	if err = r.QueryCache(ctx, buildCategoryCacheKey(id), &name, 0, func(data interface{}) error {
		// 从数据库中获取
		var cat model.CategoryModel
		if err := r.DB.WithContext(ctx).First(&cat, id).Error; err != nil && err != gorm.ErrRecordNotFound {
			return errors.Wrapf(err, "[repo.category] query db")
		}
		reflect.ValueOf(data).Elem().Set(reflect.ValueOf(cat.Name))
		return nil
	}); err != nil {
		return "", errors.Wrapf(err, "[repo.category] query cache")
	}

	return
}

// GetCategoryNamesByIds 批量获取分类名
func (r *Repo) GetCategoryNamesByIds(ctx context.Context, ids []int64) (names map[int64]string, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildCategoryCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]string)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, nil); err != nil {
		return nil, errors.Wrapf(err, "[repo.category] multi get catName cache data err")
	}

	names = make(map[int64]string, len(ids))
	// 查询未命中
	for _, id := range ids {
		name, ok := cacheMap[buildCategoryCacheKey(id)]
		if !ok {
			name, err = r.GetCategoryNameByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.sku] get catName model err: %v", err)
				continue
			}
			if name == "" {
				continue
			}
		}
		names[id] = name
	}
	return names, nil
}

// GetCategoryChild 获取分类下所有子分类
func (r *Repo) GetCategoryChild(ctx context.Context, id int64) ([]int64, error) {
	list, err := r.CategoryAll(ctx)
	if err != nil {
		return nil, err
	}
	return makeChild(id, list), nil
}

// CategoryAll 获取全部分裂
func (r *Repo) CategoryAll(ctx context.Context) (list []*model.CategoryModel, err error) {
	if err = r.QueryCache(ctx, "category_all", &list, 0, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.CategoryModel{}).Order(dbs.DefaultOrderSort).Find(&list).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.category] query cache")
	}
	return
}

// makeChild 递归所有子分类
func makeChild(id int64, list []*model.CategoryModel) (ids []int64) {
	for _, categoryModel := range list {
		if id == 0 { //所有子分类
			ids = append(ids, categoryModel.ID)
			continue
		}
		if categoryModel.ParentID == id { //递归查找
			ids = append(ids, makeChild(categoryModel.ID, list)...)
		}
	}
	if id > 0 {
		ids = append(ids, id)
	}
	return ids
}

// buildCategoryCacheKey 构建分类模型缓存键
func buildCategoryCacheKey(id int64) string {
	return fmt.Sprintf("category:%d", id)
}

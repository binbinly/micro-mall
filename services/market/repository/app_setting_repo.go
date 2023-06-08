package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/pkg/errors"

	"market/model"
	"pkg/mysql"
)

// AppPageData app其他页面配置数据
func (r *Repo) AppPageData(ctx context.Context, page int) (list []*model.AppSettingModel, err error) {
	doKey := fmt.Sprintf("app_page:%d", page)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.AppSettingModel{}).
			Where("page=?", page).Order(mysql.DefaultOrderSort).Scan(&list).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.config] query cache")
	}

	return
}

// AppHomePageData app首页配置数据
func (r *Repo) AppHomePageData(ctx context.Context, catID int) (list []*model.AppSettingModel, err error) {
	doKey := fmt.Sprintf("app_page:%d_%d", model.AppPageHome, catID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data interface{}) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.AppSettingModel{}).
			Where("page=? and cat_id=?", model.AppPageHome, catID).
			Order(mysql.DefaultOrderSort).Scan(&list).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.config] query cache")
	}

	return
}

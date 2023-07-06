package repository

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"market/model"
)

// GetConfigByName 获取配置
func (r *Repo) GetConfigByName(ctx context.Context, name string, v any) (err error) {
	doKey := "config:" + name
	var config *model.Config
	if err = r.QueryCache(ctx, doKey, &config, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.ConfigModel{}).
			Where("name=?", name).First(data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "[repo.config] query cache")
	}

	if config.Value == "" {
		return nil
	}
	if v != nil {
		if err = json.Unmarshal([]byte(config.Value), &v); err != nil {
			return errors.Wrapf(err, "[repo.config] json unmarshal by value: %v", config.Value)
		}
	} else {
		v = config.Value
	}
	return nil
}

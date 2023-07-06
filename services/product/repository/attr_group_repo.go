package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"pkg/dbs"
	"product/model"
)

// GetAttrGroupByCatID 获取当前分类下的所有属性分组
func (r *Repo) GetAttrGroupByCatID(ctx context.Context, catID int64) (list []*model.AttrGroupModel, err error) {
	doKey := fmt.Sprintf("attr_group:%d", catID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data any) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.AttrGroupModel{}).Where("cat_id=?", catID).
			Order(dbs.DefaultOrderSort).Find(&list).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.attrGroup] query cache")
	}
	return
}

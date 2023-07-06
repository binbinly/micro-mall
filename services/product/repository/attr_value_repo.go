package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"pkg/dbs"
	"product/model"
)

// GetAttrsBySpuID 批量获取spu下所有属性信息
func (r *Repo) GetAttrsBySpuID(ctx context.Context, spuID int64) (list []*model.Attrs, err error) {
	doKey := fmt.Sprintf("attrs:%v", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.AttrValueModel{}).Select("pms_attr_value.*, g.group_id").
			Joins("left join pms_attr_rel_group as g on pms_attr_value.attr_id=g.attr_id").Where("spu_id=?", spuID).
			Order(dbs.DefaultOrderSort).Scan(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.attrValue] query cache")
	}
	return
}

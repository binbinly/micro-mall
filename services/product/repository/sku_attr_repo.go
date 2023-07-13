package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"pkg/dbs"
	"product/model"
)

// GetSkuAttrsBySpuID 批量获取spu下的所有销售属性
func (r *Repo) GetSkuAttrsBySpuID(ctx context.Context, spuID int64) (list []*model.SkuAttrModel, err error) {
	doKey := fmt.Sprintf("skuAttrs:%v", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data any) error {
		// 从数据库中获取
		// EXPLAIN select a.* from pms_sku_attr as a LEFT JOIN pms_sku as s on s.id=a.sku_id where s.spu_id=17664;
		// EXPLAIN select * from pms_sku_attr where sku_id in (select id from pms_sku where spu_id=17664);
		// 此处使用子查询
		if err := r.DB.WithContext(ctx).Model(&model.SkuAttrModel{}).
			Where("sku_id in (select id from pms_sku where spu_id=?)", spuID).
			Order(dbs.DefaultOrderSort).Find(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[r.skuAttr] query cache")
	}
	return
}

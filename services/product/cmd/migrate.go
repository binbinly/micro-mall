package cmd

import (
	"pkg/dbs"
	"product/config"
	"product/model"
)

// Migrate 数据迁移
func Migrate() {
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			new(model.AttrModel),
			new(model.AttrGroupModel),
			new(model.AttrRelGroupModel),
			new(model.BrandModel),
			new(model.CategoryModel),
			new(model.CategoryRelBrandModel),
			new(model.AttrValueModel),
			new(model.SkuAttrModel),
			new(model.SkuImageModel),
			new(model.SkuModel),
			new(model.SpuCommentModel),
			new(model.SpuDescModel),
			new(model.SpuImageModel),
			new(model.SpuModel),
		}
	})
}

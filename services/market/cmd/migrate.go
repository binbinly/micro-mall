package cmd

import (
	"market/config"
	"market/model"
	"pkg/dbs"
)

// Migrate 数据迁移
func Migrate() {
	// 初始化数据库链接
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			new(model.ConfigModel),
			new(model.AppNoticeModel),
			new(model.AppSettingModel),
			new(model.CouponModel),
			new(model.CouponMemberModel),
			new(model.CouponRelCatModel),
			new(model.CouponRelSpuModel),
			new(model.HomeAdvModel),
			new(model.HomeSubjectModel),
			new(model.HomeSubjectSpuModel),
			new(model.MemberPriceModel),
			new(model.SeckillActivityModel),
			new(model.SeckillSkuModel),
			new(model.SeckillSessionModel),
			new(model.SeckillSkuNoticeModel),
			new(model.SkuFullReductionModel),
			new(model.SkuLadderModel),
			new(model.SpuBoundsModel),
		}
	})
}

package cmd

import (
	"member/config"
	"member/model"
	"pkg/dbs"
)

// Migrate 数据迁移
func Migrate() {
	// 初始化数据库链接
	dbs.Migrate(&config.Cfg.MySQL, func() []any {
		return []any{
			new(model.GrowthLogModel),
			new(model.IntegrationLogModel),
			new(model.MemberAddressModel),
			new(model.CollectSpuModel),
			new(model.CollectSubjectModel),
			new(model.MemberLevelModel),
			new(model.LoginLogModel),
			new(model.MemberModel),
			new(model.MemberStatModel),
		}
	})
}

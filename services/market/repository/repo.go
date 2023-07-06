package repository

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"gorm.io/gorm"

	"market/model"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	ICoupon

	GetConfigByName(ctx context.Context, name string, v any) (err error)
	AppPageData(ctx context.Context, page int) (list []*model.AppSettingModel, err error)
	AppHomePageData(ctx context.Context, catID int) (list []*model.AppSettingModel, err error)
	GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error)

	Close() error
}

// Repo struct
type Repo struct {
	repo.Repo
}

// New new a Dao and return
func New(db *gorm.DB, c cache.Cache) IRepo {
	return &Repo{repo.Repo{
		DB:    db,
		Cache: c,
	}}
}

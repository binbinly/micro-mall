package repository

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"

	"gorm.io/gorm"

	"product/model"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	CategoryAll(ctx context.Context) (list []*model.CategoryModel, err error)
	GetCategoryNameByID(ctx context.Context, id int64) (name string, err error)
	GetCategoryNamesByIds(ctx context.Context, ids []int64) (list map[int64]string, err error)
	GetBrandByID(ctx context.Context, id int64) (brand *model.Brand, err error)
	GetBrandsByIds(ctx context.Context, ids []int64) (list map[int64]*model.Brand, err error)
	GetCategoryChild(ctx context.Context, id int64) ([]int64, error)
	GetAttrGroupByCatID(ctx context.Context, catID int64) (list []*model.AttrGroupModel, err error)
	GetAttrsBySpuID(ctx context.Context, spuID int64) (list []*model.Attrs, err error)
	GetSkuAttrsBySpuID(ctx context.Context, spuID int64) (list []*model.SkuAttrModel, err error)
	GetSpuByID(ctx context.Context, id int64) (spu *model.SpuModel, err error)
	GetSkusByIds(ctx context.Context, ids []int64) (list []*model.SkuModel, err error)
	GetSkuByID(ctx context.Context, id int64) (sku *model.SkuModel, err error)
	GetSkusBySpuID(ctx context.Context, spuID int64) (list []*model.SkuModel, err error)
	GetImagesBySkuID(ctx context.Context, skuID int64) (list []*model.SkuImageModel, err error)
	GetSpuDescBySpuID(ctx context.Context, spuID int64) (desc *model.SpuDescModel, err error)
	CreateSpuComment(ctx context.Context, comments []*model.SpuCommentModel) error

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

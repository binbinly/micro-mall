package logic

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"pkg/dbs"
	"pkg/errno"
	"product/model"
)

// CategoryTree 获取产品分类树结构
func (l *logic) CategoryTree(ctx context.Context) ([]*model.CategoryModel, error) {
	return l.repo.CategoryAll(ctx)
}

// SkuDetail sku商品详情
func (l *logic) SkuDetail(ctx context.Context, id int64) (*model.Sku, error) {
	//1, sku基本信息获取
	sku, err := l.skuInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &model.Sku{Sku: sku}

	// 并发执行
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		//1，获取当前spu下所有sku
		skus, err := l.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] skus by id: %v", sku.SpuID)
		}
		res.Skus = skus
		return nil
	})
	g.Go(func() error {
		//2, sku图片信息
		skuImages, err := l.repo.GetImagesBySkuID(ctx, id)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] sku images by id: %v", id)
		}
		res.SkuImages = skuImages
		return nil
	})
	g.Go(func() error {
		//3, sku分类下的属性分组
		attrGroups, err := l.repo.GetAttrGroupByCatID(ctx, sku.CatID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] attr groups by id: %v", sku.CatID)
		}
		res.AttrGroups = attrGroups
		return nil
	})
	g.Go(func() error {
		//4, 获取spu的详情
		spu, err := l.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return errno.ErrProductNotFound
		}
		res.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取sku的销售属性组合
		skuAttrs, err := l.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu skus by id: %v", sku.SpuID)
		}
		res.SkuAttrs = skuAttrs
		return nil
	})
	g.Go(func() error {
		//6, 获取spu介绍
		spuDesc, err := l.repo.GetSpuDescBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu desc by id: %v", sku.SpuID)
		}
		res.SpuDesc = spuDesc
		return nil
	})
	g.Go(func() error {
		//7, 获取spu的规格参数
		spuAttrs, err := l.repo.GetAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu attrs by id: %v", sku.SpuID)
		}
		res.SpuAttrs = spuAttrs
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetSkuSaleAttrs 获取sku销售属性
func (l *logic) GetSkuSaleAttrs(ctx context.Context, id int64) (*model.Sku, error) {
	// sku基本信息获取
	sku, err := l.skuInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &model.Sku{Sku: sku}

	//并发执行
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// 获取当前spu下所有sku
		skus, err := l.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] skus by id: %v", sku.SpuID)
		}
		res.Skus = skus
		return nil
	})
	g.Go(func() error {
		// 获取spu的详情
		spu, err := l.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return errno.ErrProductNotFound
		}
		res.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取spu的销售属性组合
		skuAttrs, err := l.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[product.sku] spu skus by id: %v", sku.SpuID)
		}
		res.SkuAttrs = skuAttrs
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

// GetSkuInfo 获取sku信息
func (l *logic) GetSkuInfo(ctx context.Context, id int64) (*model.SkuModel, error) {
	return l.skuInfo(ctx, id)
}

// SpuComment 商品评价
func (l *logic) SpuComment(ctx context.Context, skuIds []int64, memberID, orderID int64, star int8, content, resources string) error {
	//sku基本信息获取
	skus, err := l.repo.GetSkusByIds(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[product.sku] sku info by ids: %v", skuIds)
	}
	if len(skus) == 0 {
		return errno.ErrProductNotFound
	}
	comments := make([]*model.SpuCommentModel, 0, len(skus))
	for _, sku := range skus {
		comments = append(comments, &model.SpuCommentModel{
			Spu:       dbs.Spu{SpuID: sku.SpuID},
			Sku:       dbs.Sku{SkuID: sku.ID},
			SkuName:   sku.Name,
			MemberID:  memberID,
			ReplyID:   0,
			OrderID:   orderID,
			Star:      star,
			SkuAttrs:  sku.AttrValue,
			Resources: resources,
			Content:   content,
			Release:   dbs.Release{IsRelease: 1},
		})
	}

	if err = l.repo.CreateSpuComment(ctx, comments); err != nil {
		return errors.Wrapf(err, "[product.sku] batch comment")
	}
	return nil
}

// skuInfo 获取sku信息
func (l *logic) skuInfo(ctx context.Context, id int64) (*model.SkuModel, error) {
	//1, sku基本信息获取
	sku, err := l.repo.GetSkuByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[product.sku] sku info by id: %v", id)
	}
	if sku == nil || sku.ID == 0 {
		return nil, errno.ErrProductNotFound
	}
	return sku, nil
}

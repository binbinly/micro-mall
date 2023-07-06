package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"market/model"
	"pkg/constvar"
)

// GetHomeCat 首页分类
func (l *logic) GetHomeCat(ctx context.Context) ([]*model.ConfigHomeCat, error) {
	var cats []*model.ConfigHomeCat
	if err := l.repo.GetConfigByName(ctx, model.ConfigKeyHomeCat, &cats); err != nil {
		return nil, errors.Wrapf(err, "[logic.page] get home cat")
	}

	return cats, nil
}

// GetHomeCatData 首页分类下的配置数据
func (l *logic) GetHomeCatData(ctx context.Context, cid int) ([]*model.AppSettingModel, error) {
	data, err := l.repo.AppHomePageData(ctx, cid)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.page] home data by cid: %v", cid)
	}
	return data, nil
}

// GetNoticeList 公告列表
func (l *logic) GetNoticeList(ctx context.Context, page int32) ([]*model.AppNoticeModel, error) {
	list, err := l.repo.GetNoticeList(ctx, constvar.GetPageOffset(page), constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.page] notice list page: %v", page)
	}
	return list, nil
}

// GetSearchData 搜索页配置数据
func (l *logic) GetSearchData(ctx context.Context) ([]*model.AppSettingModel, []string, error) {
	data, err := l.repo.AppPageData(ctx, model.AppPageSearch)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[logic.page] search data")
	}
	// 搜索热词
	hot, err := l.rdb.ZRevRangeByScore(ctx, constvar.HotSearchKey, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[logic.page] hot keyword by redis")
	}
	return data, hot, nil
}

// GetPayConfig 支付配置
func (l *logic) GetPayConfig(ctx context.Context) ([]*model.ConfigPayList, error) {
	var pays []*model.ConfigPayList
	if err := l.repo.GetConfigByName(ctx, model.ConfigKeyPayList, &pays); err != nil {
		return nil, errors.Wrapf(err, "[logic.page] get pay config")
	}
	return pays, nil
}

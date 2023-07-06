package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"pkg/errno"
	"seckill/repository"
)

// SeckillCheck 秒杀验证
func (l *logic) SeckillCheck(ctx context.Context, memberID, skuID, num int64, key string) (*repository.SkuModel, error) {
	// 获取当前秒杀商品信息
	sku, err := l.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] sku by id: %v", skuID)
	}
	if sku == nil {
		return nil, errno.ErrKillSkuNotFound
	}
	redisKey := fmt.Sprintf("seckill:user:%d_%d", memberID, sku.ID)
	// 验证时间合法性
	now := time.Now().Unix()
	if sku.StartAt > now || sku.EndAt < now { // 不在秒杀时间段内
		return nil, errno.ErrKillTimeInvalid
	}
	// 校验随机码和商品id
	if sku.Key != key { // 随机码错误
		return nil, errno.ErrKillKeyNotMatch
	}
	// 验证购物数量是否合理
	if sku.Limit < num {
		return nil, errno.ErrKillLimitExceed
	}
	stockKey := "seckill:stock:" + key
	lave, err := l.rdb.DecrBy(ctx, stockKey, num).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] decrby key: %v", stockKey)
	}
	if lave < 0 { // 库存已被秒杀完了
		return nil, errno.ErrKillFinish
	}
	// 验证是否已经购买过，幂等性，如果秒杀成功，写入标识
	ttl := time.Duration(sku.EndAt-now) * time.Second
	is, err := l.rdb.SetNX(ctx, redisKey, "1", ttl).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] setnx key: %v", redisKey)
	}
	if !is { // 已经购买过
		return nil, errno.ErrKillRepeat
	}

	return sku, nil
}

// GetSessionAll 获取所有场次
func (l *logic) GetSessionAll(ctx context.Context) ([]*repository.SessionModel, error) {
	list, err := l.repo.GetSessionAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] getall")
	}
	if len(list) == 0 {
		return []*repository.SessionModel{}, nil
	}
	//加载第一个场次的商品数据
	list[0].Skus, err = l.repo.GetSkusBySessionID(ctx, list[0].ID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] skus by sid: %v", list[0].ID)
	}
	return list, nil
}

// GetSessionSkus 获取场次下所有秒杀商品
func (l *logic) GetSessionSkus(ctx context.Context, sessionID int64) ([]*repository.SkuModel, error) {
	skus, err := l.repo.GetSkusBySessionID(ctx, sessionID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] skus by sid: %v", sessionID)
	}
	return skus, nil
}

// GetSkuInfo 获取秒杀商品信息
func (l *logic) GetSkuInfo(ctx context.Context, skuID int64) (*repository.SkuModel, error) {
	sku, err := l.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.session] sku by id: %v", skuID)
	}
	if sku == nil {
		return nil, errno.ErrKillSkuNotFound
	}
	return sku, nil
}

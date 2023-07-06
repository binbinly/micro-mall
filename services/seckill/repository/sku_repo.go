package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go-micro.dev/v4/logger"
)

const (
	//商品关联场次键
	skuRelSessionKey = "seckill:sku:session"
)

var (
	skuKey = "seckill:skus:%v"
)

// GetSkuByID 获取sku商品信息
func (r *Repo) GetSkuByID(ctx context.Context, skuID int64) (*SkuModel, error) {
	id := parseID(skuID)
	//获取商品所在场次
	sessionID, err := r.redis.HGet(ctx, skuRelSessionKey, id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "[repo.sku] hget session by id: %v", id)
	}
	data, err := r.redis.HGet(ctx, fmt.Sprintf(skuKey, sessionID), id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "[repo.sku] hget sku info by id: %v", id)
	}
	sku := &SkuModel{}
	if err = json.Unmarshal([]byte(data), sku); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] json unmarshal data: %v", data)
	}
	return sku, nil
}

// GetSkusBySessionID 获取场次下所有秒杀商品
func (r *Repo) GetSkusBySessionID(ctx context.Context, sessionID int64) ([]*SkuModel, error) {
	data, err := r.redis.HGetAll(ctx, buildSkuKey(sessionID)).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] getall by sid: %v", sessionID)
	}
	var skus []*SkuModel
	for _, datum := range data {
		if datum == "" {
			continue
		}
		sku := &SkuModel{}
		if err = json.Unmarshal([]byte(datum), sku); err != nil {
			logger.Warnf("[repo.sku] json.unmarshal err: %v", err)
			continue
		}
		skus = append(skus, sku)
	}
	return skus, nil
}

// buildSkuKey
func buildSkuKey(sessionID int64) string {
	return fmt.Sprintf(skuKey, sessionID)
}

func parseID(skuID int64) string {
	return strconv.FormatInt(skuID, 10)
}

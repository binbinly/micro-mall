package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"github.com/binbinly/pkg/storage/redis"
)

// AddCart 添加购物车
func (r *Repo) AddCart(ctx context.Context, userID int64, cart *CartModel) error {
	cart.UTime = time.Now().Unix()
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrap(err, "[repo.cart] json marshal")
	}
	return r.redis.HSet(ctx, BuildCartKey(userID), parseID(cart.SkuID), data).Err()
}

// EditCart 更新购物车
func (r *Repo) EditCart(ctx context.Context, userID int64, oldID int64, cart *CartModel) error {
	cart.UTime = time.Now().Unix()
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrap(err, "[repo.cart] json marshal")
	}
	if oldID == 0 { // 未修改商品
		err = r.redis.HSet(ctx, BuildCartKey(userID), cart.SkuID, data).Err()
	} else {
		pipe := r.redis.Pipeline()
		pipe.HDel(ctx, BuildCartKey(userID), parseID(oldID))
		pipe.HSet(ctx, BuildCartKey(userID), cart.SkuID, data)
		_, err = pipe.Exec(ctx)
	}
	if err != nil {
		return errors.Wrap(err, "[repo.cart] pipe exec")
	}
	return nil
}

// GetCartByID 获取购物车数据
func (r *Repo) GetCartByID(ctx context.Context, userID int64, skuID int64) (*CartModel, error) {
	data, err := r.redis.HGet(ctx, BuildCartKey(userID), parseID(skuID)).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "[repo.cart] hget")
	}
	cart := &CartModel{}
	if err == redis.Nil {
		return cart, nil
	}
	err = json.Unmarshal([]byte(data), cart)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.cart] json unmarshal")
	}
	return cart, nil
}

// GetCartsByIds 批量获取购物车数据
func (r *Repo) GetCartsByIds(ctx context.Context, userID int64, ids []int64) ([]*CartModel, error) {
	data, err := r.redis.HMGet(ctx, BuildCartKey(userID), parseIds(ids)...).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "[repo.cart] hmget db")
	}
	var carts []*CartModel
	for _, datum := range data {
		if datum == nil {
			continue
		}
		cart := &CartModel{}
		err = json.Unmarshal([]byte(datum.(string)), cart)
		if err != nil {
			logger.Warnf("[repo.cart] json.unmarshal err: %v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

// DelCart 移除购物车
func (r *Repo) DelCart(ctx context.Context, userID int64, ids []int64) error {
	return r.redis.HDel(ctx, BuildCartKey(userID), parseIds(ids)...).Err()
}

// EmptyCart 清空购物车
func (r *Repo) EmptyCart(ctx context.Context, userID int64) error {
	return r.redis.Del(ctx, BuildCartKey(userID)).Err()
}

// CartList 我的购物车
func (r *Repo) CartList(ctx context.Context, userID int64) ([]*CartModel, error) {
	data, err := r.redis.HGetAll(ctx, BuildCartKey(userID)).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrapf(err, "[repo.cart] json marshal")
	}
	carts := make([]*CartModel, 0, len(data))
	for _, v := range data {
		cart := &CartModel{}
		err = json.Unmarshal([]byte(v), cart)
		if err != nil {
			logger.Warnf("[repo.cart] list json unmarshal err:%v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func parseID(skuID int64) string {
	return strconv.FormatInt(skuID, 10)
}

func parseIds(skuIds []int64) (res []string) {
	for _, id := range skuIds {
		res = append(res, parseID(id))
	}
	return
}

package logic

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"cart/repository"
)

// AddCart 添加购物车
func (l *logic) AddCart(ctx context.Context, userID int64, num int, sku *repository.Sku) error {
	// 获取当前sku_id 的购物车信息
	cart, err := l.repo.GetCartByID(ctx, userID, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[logic.cart] cart uid: %v id: %v", userID, sku.ID)
	}

	if cart.SkuID != 0 { //商品已存在,加数量即可
		cart.Num += num
	} else {
		setCart(sku, num, cart)
	}

	if err = l.repo.AddCart(ctx, userID, cart); err != nil {
		return errors.Wrapf(err, "[logic.cart] add cart")
	}

	return nil
}

// EditCart 更新购物车
func (l *logic) EditCart(ctx context.Context, userID, oldSkuID int64, sku *repository.Sku, num int) error {
	cart, err := l.repo.GetCartByID(ctx, userID, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[logic.cart] get by uid: %v, id: %v", userID, sku.ID)
	}
	if cart.SkuID == sku.ID && cart.Num == num { //购物车未更新，直接返回
		return nil
	}
	if cart.SkuID == 0 { // 商品不存在，直接添加
		setCart(sku, num, cart)
	} else {
		cart.Num = num
	}

	if err = l.repo.EditCart(ctx, userID, oldSkuID, cart); err != nil {
		return errors.Wrapf(err, "[logic.cart] edit sku")
	}

	return nil
}

// EditCartNum 修改商品数量
func (l *logic) EditCartNum(ctx context.Context, userID int64, sku *repository.Sku, num int) error {
	cart, err := l.repo.GetCartByID(ctx, userID, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[logic.cart] get by uid: %v, id: %v", userID, sku.ID)
	}
	if cart.SkuID == 0 { // 商品不存在，直接添加
		setCart(sku, num, cart)
	}
	if cart.Num == num { // 数量未修改，直接返回
		return nil
	} else {
		cart.Num = num
	}

	if err = l.repo.EditCart(ctx, userID, 0, cart); err != nil {
		return errors.Wrapf(err, "[logic.cart] edit num")
	}

	return nil
}

// DelCart 删除购物车
func (l *logic) DelCart(ctx context.Context, userID int64, skuIds []int64) error {
	return l.repo.DelCart(ctx, userID, skuIds)
}

// ClearCart 清空购物车
func (l *logic) ClearCart(ctx context.Context, userID int64) error {
	return l.repo.EmptyCart(ctx, userID)
}

// CartList 购物车
func (l *logic) CartList(ctx context.Context, userID int64) ([]*repository.CartModel, error) {
	list, err := l.repo.CartList(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.cart] list uid: %v", userID)
	}
	sort.Sort(repository.CartSort(list))

	return list, nil
}

// BatchGetCarts 批量获取购物车信息
func (l *logic) BatchGetCarts(ctx context.Context, userID int64, skuIds []int64) ([]*repository.CartModel, error) {
	return l.repo.GetCartsByIds(ctx, userID, skuIds)
}

func setCart(sku *repository.Sku, num int, cart *repository.CartModel) {
	cart.SkuID = sku.ID
	cart.Num = num
	cart.Cover = sku.Cover
	cart.Price = int(sku.Price)
	cart.Title = sku.Title
	cart.SkuAttr = sku.AttrValue
}

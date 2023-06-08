package logic

import (
	"context"
	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/emptypb"
	"pkg/app"
	"pkg/constvar"
	"pkg/errno"
	"pkg/message"
	"pkg/mysql"
	"pkg/proto/cart"
	cpb "pkg/proto/common"
	"pkg/proto/core"
	"pkg/proto/market"
	"pkg/proto/member"
	"pkg/proto/product"
	"pkg/proto/warehouse"
	"strings"
	"time"

	"github.com/pkg/errors"
	"order/model"
)

// OrderSubmit 购物车提交订单
func (l *logic) OrderSubmit(ctx context.Context, memberID, addressID, couponID int64, skuIds []int64, note string) error {
	// 前往购物车服务获取需要购买的商品信息
	carts, err := l.cartService.BatchGetCarts(ctx, &cart.SkusReq{
		SkuIds: skuIds,
	})
	if err != nil {
		return err
	}
	if len(carts.Data) == 0 {
		return errno.ErrOrderSkuEmpty
	}

	// 商品总金额
	var total int
	items := make([]*model.OrderItemModel, 0, len(carts.Data))
	for _, c := range carts.Data {
		items = append(items, buildOrderItem(int64(c.SkuId), int(c.Price), int(c.Num), c.Title, c.Cover, c.SkuAttr))
		total += int(c.Price) * int(c.Num)
	}

	order, err := l.buildOrder(ctx, addressID, memberID, couponID, total, note, "")
	if err != nil {
		return err
	}

	if err = l.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	// 购物车服务删除已购买的sku，不抛出错误，只记录日志
	if _, err = l.cartService.BatchDelCarts(ctx, &cart.SkusReq{SkuIds: skuIds}); err != nil {
		logger.Warnf("[handler.order] del cart by ids: %v err: %v", skuIds, err)
	}

	return nil
}

// SubmitSkuOrder 不通过购物车，直接在商品页面提交sku订单
func (l *logic) SubmitSkuOrder(ctx context.Context, memberID, skuID, addressID, couponID int64, num int, note string) error {
	// 产品服务获取sku信息
	sku, err := l.productService.GetSkuByID(ctx, &cpb.SkuIDReq{Id: skuID})
	if err != nil {
		return err
	}

	total := int(sku.Sku.Price) * num
	items := []*model.OrderItemModel{
		buildOrderItem(sku.Sku.Id, int(sku.Sku.Price), num, sku.Sku.Title, sku.Sku.Cover, sku.Sku.AttrValue),
	}

	order, err := l.buildOrder(ctx, addressID, memberID, couponID, total, note, "")
	if err != nil {
		return err
	}

	if err = l.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	return nil
}

// SubmitKillOrder 秒杀订单提交
func (l *logic) SubmitKillOrder(ctx context.Context, memberID, skuID, addressID int64, price, num int, orderNo string) error {
	// 产品服务获取sku信息
	sku, err := l.productService.GetSkuByID(ctx, &cpb.SkuIDReq{Id: skuID})
	if err != nil {
		return err
	}

	total := price * num
	items := []*model.OrderItemModel{
		buildOrderItem(sku.Sku.Id, price, num, sku.Sku.Title, sku.Sku.Cover, sku.Sku.AttrValue),
	}

	order, err := l.buildOrder(ctx, addressID, memberID, 0, total, "", orderNo)
	if err != nil {
		return err
	}

	if err = l.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	return nil
}

// OrderDetail 订单详情
func (l *logic) OrderDetail(ctx context.Context, mid, id int64) (*model.OrderModel, error) {
	order, err := l.repo.GetOrderDetail(ctx, id, mid)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, errno.ErrOrderNotFound
	}

	return order, nil
}

// OrderCancel 取消订单
func (l *logic) OrderCancel(ctx context.Context, mid, id int64) error {
	order, err := l.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	if order.MemberID != mid || order.Status != model.OrderStatusInit {
		return errno.ErrOrderNotFound
	}

	if err = l.repo.OrderDelete(ctx, order); err != nil {
		return errors.Wrapf(err, "[logic.order] delete")
	}

	// 订单取消，发送订单消息，其他服务订阅处理
	if err = l.releaseEvent.Publish(ctx, &message.OrderMessage{
		OrderId: id,
		Finish:  false,
	}); err != nil {
		logger.Warnf("[logic.order] push release event err: %v", err)
	}

	return nil
}

// MyOrderList 订单列表
func (l *logic) MyOrderList(ctx context.Context, mid int64, status int, page int32) ([]*model.OrderModel, error) {
	return l.repo.GetOrderList(ctx, mid, status, constvar.GetPageOffset(page), constvar.DefaultLimit)
}

// OrderPayNotify 支付成功回调处理
func (l *logic) OrderPayNotify(ctx context.Context, mid int64, amount int, pType int64, orderNo, tradeNo, transHash string) error {
	order, err := l.repo.GetOrderByNo(ctx, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[logic.order] get by no: %v", orderNo)
	}
	if order.MemberID != mid || order.Status != model.OrderStatusInit {
		return errno.ErrOrderNotFound
	}

	// 营销服务获取支付配置
	pays, err := l.marketService.GetPayConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	// 获取当前支付的以太坊合约地址
	var address string
	for _, pay := range pays.Data {
		if pay.Id == int32(pType) {
			address = pay.Address
		}
	}
	if address == "" {
		return errno.ErrPayActionInvalid
	}

	// 第三方服务中检查 当前订单号是否已支付
	if _, err = l.thirdService.CheckETHPay(ctx, &core.ETHPayReq{
		Id:      pType,
		Address: address,
		OrderNo: orderNo,
	}); err != nil {
		return err
	}

	order.Status = model.OrderStatusDelivered
	order.PayAt = time.Now().Unix()
	order.PayType = int8(pType)
	order.PayAmount = amount
	order.TradeNo = tradeNo
	order.TransHash = transHash
	if err = l.repo.OrderSave(ctx, nil, order); err != nil {
		return errors.Wrap(err, "[logic.order] save")
	}

	return nil
}

// OrderConfirmReceipt 确认收货
func (l *logic) OrderConfirmReceipt(ctx context.Context, mid, id int64) error {
	order, err := l.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已支付，已发货，才可以确认收货
	if order.MemberID != mid || (order.Status != model.OrderStatusDelivered && order.Status != model.OrderStatusShipped) {
		return errno.ErrOrderNotFound
	}

	order.Status = model.OrderStatusReceived
	order.ReceiveAt = time.Now().Unix()
	order.IsConfirm = 1
	if err = l.repo.OrderSave(ctx, nil, order); err != nil {
		return errors.Wrapf(err, "[logic.order] save")
	}

	return nil
}

// OrderRefund 退款
func (l *logic) OrderRefund(ctx context.Context, mid, id int64, content string) error {
	order, err := l.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已支付，已收货状态下方可退款
	if order.MemberID != mid || (order.Status != model.OrderStatusReceived && order.Status != model.OrderStatusDelivered) {
		return errno.ErrOrderNotFound
	}

	// 开启事务
	tx := l.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusPendingRefund
	if err = l.repo.OrderSave(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[logic.order] save")
	}

	if err = l.repo.CreateOrderRefund(ctx, tx, &model.OrderRefundModel{
		OID:     model.OID{OrderID: order.ID},
		Amount:  order.PayAmount,
		Content: content,
	}); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[logic.order] create refund")
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.order] tx commit refund")
	}

	return nil
}

// OrderComment 评价
func (l *logic) OrderComment(ctx context.Context, mid, id int64, skuIds []int64, star int32, content, resources string) error {
	order, err := l.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已收货方可评价
	if order.MemberID != mid || order.Status != model.OrderStatusReceived {
		return errno.ErrOrderNotFound
	}

	// 开启事务
	tx := l.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusFinish
	order.CommentAt = time.Now().Unix()
	if err = l.repo.OrderSave(ctx, tx, order); err != nil {
		return errors.Wrapf(err, "[logic.order] save")
	}

	// 产品服务 add sku comment record
	if _, err = l.productService.SpuComment(ctx, &product.CommentReq{
		SkuIds:    skuIds,
		OrderId:   id,
		Star:      star,
		Content:   content,
		Resources: resources,
	}); err != nil {
		return err
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.order] tx commit refund")
	}

	// 订单完成，发送订单完成任务，其他服务监听异步处理，比如库存服务监听扣减库存
	if err = l.releaseEvent.Publish(ctx, &message.OrderMessage{
		OrderId: id,
		Finish:  true,
	}); err != nil {
		logger.Warnf("[logic.order] push release event err: %v", err)
	}

	return nil
}

// OrderInfo 订单详情
func (l *logic) OrderInfo(ctx context.Context, id int64) (*model.OrderModel, error) {
	return l.orderInfo(ctx, id)
}

func (l *logic) orderInfo(ctx context.Context, id int64) (*model.OrderModel, error) {
	order, err := l.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, errno.ErrOrderNotFound
	}
	return order, nil
}

// orderSubmit 商品不通过购物车，直接提交订单
// 分布式事务 - 柔性事务-可靠消息 + 最终一致性方案->rabbitmq延迟队列
func (l *logic) orderSubmit(ctx context.Context, order *model.OrderModel, items []*model.OrderItemModel) error {
	// 开启事务
	tx := l.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	if err := l.repo.CreateOrder(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[logic.order] create")
	}

	// 批量创建子订单
	if err := l.repo.BatchCreateOrderItem(ctx, tx, items); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[logic.order] create items")
	}

	if err := l.submitBefore(ctx, order, items); err != nil {
		tx.Rollback()
		return err
	}

	//提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		// 提交订单失败了，库存已锁定成功，需要告知其他服务回滚，其他服务自行订阅 topic
		if err = l.releaseEvent.Publish(ctx, &message.OrderMessage{
			OrderId: order.ID,
			Finish:  false,
		}); err != nil {
			logger.Warnf("[logic.order] push release event err: %v", err)
		}
		return errors.Wrap(err, "[logic.order] tx commit")
	}

	// 创建订单成功，发送延迟队列，自动取消订单
	if err := l.createEvent.Publish(ctx, &message.OrderCancelMessage{
		OrderID:  order.ID,
		MemberID: order.MemberID,
	}); err != nil {
		logger.Warnf("[logic.order] push create event err: %v", err)
	}

	return nil
}

// submitBefore 订单提交之前执行其他操作，锁定库存，优惠券使用
func (l *logic) submitBefore(ctx context.Context, order *model.OrderModel, items []*model.OrderItemModel) error {
	if order.CouponID > 0 { // 设置优惠券已使用
		if _, err := l.marketService.CouponUsed(ctx, &market.CouponUsedReq{
			CouponId: order.CouponID,
			OrderId:  order.ID,
		}); err != nil {
			return err
		}
	}

	skuNums := getSkuNums(order, items)
	// 仓储服务 锁定库存
	if _, err := l.warehouseService.SKuStockLock(ctx, &warehouse.SkuStockLockReq{
		OrderId:   0,
		OrderNo:   order.OrderNo,
		Consignee: order.AddressName,
		Phone:     order.AddressPhone,
		Address:   strings.Join([]string{order.AddressProvince, order.AddressCity, order.AddressCounty, order.AddressDetail}, " "),
		Note:      order.Note,
		SkuNums:   skuNums,
	}); err != nil {
		return err
	}

	return nil
}

// buildOrder 生成订单
// 1，订单信息
// 2，商品spu信息（暂不处理）
// 3，商品sku信息
// 4，优惠信息
// 5，积分信息,与价格同步 1元 = 1积分，1成长值
func (l *logic) buildOrder(ctx context.Context, addressId, memberID, couponID int64, total int, note, orderNo string) (*model.OrderModel, error) {
	// 生成订单号
	if orderNo == "" {
		orderNo = app.GenOrderNo()
	}

	// 会员服务获取收货地址信息
	address, err := l.memberService.GetAddressInfo(ctx, &member.AddressIDReq{Id: addressId})
	if err != nil {
		return nil, err
	}

	// 实际应该支付总金额
	amount := total
	// 优惠券优惠金额
	var couponAmount int
	if couponID > 0 {
		// 营销服务获取优惠券信息
		coupon, err := l.marketService.GetCouponInfo(ctx, &market.CouponIDReq{Id: couponID})
		if err != nil {
			return nil, err
		}
		couponAmount = int(coupon.Coupon.Amount)
		amount -= couponAmount
	}

	return &model.OrderModel{
		MID:             mysql.MID{MemberID: memberID},
		OrderNo:         orderNo,
		CouponID:        couponID,
		TotalAmount:     total,
		CouponAmount:    couponAmount,
		Amount:          amount,
		SourceType:      model.OrderSourceTypeApp,
		Status:          model.OrderStatusInit,
		AutoConfirmDay:  15,
		Integration:     total / 100,
		Growth:          total / 100,
		AddressName:     address.Address.Name,
		AddressPhone:    address.Address.Phone,
		AddressProvince: address.Address.Province,
		AddressCity:     address.Address.City,
		AddressCounty:   address.Address.County,
		AddressDetail:   address.Address.Detail,
		Note:            note,
	}, nil
}

// buildOrderItem 构建 order items
// 积分信息,与价格同步 1元 = 1积分，1成长值
func buildOrderItem(skuId int64, price, num int, title, cover, skuAttrs string) *model.OrderItemModel {
	amount := price * num
	return &model.OrderItemModel{
		Sku:             mysql.Sku{SkuID: skuId},
		SkuName:         title,
		SkuImg:          cover,
		SkuPrice:        amount,
		SkuAttrs:        skuAttrs,
		Num:             num,
		RealAmount:      amount,
		GiveIntegration: amount / 100,
		GiveGrowth:      amount / 100,
	}
}

func getSkuNums(order *model.OrderModel, items []*model.OrderItemModel) map[int64]int32 {
	skuNums := make(map[int64]int32, len(items))
	// 子订单批量创建
	for _, item := range items {
		item.OrderID = order.ID
		item.OrderNo = order.OrderNo
		skuNums[item.SkuID] = int32(item.Num)
	}

	return skuNums
}

package errno

import (
	"go-micro.dev/v4/errors"
	"pkg/constvar"
)

// 对外错误响应定义
// 全局错误定义
var (
	// ErrParamsCheckInvalid 参数验证不通过
	ErrParamsCheckInvalid = errors.New("global", "参数非法", 400)
)

// MarketReplyErr 营销服错误响应处理
func MarketReplyErr(err error) error {
	switch err {
	case ErrCouponFinished:
		return errors.New(constvar.ServiceMarket, "优惠券已领完", int32(CouponFinished))
	case ErrCouponReceived:
		return errors.New(constvar.ServiceMarket, "优惠券已领取过哦", int32(CouponReceived))
	case ErrCouponNotFound:
		return errors.New(constvar.ServiceMarket, "优惠券不存在", int32(CouponNotFound))
	}
	return err
}

// ProductReplyErr 产品服错误响应处理
func ProductReplyErr(err error) error {
	switch err {
	case ErrProductNotFound:
		return errors.New(constvar.ServiceProduct, "商品不存在哦", int32(ProductNotFound))
	}
	return err
}

// MemberReplyErr 会员服错误响应处理
func MemberReplyErr(err error) error {
	switch err {
	case ErrMemberPhoneValid:
		return errors.New(constvar.ServiceMember, "手机号格式错误", int32(MemberPhoneValid))
	case ErrMemberAddressNotFound:
		return errors.New(constvar.ServiceMember, "收货地址不存在", int32(MemberAddressNotFound))
	case ErrMemberFrozen:
		return errors.New(constvar.ServiceMember, "账号已被冻结，如有疑问，请联系客服", int32(MemberFrozen))
	case ErrMemberNotFound:
		return errors.New(constvar.ServiceMember, "账号不存在", int32(MemberNotFound))
	case ErrMemberExisted:
		return errors.New(constvar.ServiceMember, "账号已存在", int32(MemberExisted))
	}
	return err
}

// OrderReplyErr 订单服错误响应处理
func OrderReplyErr(err error) error {
	switch err {
	case ErrOrderNotFound:
		return errors.New(constvar.ServiceOrder, "订单不存在", int32(OrderNotFound))
	case ErrOrderSkuEmpty:
		return errors.New(constvar.ServiceOrder, "商品不存在", int32(OrderSkuEmpty))
	case ErrPayActionInvalid:
		return errors.New(constvar.ServiceOrder, "无效的支付方式", int32(PayActionInvalid))
	}
	return err
}

// SeckillReplyErr 秒杀服务错误响应处理
func SeckillReplyErr(err error) error {
	switch err {
	case ErrKillTimeInvalid:
		return errors.New(constvar.ServiceSeckill, "不在秒杀时间内哦", int32(KillTimeInvalid))
	case ErrKillKeyNotMatch:
		return errors.New(constvar.ServiceSeckill, "商品令牌错误", int32(KillKeyNotMatch))
	case ErrKillLimitExceed:
		return errors.New(constvar.ServiceSeckill, "超出数量限制哦", int32(KillLimitExceed))
	case ErrKillRepeat:
		return errors.New(constvar.ServiceSeckill, "不可重复秒杀哦", int32(KillRepeat))
	case ErrKillFinish:
		return errors.New(constvar.ServiceSeckill, "来晚了，该商品已秒杀完了哦", int32(KillFinish))
	case ErrKillSkuNotFound:
		return errors.New(constvar.ServiceSeckill, "商品不存在", int32(KillSkuNotFound))
	}
	return err
}

package errno

import (
	"go-micro.dev/v4/errors"
)

const (
	warehouseID = "warehouse"
	centerID    = "center"
	thirdID     = "third-party"
)

// 内部服务错误定义
// WarehouseReplyErr 仓储服错误响应处理
func WarehouseReplyErr(err error) error {
	switch err {
	case ErrWareInventoryShortage:
		return errors.New(warehouseID, "库存不足", int32(WareInventoryShortage))
	}
	return err
}

// CenterReplyErr 中心服错误响应处理
func CenterReplyErr(err error) error {
	switch err {
	case ErrUserExisted:
		return errors.New(centerID, "用户已存在", int32(UserExisted))
	case ErrUserNotFound:
		return errors.New(centerID, "用户不存在", int32(UserNotFound))
	case ErrUserNotMatch:
		return errors.New(centerID, "用户密码不匹配", int32(UserNotMatch))
	case ErrUserFrozen:
		return errors.New(centerID, "账号已被冻结", int32(UserFrozen))
	case ErrUserTokenExpired:
		return errors.New(centerID, "令牌已过期", int32(UserTokenExpired))
	case ErrUserTokenEmpty:
		return errors.New(centerID, "请先登录", int32(UserTokenEmpty))
	}
	return err
}

// ThirdReplyErr 第三方服务错误响应处理
func ThirdReplyErr(err error) error {
	switch err {
	case ErrVerifyCodeRuleMinute:
		return errors.New(thirdID, "触发分钟流控", int32(VerifyCodeRuleMinute))
	case ErrVerifyCodeRuleHour:
		return errors.New(thirdID, "触发小时流控", int32(VerifyCodeRuleHour))
	case ErrVerifyCodeRuleDay:
		return errors.New(thirdID, "触发天极流控", int32(VerifyCodeRuleDay))
	case ErrVerifyCodeNotMatch:
		return errors.New(thirdID, "验证码不匹配", int32(VerifyCodeNotMatch))
	case ErrETHPayNotFound:
		return errors.New(thirdID, "支付订单未找到", int32(ETHPayNotFound))
	}
	return err
}

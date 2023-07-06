package errno

import (
	"go-micro.dev/v4/errors"
	"pkg/constvar"
)

// 内部服务错误定义

// WarehouseReplyErr 仓储服错误响应处理
func WarehouseReplyErr(err error) error {
	switch err {
	case ErrWareInventoryShortage:
		return errors.New(constvar.ServiceWarehouse, "库存不足", int32(WareInventoryShortage))
	}
	return err
}

// CenterReplyErr 中心服错误响应处理
func CenterReplyErr(err error) error {
	switch err {
	case ErrUserExisted:
		return errors.New(constvar.ServiceCenter, "用户已存在", int32(UserExisted))
	case ErrUserNotFound:
		return errors.New(constvar.ServiceCenter, "用户不存在", int32(UserNotFound))
	case ErrUserNotMatch:
		return errors.New(constvar.ServiceCenter, "用户密码不匹配", int32(UserNotMatch))
	case ErrUserPwd:
		return errors.New(constvar.ServiceCenter, "密码错误", int32(UserPwd))
	case ErrUserFrozen:
		return errors.New(constvar.ServiceCenter, "账号已被冻结", int32(UserFrozen))
	case ErrUserTokenExpired:
		return errors.New(constvar.ServiceCenter, "令牌已过期", int32(UserTokenExpired))
	case ErrUserTokenEmpty:
		return errors.New(constvar.ServiceCenter, "请先登录", int32(UserTokenEmpty))
	}
	return err
}

// ThirdReplyErr 第三方服务错误响应处理
func ThirdReplyErr(err error) error {
	switch err {
	case ErrVerifyCodeRuleMinute:
		return errors.New(constvar.ServiceThird, "触发分钟流控", int32(VerifyCodeRuleMinute))
	case ErrVerifyCodeRuleHour:
		return errors.New(constvar.ServiceThird, "触发小时流控", int32(VerifyCodeRuleHour))
	case ErrVerifyCodeRuleDay:
		return errors.New(constvar.ServiceThird, "触发天极流控", int32(VerifyCodeRuleDay))
	case ErrVerifyCodeNotMatch:
		return errors.New(constvar.ServiceThird, "验证码不匹配", int32(VerifyCodeNotMatch))
	case ErrETHPayNotFound:
		return errors.New(constvar.ServiceThird, "支付订单未找到", int32(ETHPayNotFound))
	}
	return err
}

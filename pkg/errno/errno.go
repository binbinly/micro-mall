package errno

import (
	"github.com/pkg/errors"
)

// 中心服错误定义
var (
	//ErrUserExisted 用户已存在
	ErrUserExisted = errors.New("user:existed")
	//ErrUserNotFound 用户不存在
	ErrUserNotFound = errors.New("user:not found")
	//ErrUserNotMatch 用户名密码不匹配
	ErrUserNotMatch = errors.New("user:not match")
	//ErrUserPwd 密码错误
	ErrUserPwd = errors.New("user:pwd")
	//ErrUserFrozen 账号已冻结
	ErrUserFrozen = errors.New("user:frozen")
	//ErrUserTokenExpired 用户令牌过期
	ErrUserTokenExpired = errors.New("user: token expired")
	//ErrUserTokenEmpty 用户令牌空
	ErrUserTokenEmpty = errors.New("user: token empty")
)

// 营销服错误定义
var (
	//ErrCouponNotFound 优惠券不存在
	ErrCouponNotFound = errors.New("coupon:not found")
	//ErrCouponFinished 优惠券已领完
	ErrCouponFinished = errors.New("coupon:finished")
	//ErrCouponReceived 已领取过
	ErrCouponReceived = errors.New("coupon:received")
)

// 会员服错误定义
var (
	//ErrMemberExisted 用户已存在
	ErrMemberExisted = errors.New("member:existed")
	//ErrMemberNotFound 用户不存在
	ErrMemberNotFound = errors.New("member:not found")
	//ErrMemberFrozen 账号已冻结
	ErrMemberFrozen = errors.New("member:frozen")
	//ErrMemberAddressNotFound 会员收货地址不存在
	ErrMemberAddressNotFound = errors.New("member:address not found")
	//ErrMemberPhoneValid 手机号格式错误
	ErrMemberPhoneValid = errors.New("member:phone valid")
	// ErrMemberPasswordNotMatch 两次输入密码不匹配
	ErrMemberPasswordNotMatch = errors.New("The two passwords do not match")
)

// 订单服错误定义
var (
	//ErrOrderNotFound 订单不存在
	ErrOrderNotFound = errors.New("order:not found")
	//ErrOrderSkuEmpty 订单提交商品为空
	ErrOrderSkuEmpty = errors.New("order sku empty")
	//ErrPayActionInvalid 无效的支付方式
	ErrPayActionInvalid = errors.New("pay action invalid")
)

// 产品服错误定义
var (
	//ErrProductNotFound 产品不存在
	ErrProductNotFound = errors.New("product:not found")
)

// 秒杀服错误定义
var (
	//ErrVerifyCodeRuleMinute 分钟限制
	ErrVerifyCodeRuleMinute = errors.New("sms:minute limit")
	//ErrVerifyCodeRuleHour 小时限制
	ErrVerifyCodeRuleHour = errors.New("sms:hour limit")
	//ErrVerifyCodeRuleDay 天级限制
	ErrVerifyCodeRuleDay = errors.New("sms:day limit")
	//ErrVerifyCodeNotMatch 验证码不匹配
	ErrVerifyCodeNotMatch = errors.New("code:not match")
	//ErrETHPayNotFound 支付未找到
	ErrETHPayNotFound = errors.New("eth:not found")
)

// 仓储服错误定义
var (
	//ErrWareInventoryShortage 库存不足
	ErrWareInventoryShortage = errors.New("ware:inventory shortage")
)

var (
	//ErrKillTimeInvalid 时间无效，不在秒杀时间内
	ErrKillTimeInvalid = errors.New("kill:time invalid")
	//ErrKillKeyNotMatch 随机码不匹配
	ErrKillKeyNotMatch = errors.New("kill:key not match")
	//ErrKillLimitExceed 超出数量限制
	ErrKillLimitExceed = errors.New("kill:limit exceed")
	//ErrKillRepeat 不可重复购买
	ErrKillRepeat = errors.New("kill:repeat")
	//ErrKillFinish 秒杀已结束
	ErrKillFinish = errors.New("kill:finish")
	//ErrKillSkuNotFound 秒杀商品不存在
	ErrKillSkuNotFound = errors.New("kill:sku not found")
)

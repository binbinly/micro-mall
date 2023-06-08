package errno

import "google.golang.org/grpc/codes"

// 中心服错误码定义
const (
	//grpc 返回错误code
	//UserExisted 用户已存在
	UserExisted codes.Code = 100
	//UserNotFound 用户不存在
	UserNotFound codes.Code = 101
	//UserNotMatch 用户账号不匹配
	UserNotMatch codes.Code = 102
	//UserFrozen 用户已冻结
	UserFrozen codes.Code = 103
	//UserTokenExpired 用户令牌已过期
	UserTokenExpired codes.Code = 104
	//UserTokenEmpty 用户令牌空
	UserTokenEmpty codes.Code = 105
	//UserPwd 密码错误
	UserPwd codes.Code = 106
)

// 第三方服务错误码定义
const (
	//grpc 返回错误code
	//VerifyCodeRuleMinute 分钟流控
	VerifyCodeRuleMinute codes.Code = 200
	//VerifyCodeRuleHour 小时流控
	VerifyCodeRuleHour codes.Code = 201
	//VerifyCodeRuleDay 天极流控
	VerifyCodeRuleDay codes.Code = 202
	//VerifyCodeNotMatch 验证码不匹配
	VerifyCodeNotMatch codes.Code = 203
	//ETHPayNotFound 以太币支付不存在
	ETHPayNotFound codes.Code = 204
)

// 仓储服错误码定义
const (
	//grpc 返回错误code
	//WareInventoryShortage 库存不足
	WareInventoryShortage codes.Code = 300
)

// 产品服错误码定义
const (
	//grpc 返回错误code
	//ProductNotFound 产品不存在
	ProductNotFound codes.Code = 1000
)

// 秒杀服错误码定义
const (
	//grpc 返回错误code
	//KillTimeInvalid 时间无效
	KillTimeInvalid codes.Code = 2000
	//KillKeyNotMatch 随机码不匹配
	KillKeyNotMatch codes.Code = 2001
	//KillLimitExceed 超时数量限制
	KillLimitExceed codes.Code = 2002
	//KillRepeat 不可以重复秒杀
	KillRepeat codes.Code = 2003
	//KillFinish 秒杀已结束
	KillFinish codes.Code = 2004
	//KillSkuNotFound 秒杀商品不存在
	KillSkuNotFound codes.Code = 2005
)

// 会员服错误码定义
const (
	//grpc 返回错误code
	//MemberExisted 会员已存在
	MemberExisted codes.Code = 3000
	//MemberNotFound 会员不存在
	MemberNotFound codes.Code = 3001
	//MemberFrozen 会员已冻结
	MemberFrozen codes.Code = 3002
	//MemberAddressNotFound 收货地址不存在
	MemberAddressNotFound codes.Code = 3003
	//MemberPhoneValid 手机号格式错误
	MemberPhoneValid codes.Code = 3004
)

// 营销服错误码定义
const (
	//grpc 返回错误code
	//CouponNotFound 优惠券不存在
	CouponNotFound codes.Code = 4000
	//CouponFinished 优惠券已领完
	CouponFinished codes.Code = 4001
	//CouponReceived 已领取过
	CouponReceived codes.Code = 4002
)

// 订单服错误码定义
const (
	//grpc 返回错误code
	//OrderNotFound 订单不存在
	OrderNotFound codes.Code = 5000
	//OrderSkuEmpty 提交商品为空
	OrderSkuEmpty codes.Code = 5001
	//PayActionInvallid 无效的支付方式
	PayActionInvallid codes.Code = 5002
)

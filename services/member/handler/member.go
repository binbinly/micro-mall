package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/binbinly/pkg/util/validator"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/client"

	"member/logic"
	"member/resource"
	"pkg/constvar"
	"pkg/errno"
	"pkg/handler"
	cpb "pkg/proto/common"
	"pkg/proto/core"
	pb "pkg/proto/member"
)

// Auth 会员服身份验证
func Auth(method string) bool {
	switch method {
	case "Member.Register", "Member.Login", "Member.PhoneLogin", "Member.SendCode":
		//这些路由不需要身份验证
		return false
	}
	return true
}

// Member 会员服务处理器
type Member struct {
	logic         logic.Logic
	thirdService  core.ThirdPartyService
	centerService core.CenterService
}

func New(l logic.Logic, c client.Client) *Member {
	return &Member{
		logic:         l,
		thirdService:  core.NewThirdPartyService(constvar.ServiceThird, c),
		centerService: core.NewCenterService(constvar.ServiceCenter, c),
	}
}

// Register 注册
func (m *Member) Register(ctx context.Context, req *pb.RegisterReq, empty *empty.Empty) error {
	phone, _ := strconv.ParseInt(req.Phone, 10, 64)
	if !validator.RegexMatch(req.Phone, validator.ChineseMobileMatcher) || phone == 0 {
		return errno.ErrMemberPhoneValid
	}
	if req.Password != req.ConfirmPassword {
		return errno.ErrMemberPasswordNotMatch
	}

	// 验证验证码是否正确
	if _, err := m.thirdService.CheckVCode(ctx, &core.VCodeReq{
		Phone: phone,
		Code:  req.Code,
	}); err != nil {
		return err
	}

	// 先判断用户名是否已经注册
	if err := m.logic.MemberExist(ctx, req.Username, phone); err != nil {
		return err
	}

	// 前往用户中心注册
	user, err := m.centerService.Register(ctx, &core.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    phone,
	})
	if err != nil {
		return err
	}

	// 本地注册
	if err = m.logic.MemberRegister(ctx, user.Id, phone, req.Username); err != nil {
		return err
	}

	return nil
}

// Login 登录
func (m *Member) Login(ctx context.Context, req *pb.LoginReq, reply *pb.MemberTokenReply) error {
	// 前往用户中心登录
	user, err := m.centerService.Login(ctx, &core.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	// 本地会员创建
	if err = m.logic.MemberCreate(ctx, user.User); err != nil {
		return err
	}

	reply.Data = resource.MemberTokenResource(user)
	return nil
}

// PhoneLogin 手机号登录
func (m *Member) PhoneLogin(ctx context.Context, req *pb.PhoneLoginReq, reply *pb.MemberTokenReply) error {
	phone, _ := strconv.ParseInt(req.Phone, 10, 64)
	if !validator.RegexMatch(req.Phone, validator.ChineseMobileMatcher) || phone == 0 {
		return errno.ErrMemberPhoneValid
	}

	// 验证验证码是否正确
	if _, err := m.thirdService.CheckVCode(ctx, &core.VCodeReq{
		Phone: phone,
		Code:  req.Code,
	}); err != nil {
		return err
	}

	// 前往用户中心登录
	user, err := m.centerService.PhoneLogin(ctx, &core.PhoneReq{
		Phone: phone,
	})
	if err != nil {
		return err
	}

	// 本地会员创建
	if err = m.logic.MemberCreate(ctx, user.User); err != nil {
		return err
	}

	reply.Data = resource.MemberTokenResource(user)
	return nil
}

// Edit 修改信息
func (m *Member) Edit(ctx context.Context, req *pb.EditReq, empty *empty.Empty) error {
	um := make(map[string]any)
	if req.Avatar != "" {
		um["avatar"] = req.Avatar
	}
	if req.Nickname != "" {
		um["nickname"] = req.Nickname
	}
	if req.Sign != "" {
		um["sign"] = req.Sign
	}
	if len(um) == 0 { //数据为空
		return errno.ErrParamsCheckInvalid
	}

	if err := m.logic.MemberEdit(ctx, handler.GetUserID(ctx), um); err != nil {
		return errno.MemberReplyErr(err)
	}

	// 更新中心服信息
	content, _ := json.Marshal(um)
	if _, err := m.centerService.Edit(ctx, &core.EditReq{
		Content: content,
	}); err != nil {
		return err
	}

	return nil
}

// PwdEdit 修改密码
func (m *Member) PwdEdit(ctx context.Context, req *pb.PwdEditReq, empty *empty.Empty) error {
	if req.Password != req.ConfirmPassword {
		return errno.ErrMemberPasswordNotMatch
	}
	if req.Password == req.OldPassword {
		return nil
	}

	// 密码只保存在用户中心
	if _, err := m.centerService.EditPwd(ctx, &core.EditPwdReq{
		OldPwd: req.OldPassword,
		Pwd:    req.Password,
	}); err != nil {
		return err
	}

	return nil
}

// Profile 用户信息
func (m *Member) Profile(ctx context.Context, empty *empty.Empty, reply *pb.MemberReply) error {
	member, err := m.logic.MemberInfo(ctx, handler.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}

	reply.Data = resource.MemberResource(member)
	return nil
}

// Logout 退出
func (m *Member) Logout(ctx context.Context, empty *empty.Empty, empty2 *empty.Empty) error {
	// 会员中心登出
	if _, err := m.centerService.Logout(ctx, empty); err != nil {
		return err
	}
	return nil
}

// AddressAdd 添加收货地址
func (m *Member) AddressAdd(ctx context.Context, req *pb.AddressAddReq, reply *pb.AddressIDReply) error {
	id, err := m.logic.MemberAddressAdd(ctx, handler.GetUserID(ctx), req.Name, req.Phone, req.Province,
		req.City, req.County, req.Detail, int(req.AreaCode), int8(req.IsDefault))
	if err != nil {
		return errno.MemberReplyErr(err)
	}

	reply.Data = int32(id)
	return nil
}

// AddressEdit 修改收货地址
func (m *Member) AddressEdit(ctx context.Context, req *cpb.Address, empty *empty.Empty) error {
	if err := m.logic.MemberAddressEdit(ctx, int64(req.Id), handler.GetUserID(ctx), req.Name, req.Phone, req.Province,
		req.City, req.County, req.Detail, int(req.AreaCode), int8(req.IsDefault)); err != nil {
		return errno.MemberReplyErr(err)
	}

	return nil
}

// GetAddressList 收货地址列表
func (m *Member) GetAddressList(ctx context.Context, empty *empty.Empty, reply *pb.AddressReply) error {
	list, err := m.logic.MemberAddressList(ctx, handler.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}

	reply.Data = resource.AddressListResource(list)
	return nil
}

// AddressDel 删除收货地址
func (m *Member) AddressDel(ctx context.Context, req *pb.AddressIDReq, empty *empty.Empty) error {
	if err := m.logic.MemberAddressDel(ctx, req.Id, handler.GetUserID(ctx)); err != nil {
		return errno.MemberReplyErr(err)
	}

	return nil
}

// SendCode 发送验证码
func (m *Member) SendCode(ctx context.Context, req *pb.PhoneReq, reply *pb.CodeReply) error {
	// 调用第三方服务发送验证码
	code, err := m.thirdService.SendSMS(ctx, &core.SendSMSReq{Phone: req.Phone})
	if err != nil {
		return err
	}

	reply.Data = code.Code
	return nil
}

// GetAddressInfo 获取收货地址信息
func (m *Member) GetAddressInfo(ctx context.Context, req *pb.AddressIDReq, internal *pb.AddressInfoInternal) error {
	addr, err := m.logic.GetMemberAddressInfo(ctx, req.Id, handler.GetUserID(ctx))
	if err != nil {
		return errno.MemberReplyErr(err)
	}
	internal.Address = resource.AddressResource(addr)
	return nil
}

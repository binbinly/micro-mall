package handler

import (
	"context"
	"pkg/errno"
	"third-party/logic"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "pkg/proto/core"
)

// ThirdParty 第三方服务处理器
type ThirdParty struct {
	logic logic.Logic
}

func New(l logic.Logic) *ThirdParty {
	return &ThirdParty{logic: l}
}

// SendSMS 发送短信验证码
func (t *ThirdParty) SendSMS(ctx context.Context, req *pb.SendSMSReq, reply *pb.CodeReply) error {
	vCode, err := t.logic.SendSMS(ctx, req.Phone)
	if err != nil {
		return errno.ThirdReplyErr(err)
	}
	reply.Code = vCode

	return nil
}

// CheckVCode 短信验证码验证
func (t *ThirdParty) CheckVCode(ctx context.Context, req *pb.VCodeReq, empty *emptypb.Empty) error {
	if err := t.logic.CheckVCode(ctx, req.Phone, req.Code); err != nil {
		return errno.ThirdReplyErr(err)
	}

	return nil
}

// CheckETHPay 以太币支付验证
func (t *ThirdParty) CheckETHPay(ctx context.Context, req *pb.ETHPayReq, empty *emptypb.Empty) error {
	if err := t.logic.CheckPay(ctx, req.Id, req.Address, req.OrderNo); err != nil {
		return errno.ThirdReplyErr(err)
	}

	return nil
}

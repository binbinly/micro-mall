package handler

import (
	"context"

	"github.com/pkg/errors"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"google.golang.org/protobuf/types/known/emptypb"

	"pkg/app"
	"pkg/constvar"
	"pkg/errno"
	"pkg/handler"
	"pkg/message"
	cpb "pkg/proto/common"
	pb "pkg/proto/seckill"
	"seckill/logic"
	"seckill/resource"
)

// Auth 秒杀服身份验证
func Auth(method string) bool {
	switch method {
	case "Seckill.Kill":
		//这些路由需要身份验证
		return true
	}
	return false
}

// Seckill 秒杀处理器
type Seckill struct {
	Srv        logic.Logic
	OrderEvent micro.Event
}

func New(l logic.Logic, c client.Client) *Seckill {
	return &Seckill{
		Srv:        l,
		OrderEvent: micro.NewEvent(constvar.TopicOrderSeckill, c),
	}
}

// Kill 秒杀商品
func (s *Seckill) Kill(ctx context.Context, req *pb.KillReq, reply *pb.KillReply) error {
	memberID := handler.GetUserID(ctx)
	sku, err := s.Srv.SeckillCheck(ctx, memberID, req.SkuId, req.Num, req.Key)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}
	// 生成订单号
	no := app.GenOrderNo()
	// 快速下单，发送mq消息
	if err = s.OrderEvent.Publish(ctx, &message.SeckillMessage{
		MemberID:  memberID,
		SkuID:     req.SkuId,
		AddressID: req.AddressId,
		Num:       req.Num,
		Price:     sku.Price,
		OrderNo:   no,
	}); err != nil {
		return errors.Wrapf(err, "[seckill.Kill] publish order failed")
	}

	reply.Data = no
	return nil
}

// GetSessionAll 获取所有场次
func (s *Seckill) GetSessionAll(ctx context.Context, empty *emptypb.Empty, reply *pb.SessionsReply) error {
	list, err := s.Srv.GetSessionAll(ctx)
	if err != nil {
		return err
	}

	reply.Data = resource.SessionsResource(list)
	return nil
}

// GetSessionSkus 获取场次下的商品
func (s *Seckill) GetSessionSkus(ctx context.Context, req *pb.SessionIdReq, reply *pb.SkusReply) error {
	list, err := s.Srv.GetSessionSkus(ctx, req.Id)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}

	reply.Data = resource.SkusResource(list)
	return nil
}

// GetSkuDetail 获取秒杀商品信息
func (s *Seckill) GetSkuDetail(ctx context.Context, req *cpb.SkuIDReq, reply *pb.SkuReply) error {
	info, err := s.Srv.GetSkuInfo(ctx, req.Id)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}

	reply.Data = resource.SkuResource(info)
	return nil
}

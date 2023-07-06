package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"

	"center/logic"
	"center/resource"
	"encoding/json"
	"pkg/errno"
	"pkg/handler"
	pb "pkg/proto/core"
)

// Center 中心服务
type Center struct {
	logic logic.Logic
}

func New(l logic.Logic) *Center {
	return &Center{logic: l}
}

func (c *Center) Register(ctx context.Context, req *pb.RegisterReq, reply *pb.RegisterReply) error {
	id, err := c.logic.UserRegister(ctx, req.Username, req.Password, req.Phone)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.Id = id
	return nil
}

func (c *Center) Login(ctx context.Context, req *pb.LoginReq, reply *pb.UserReply) error {
	user, token, err := c.logic.UsernameLogin(ctx, req.Username, req.Password)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.User = resource.UserResource(user)
	reply.Token = token
	return nil
}

func (c *Center) PhoneLogin(ctx context.Context, req *pb.PhoneReq, reply *pb.UserReply) error {
	user, token, err := c.logic.UserPhoneLogin(ctx, req.Phone)
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	reply.User = resource.UserResource(user)
	reply.Token = token
	return nil
}

func (c *Center) Edit(ctx context.Context, req *pb.EditReq, empty *empty.Empty) error {
	data := make(map[string]any)
	if err := json.Unmarshal(req.Content, &data); err != nil {
		return errors.Wrapf(err, "[grpc.center] json unmarshal with data:%s", req.Content)
	}

	if err := c.logic.UserEdit(ctx, handler.GetUserID(ctx), data); err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

func (c *Center) EditPwd(ctx context.Context, req *pb.EditPwdReq, empty *empty.Empty) error {
	if err := c.logic.UserEditPwd(ctx, handler.GetUserID(ctx), req.OldPwd, req.Pwd); err != nil {
		return errno.CenterReplyErr(err)
	}

	return nil
}

func (c *Center) Info(ctx context.Context, empty *empty.Empty, info *pb.UserInfo) error {
	user, err := c.logic.UserInfoByID(ctx, handler.GetUserID(ctx))
	if err != nil {
		return errno.CenterReplyErr(err)
	}
	info = resource.UserResource(user)
	return nil
}

func (c *Center) Logout(ctx context.Context, inEmpty *empty.Empty, empty *empty.Empty) error {
	if err := c.logic.UserLogout(ctx, handler.GetUserID(ctx)); err != nil {
		return errno.CenterReplyErr(err)
	}
	return nil
}

package resource

import (
	"member/model"
	pb "pkg/proto/common"
	center "pkg/proto/core"
	"strconv"
)

// MemberTokenResource 转化会员登录信息
func MemberTokenResource(user *center.UserReply) *pb.MemberToken {
	return &pb.MemberToken{
		Member: &pb.Member{
			Id:       int32(user.User.Id),
			Username: user.User.Username,
			Nickname: user.User.Nickname,
			Avatar:   user.User.Avatar,
			Phone:    strconv.FormatInt(user.User.Phone, 10),
		},
		Token: user.Token,
	}
}

// MemberResource 会员信息转换输出
func MemberResource(model *model.MemberModel) *pb.Member {
	return &pb.Member{
		Id:       int32(model.ID),
		Username: model.Username,
		Nickname: model.Nickname,
		Sign:     model.Sign,
		Avatar:   model.Avatar,
		Area:     model.Area,
		Phone:    strconv.FormatInt(model.Phone, 10),
	}
}

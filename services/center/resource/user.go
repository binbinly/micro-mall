package resource

import (
	"center/model"
	pb "pkg/proto/core"
)

// UserResource 组装数据并输出
func UserResource(input *model.UserModel) *pb.UserInfo {
	if input == nil {
		return &pb.UserInfo{}
	}

	return &pb.UserInfo{
		Id:       input.ID,
		Username: input.Username,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Gender:   pb.UserInfo_Gender(input.Gender),
	}
}

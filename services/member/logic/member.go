package logic

import (
	"context"

	"github.com/pkg/errors"

	"member/model"
	"pkg/dbs"
	"pkg/errno"
	center "pkg/proto/core"
)

// MemberRegister 注册用户
func (s *logic) MemberRegister(ctx context.Context, uid, phone int64, username string) error {
	u := &model.MemberModel{
		PriID:    dbs.PriID{ID: uid},
		Username: username,
		Phone:    phone,
		Status:   model.MemberStatusNormal,
	}
	if err := s.repo.MemberCreate(ctx, u); err != nil {
		return errors.Wrapf(err, "[logic.member] create member")
	}

	return nil
}

// MemberExist 检查用户是否存在
func (s *logic) MemberExist(ctx context.Context, username string, phone int64) error {
	exist, err := s.repo.MemberExist(ctx, username, phone)
	if err != nil {
		return errors.Wrapf(err, "[logic.member] member exist")
	}
	if exist {
		return errno.ErrMemberExisted
	}

	return nil
}

// MemberEdit 修改会员信息
func (s *logic) MemberEdit(ctx context.Context, id int64, um map[string]any) error {
	return s.repo.MemberUpdate(ctx, id, um)
}

// MemberInfo 获取用户信息
func (s *logic) MemberInfo(ctx context.Context, id int64) (*model.MemberModel, error) {
	return s.memberInfo(ctx, id)
}

// MemberCreate 会员信息创建
func (s *logic) MemberCreate(ctx context.Context, user *center.UserInfo) error {
	member, err := s.repo.GetMemberByID(ctx, user.Id)
	if err != nil {
		return errors.Wrapf(err, "[logic.member] by uid: %v", user.Id)
	}
	if member == nil || member.ID == 0 { // 会员不存在，注册
		u := &model.MemberModel{
			PriID:    dbs.PriID{ID: user.Id},
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Gender:   int8(user.Gender),
			Phone:    user.Phone,
			Email:    user.Email,
			Status:   model.MemberStatusNormal,
		}
		if err = s.repo.MemberCreate(ctx, u); err != nil {
			return errors.Wrapf(err, "[logic.member] create member")
		}
	}

	return nil
}

// memberInfo 获取会员模型
func (s *logic) memberInfo(ctx context.Context, id int64) (*model.MemberModel, error) {
	member, err := s.repo.GetMemberByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.member] from repo by id: %d", id)
	}
	if member.ID == 0 {
		return nil, errno.ErrMemberNotFound
	}
	if member.Status != model.MemberStatusNormal {
		return nil, errno.ErrMemberFrozen
	}
	return member, nil
}

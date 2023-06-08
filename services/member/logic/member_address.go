package logic

import (
	"context"
	"pkg/errno"
	"pkg/mysql"

	"github.com/pkg/errors"

	"member/model"
)

// MemberAddressAdd 添加收货地址
func (s *logic) MemberAddressAdd(ctx context.Context, MemberID int64, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) (int64, error) {
	addr := &model.MemberAddressModel{
		MID:       mysql.MID{MemberID: MemberID},
		Name:      name,
		Phone:     phone,
		Province:  province,
		City:      city,
		County:    county,
		AreaCode:  areaCode,
		Detail:    detail,
		IsDefault: isDefault,
	}
	return s.repo.MemberAddressCreate(ctx, addr)
}

// MemberAddressEdit 收货地址修改
func (s *logic) MemberAddressEdit(ctx context.Context, id, MemberID int64, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) error {
	address, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[logic.memberAddress] get id: %v", id)
	}
	if address == nil || address.MemberID != MemberID {
		return errno.ErrMemberAddressNotFound
	}
	isEdit := false
	if name != "" && address.Name != name {
		address.Name = name
		isEdit = true
	}
	if phone != "" && address.Phone != phone {
		address.Phone = phone
		isEdit = true
	}
	if province != "" && address.Province != province {
		address.Province = province
		isEdit = true
	}
	if city != "" && address.City != city {
		address.City = city
		isEdit = true
	}
	if county != "" && address.County != county {
		address.County = city
		isEdit = true
	}
	if detail != "" && address.Detail != detail {
		address.Detail = detail
		isEdit = true
	}
	if areaCode != 0 && address.AreaCode != areaCode {
		address.AreaCode = areaCode
		isEdit = true
	}
	if address.IsDefault != isDefault {
		address.IsDefault = isDefault
		isEdit = true
	}
	if isEdit {
		if err = s.repo.MemberAddressSave(ctx, address); err != nil {
			return errors.Wrapf(err, "[logic.memberAddress] update id: %v, uid: %v", id, MemberID)
		}
	}

	return nil
}

// MemberAddressList 收货地址列表
func (s *logic) MemberAddressList(ctx context.Context, MemberID int64) ([]*model.MemberAddressModel, error) {
	return s.repo.GetMemberAddressList(ctx, MemberID)
}

// MemberAddressDel 删除用户收货地址
func (s *logic) MemberAddressDel(ctx context.Context, id, memberID int64) error {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[logic.memberAddress] check id : %v uid: %v", id, memberID)
	}
	if addr == nil || addr.MemberID != memberID {
		return errno.ErrMemberAddressNotFound
	}
	return s.repo.MemberAddressDelete(ctx, addr)
}

// GetMemberAddressInfo 获取收货地址信息
func (s *logic) GetMemberAddressInfo(ctx context.Context, id, memberID int64) (*model.MemberAddressModel, error) {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.memberAddress] check id : %v uid: %v", id, memberID)
	}
	if addr == nil || addr.MemberID != memberID {
		return nil, errno.ErrMemberAddressNotFound
	}
	return addr, nil
}

package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"member/model"
	"pkg/dbs"
)

// IMemberAddress 用户收货地址接口
type IMemberAddress interface {
	MemberAddressCreate(ctx context.Context, addr *model.MemberAddressModel) (int64, error)
	MemberAddressSave(ctx context.Context, addr *model.MemberAddressModel) error
	GetMemberAddressList(ctx context.Context, memberID int64) (list []*model.MemberAddressModel, err error)
	GetMemberAddressByID(ctx context.Context, id int64) (address *model.MemberAddressModel, err error)
	MemberAddressDelete(ctx context.Context, addr *model.MemberAddressModel) error
}

// MemberAddressCreate 创建收货地址
func (r *Repo) MemberAddressCreate(ctx context.Context, addr *model.MemberAddressModel) (int64, error) {
	if addr.IsDefault == 1 {
		if err := r.cancelDefaultAddress(ctx, addr.MemberID); err != nil {
			return 0, err
		}
	}

	if err := r.DB.WithContext(ctx).Create(addr).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.memberAddress] create")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return addr.ID, nil
}

// MemberAddressSave 更新用户收货地址
func (r *Repo) MemberAddressSave(ctx context.Context, addr *model.MemberAddressModel) error {
	if addr.IsDefault == 1 {
		if err := r.cancelDefaultAddress(ctx, addr.MemberID); err != nil {
			return err
		}
	}

	if err := r.DB.WithContext(ctx).Save(addr).Error; err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] save")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return nil
}

// GetMemberAddressList 用户收货地址列表
func (r *Repo) GetMemberAddressList(ctx context.Context, memberID int64) (list []*model.MemberAddressModel, err error) {
	doKey := fmt.Sprintf("member_address_all:%d", memberID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func(data interface{}) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Where("member_id=?", memberID).Order(dbs.DefaultOrder).Find(&list).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.member] query cache")
	}

	return list, nil
}

// GetMemberAddressByID 用户收货地址详情
func (r *Repo) GetMemberAddressByID(ctx context.Context, id int64) (address *model.MemberAddressModel, err error) {
	doKey := fmt.Sprintf("member_address:%d", id)
	if err = r.QueryCache(ctx, doKey, &address, 0, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).First(data, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.member] query cache")
	}

	return address, nil
}

// MemberAddressDelete 删除收货地址
func (r *Repo) MemberAddressDelete(ctx context.Context, addr *model.MemberAddressModel) error {
	if err := r.DB.WithContext(ctx).Delete(addr).Error; err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] delete")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return nil
}

// cancelDefaultAddress 取消默认地址
func (r *Repo) cancelDefaultAddress(ctx context.Context, memberID int64) error {
	if err := r.DB.WithContext(ctx).Model(&model.MemberAddressModel{}).
		Where("member_id=? and is_default=1", memberID).Update("is_default", 0).Error; err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] set other not defualt by uid: %v", memberID)
	}
	return nil
}

// delAddressCache 删除售后地址缓存
func (r *Repo) delAddressCache(ctx context.Context, id, memberID int64) {
	r.DelCache(ctx, buildAddressAllCacheKey(memberID))
	r.DelCache(ctx, buildAddressCacheKey(id))
}

// buildAddressCacheKey 构建收货地址缓存键
func buildAddressCacheKey(id int64) string {
	return fmt.Sprintf("member_address:%d", id)
}

// buildAddressAllCacheKey 当前用户所有收货地址键
func buildAddressAllCacheKey(memberID int64) string {
	return fmt.Sprintf("member_address_all:%d", memberID)
}

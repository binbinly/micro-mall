package logic

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"member/model"
	"member/repository"
	center "pkg/proto/core"
)

var _ Logic = (*logic)(nil)

// Logic 会员服务接口
type Logic interface {
	MemberRegister(ctx context.Context, uid, phone int64, username string) error
	MemberExist(ctx context.Context, username string, phone int64) error
	MemberCreate(ctx context.Context, user *center.UserInfo) error
	MemberEdit(ctx context.Context, id int64, um map[string]any) error
	MemberInfo(ctx context.Context, id int64) (*model.MemberModel, error)
	MemberAddressAdd(ctx context.Context, MemberID int64, name, phone,
		province, city, county, detail string, areaCode int, isDefault int8) (int64, error)
	MemberAddressEdit(ctx context.Context, id, MemberID int64, name, phone,
		province, city, county, detail string, areaCode int, isDefault int8) error
	MemberAddressList(ctx context.Context, MemberID int64) ([]*model.MemberAddressModel, error)
	MemberAddressDel(ctx context.Context, id, memberID int64) error
	GetMemberAddressInfo(ctx context.Context, id, memberID int64) (*model.MemberAddressModel, error)

	Close() error
}

type logic struct {
	repo repository.IRepo
}

// New init logic
func New(db *gorm.DB, rdb *redis.Client) Logic {
	return &logic{
		repo: repository.New(db, cache.NewRedisCache(rdb)),
	}
}

// Close logic
func (s *logic) Close() error {
	return s.repo.Close()
}

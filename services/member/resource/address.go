package resource

import (
	"member/model"
	pb "pkg/proto/common"
)

// AddressListResource 转换会员收货地址列表输出
func AddressListResource(list []*model.MemberAddressModel) []*pb.Address {
	if len(list) == 0 {
		return []*pb.Address{}
	}
	res := make([]*pb.Address, 0, len(list))
	for _, addr := range list {
		res = append(res, AddressResource(addr))
	}
	return res
}

// AddressResource 转换会员收货地址输出
func AddressResource(addr *model.MemberAddressModel) *pb.Address {
	return &pb.Address{
		Id:        int32(addr.ID),
		Name:      addr.Name,
		Phone:     addr.Phone,
		Province:  addr.Province,
		City:      addr.City,
		County:    addr.County,
		AreaCode:  int32(addr.AreaCode),
		Detail:    addr.Detail,
		IsDefault: int32(addr.IsDefault),
	}
}

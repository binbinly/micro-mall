package resource

import (
	"time"

	"github.com/binbinly/pkg/util"

	pb "pkg/proto/common"
	"seckill/config"
	"seckill/repository"
)

// SkusResource 转换秒杀商品列表
func SkusResource(list []*repository.SkuModel) []*pb.KillSku {
	if len(list) == 0 {
		return []*pb.KillSku{}
	}

	res := make([]*pb.KillSku, 0, len(list))
	for _, sku := range list {
		res = append(res, SkuResource(sku))
	}
	return res
}

// SkuResource 转换sku信息
func SkuResource(sku *repository.SkuModel) *pb.KillSku {
	now := time.Now().Unix()
	var open bool
	var key string
	if sku.StartAt <= now && sku.EndAt >= now {
		open = true
		key = sku.Key //只有秒杀开始才发送key
	}
	return &pb.KillSku{
		Id:            int32(sku.ID),
		Price:         util.ParseAmount(sku.Price),
		Count:         int32(sku.Count),
		Limit:         int32(sku.Limit),
		OriginalPrice: util.ParseAmount(sku.OriginalPrice),
		Title:         sku.Title,
		Cover:         util.FormatResUrl(config.Cfg.DFS.Endpoint, sku.Cover),
		Key:           key,
		Open:          open,
		StartAt:       int32(sku.StartAt),
	}
}

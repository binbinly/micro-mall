package handler

import (
	"context"
	"pkg/errno"
	"warehouse/logic"

	"google.golang.org/protobuf/types/known/emptypb"
	cpb "pkg/proto/common"
	pb "pkg/proto/warehouse"
)

// Warehouse 仓储服务处理器
type Warehouse struct {
	logic logic.Logic
}

func New(l logic.Logic) *Warehouse {
	return &Warehouse{logic: l}
}

// GetSkuStock 获取sku库存数量
func (w *Warehouse) GetSkuStock(ctx context.Context, req *cpb.SkuIDReq, reply *pb.SkuStockReply) error {
	num, err := w.logic.GetSkuStock(ctx, req.Id)
	if err != nil {
		return err
	}

	reply.Num = int32(num)
	return nil
}

// BatchSkusStock 批量获取sku库存数量
func (w *Warehouse) BatchSkusStock(ctx context.Context, req *pb.SkusReq, reply *pb.SkusStockReply) error {
	nums, err := w.logic.GetSkusStock(ctx, req.Ids)
	if err != nil {
		return err
	}

	reply.SkuNums = nums
	return nil
}

// SKuStockLock 锁定库存
func (w *Warehouse) SKuStockLock(ctx context.Context, req *pb.SkuStockLockReq, empty *emptypb.Empty) error {
	if err := w.logic.SKuStockLock(ctx, req.OrderId, req.OrderNo, req.Consignee, req.Phone, req.Address, req.Note, req.SkuNums); err != nil {
		return errno.WarehouseReplyErr(err)
	}

	return nil
}

// SkuStockUnlock 解锁库存
func (w *Warehouse) SkuStockUnlock(ctx context.Context, req *pb.SkuStockUnlockReq, empty *emptypb.Empty) error {
	if err := w.logic.SkuStockUnlock(ctx, req.OrderId, req.Finish); err != nil {
		return err
	}

	return nil
}

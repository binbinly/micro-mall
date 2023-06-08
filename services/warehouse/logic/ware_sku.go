package logic

import (
	"context"
	"github.com/binbinly/pkg/repo"
	"github.com/pkg/errors"
	"pkg/errno"
	"pkg/mysql"

	"warehouse/model"
)

// GetSkuStock 获取sku库存数量
func (l *logic) GetSkuStock(ctx context.Context, skuID int64) (int, error) {
	stock, err := l.repo.GetWareSkuStock(ctx, skuID)
	if err != nil {
		return 0, errors.Wrapf(err, "[logic.wareSku] stock sku_id: %v", skuID)
	}

	return stock.Stock - stock.StockLock, nil
}

// GetSkusStock 批量获取sku的库存数量
func (l *logic) GetSkusStock(ctx context.Context, skuIds []int64) (map[int64]int32, error) {
	stocks, err := l.repo.BatchGetWareSkusStock(ctx, skuIds)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]int32, len(stocks))
	for _, stock := range stocks {
		res[stock.SkuID] = int32(stock.Stock - stock.StockLock)
	}

	return res, nil
}

// SKuStockLock sku库存锁定
func (l *logic) SKuStockLock(ctx context.Context, orderID int64, orderNo, consignee, phone, address, note string,
	sku map[int64]int32) error {

	skuIds := make([]int64, 0, len(sku))
	for id := range sku {
		skuIds = append(skuIds, id)
	}

	skus, err := l.repo.BatchGetWareSkus(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[logic.wareSku] batch skus by ids: %v", skuIds)
	}
	if len(skus) == 0 {
		return errno.ErrWareInventoryShortage
	}
	//库存工作单
	task := &model.WareTaskModel{
		OrderID:   orderID,
		OrderNo:   orderNo,
		Consignee: consignee,
		Phone:     phone,
		Address:   address,
		Note:      note,
		Status:    model.TaskStatusLock,
	}

	// 开启事务
	tx := l.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//保存库存工作单
	if err = l.repo.CreateWareTask(ctx, tx, task); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.wareSku] create ware task")
	}

	var taskDetails []*model.WareTaskDetailModel
	for _, m := range skus {
		if m.Stock-m.StockLock < int(sku[m.SkuID]) { //库存不足
			tx.Rollback()
			return errno.ErrWareInventoryShortage
		}
		m.StockLock += int(sku[m.SkuID])

		if err = l.repo.WareSkuSave(ctx, tx, m); err != nil {
			tx.Rollback()
			return errors.Wrapf(err, "[logic.wareSku] save")
		}
		taskDetails = append(taskDetails, &model.WareTaskDetailModel{
			TaskID:  task.ID,
			WareID:  m.WareID,
			Sku:     mysql.Sku{SkuID: m.SkuID},
			SkuName: m.SkuName,
			SkuNum:  int(sku[m.SkuID]),
		})
	}

	//批量保存库存工作单详情
	if err = l.repo.BatchCreateWareTaskDetail(ctx, tx, taskDetails); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.wareSku] batch create task detail")
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.wareSku] tx commit lock stock")
	}

	return nil
}

// SkuStockUnlock sku库存解锁 finish 订单是否完成 订单完成直接扣减库存
// 1，下订单成功后，订单自动过期或者用户手动取消，需要解锁库存
// 2，下订单成功，库存锁定成功，接下来业务调用失败，导致订单回滚，之前锁定成功的库存需要解锁
func (l *logic) SkuStockUnlock(ctx context.Context, orderID int64, finish bool) error {
	task, err := l.repo.GetTaskByOrderID(ctx, orderID)
	if err != nil {
		return errors.Wrapf(err, "[logic.wareSku] task by order_id:%v", orderID)
	}
	//库存工作单不存在或者工作单状态非解锁，无需解锁跳过
	if task.ID == 0 || task.Status != model.TaskStatusLock {
		return nil
	}

	//库存工作单详情
	details, err := l.repo.GetTaskDetailByID(ctx, task.ID)
	if err != nil {
		return errors.Wrapf(err, "[logic.wareSku] task detail by task_id: %v", task.ID)
	}
	skuIds := make([]int64, 0, len(details))
	skuStock := make(map[int64]int, len(details))
	for _, detail := range details {
		skuIds = append(skuIds, detail.SkuID)
		skuStock[detail.SkuID] = detail.SkuNum
	}
	skus, err := l.repo.BatchGetWareSkus(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[logic.wareSku] batch skus by ids: %v", skuIds)
	}

	// 开启事务
	tx := l.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//更新库存工作单状态
	status := model.TaskStatusUnlock
	if finish {
		status = model.TaskStatusFinish
	}
	if err = l.repo.UpdateWareTaskStatus(ctx, tx, orderID, status); err != nil {
		tx.Rollback()
		if errors.Is(err, repo.ErrRecordNotModified) {
			return nil
		}
		return errors.Wrapf(err, "[logic.wareSku] update task")
	}

	for _, m := range skus {
		if m.StockLock < skuStock[m.SkuID] { //库存不足
			continue
		}
		m.StockLock -= skuStock[m.SkuID]
		if finish { //订单完成，库存扣减
			m.Stock -= skuStock[m.SkuID]
		}
		if err = l.repo.WareSkuSave(ctx, tx, m); err != nil {
			tx.Rollback()
			return errors.Wrapf(err, "[logic.wareSku] save")
		}
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[logic.wareSku] tx commit unlock stock")
	}

	return nil
}

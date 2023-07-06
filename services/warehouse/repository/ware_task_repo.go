package repository

import (
	"context"
	"pkg/dbs"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"warehouse/model"
)

// GetTaskByOrderID 获取库存工作单
func (r *Repo) GetTaskByOrderID(ctx context.Context, orderID int64) (task *model.WareTaskModel, err error) {
	task = new(model.WareTaskModel)
	err = r.DB.WithContext(ctx).Where("order_id=?", orderID).First(task).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.wareTask] by order_id: %v", orderID)
	}

	return task, nil
}

// UpdateWareTaskStatus 更新库存工作单状态
func (r *Repo) UpdateWareTaskStatus(ctx context.Context, tx *gorm.DB, orderID int64, status int) error {
	result := tx.WithContext(ctx).Model(&model.WareTaskModel{}).
		Where("order_id=? and status=1", orderID).Update("status", status)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "[repo.wareTask] update")
	}
	if result.RowsAffected == 0 { //没有记录更新
		return dbs.ErrRecordNotModified
	}

	return nil
}

// CreateWareTask 创建库存工作单
func (r *Repo) CreateWareTask(ctx context.Context, tx *gorm.DB, task *model.WareTaskModel) error {
	if err := tx.WithContext(ctx).Create(task).Error; err != nil {
		return errors.Wrapf(err, "[repo.wareTask] create")
	}

	return nil
}

// BatchCreateWareTaskDetail 批量创建库存工作单详情
func (r *Repo) BatchCreateWareTaskDetail(ctx context.Context, tx *gorm.DB, items []*model.WareTaskDetailModel) error {
	err := tx.WithContext(ctx).Model(&model.WareTaskDetailModel{}).Create(&items).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.wareTask] batch detail create")
	}
	return nil
}

// GetTaskDetailByID 获取库存工作单详情
func (r *Repo) GetTaskDetailByID(ctx context.Context, taskID int64) (list []*model.WareTaskDetailModel, err error) {
	err = r.DB.WithContext(ctx).Model(&model.WareTaskDetailModel{}).Where("task_id=?", taskID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.wareTask] detail list")
	}

	return
}

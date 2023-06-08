package repository

import (
	"context"
	"pkg/mysql"

	"github.com/pkg/errors"

	"market/model"
)

// GetNoticeList 公告列表
func (r *Repo) GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error) {
	err = r.DB.WithContext(ctx).Scopes(mysql.OffsetPage(offset, limit)).Order(mysql.DefaultOrder).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.notice] db find")
	}
	return
}

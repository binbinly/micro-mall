package repository

import (
	"context"

	"github.com/pkg/errors"

	"market/model"
	"pkg/dbs"
)

// GetNoticeList 公告列表
func (r *Repo) GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error) {
	err = r.DB.WithContext(ctx).Scopes(dbs.OffsetPage(offset, limit)).Order(dbs.DefaultOrder).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.notice] db find")
	}
	return
}

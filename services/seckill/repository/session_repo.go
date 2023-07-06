package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
)

const (
	sessionKey = "seckill:sessions"
)

// GetSessionAll 获取当下所有秒杀场次
func (r *Repo) GetSessionAll(ctx context.Context) ([]*SessionModel, error) {
	data, err := r.redis.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.session] getall")
	}

	now := time.Now().Unix()
	var sessions []*SessionModel
	for _, datum := range data {
		if datum == "" {
			continue
		}
		session := &SessionModel{}
		if err = json.Unmarshal([]byte(datum), session); err != nil {
			logger.Warnf("[repo.session] json.unmarshal err: %v", err)
			continue
		}
		if now > session.EndAt { //当前场次已结束
			continue
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

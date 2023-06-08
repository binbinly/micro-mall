package resource

import (
	"seckill/repository"
	"time"

	pb "pkg/proto/common"
)

// SessionsResource 转换场次结构输出
func SessionsResource(list []*repository.SessionModel) (res []*pb.Session) {
	if len(list) == 0 {
		return []*pb.Session{}
	}

	now := time.Now().Unix()
	for _, session := range list {
		var open bool
		if session.StartAt <= now && session.EndAt >= now {
			open = true
		}
		s := &pb.Session{
			Id:   int32(session.ID),
			Name: session.Name,
			Open: open,
			Skus: make([]*pb.KillSku, 0, len(session.Skus)),
		}
		if len(session.Skus) > 0 {
			s.Skus = SkusResource(session.Skus)
		}
		res = append(res, s)
	}
	return
}

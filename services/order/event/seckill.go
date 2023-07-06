package event

import (
	"context"

	"github.com/binbinly/pkg/logger"

	"order/logic"
	"pkg/message"
)

// KillEvent 秒杀订单消息 event
type KillEvent struct {
	logic logic.Logic
}

// NewKill 实例化
func NewKill(l logic.Logic) *KillEvent {
	return &KillEvent{logic: l}
}

// Handler 秒杀订单消息处理
func (e *KillEvent) Handler(ctx context.Context, message *message.SeckillMessage) error {
	logger.Debugf("[order.killEvent] handler message: %v", message)
	if err := e.logic.SubmitKillOrder(ctx, message.MemberID, message.SkuID, message.AddressID,
		message.Price, int(message.Num), message.OrderNo); err != nil {
		logger.Warnf("[order.killEvent] submit: %v err: %v", message, err)
		return err
	}

	return nil
}

package event

import (
	"context"
	"order/logic"
	"pkg/message"

	"go-micro.dev/v4/logger"
)

// CancelEvent 订单取消 event
type CancelEvent struct {
	logic logic.Logic
}

// NewCancel 实例化
func NewCancel(l logic.Logic) *CancelEvent {
	return &CancelEvent{logic: l}
}

// Handler 秒杀订单消息处理
func (e *CancelEvent) Handler(ctx context.Context, message *message.OrderCancelMessage) error {
	logger.Infof("[order.CancelEvent] handler message: %v", message)

	if err := e.logic.OrderCancel(context.Background(), message.MemberID, message.OrderID); err != nil {
		logger.Warnf("[order.CancelEvent] order cancel err: %v", err)
	}
	return nil
}

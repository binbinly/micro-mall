package event

import (
	"context"

	"go-micro.dev/v4/logger"

	"pkg/message"
	"warehouse/logic"
)

type Event struct {
	logic logic.Logic
}

func New(l logic.Logic) *Event {
	return &Event{logic: l}
}

// Handler 消息处理
func (e *Event) Handler(ctx context.Context, message *message.OrderMessage) error {
	logger.Debugf("[warehouse.event] handler message: %v", message)
	if err := e.logic.SkuStockUnlock(ctx, message.OrderId, message.Finish); err != nil {
		logger.Warnf("[warehouse.event] handler message: %v", message)
		return err
	}
	return nil
}

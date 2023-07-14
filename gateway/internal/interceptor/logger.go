package interceptor

import (
	"context"

	"github.com/binbinly/pkg/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

// Logger adapts go-kit logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func Logger() logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			logger.Debug(largs...)
		case logging.LevelInfo:
			logger.Info(largs...)
		case logging.LevelWarn:
			logger.Warn(largs...)
		case logging.LevelError:
			logger.Error(largs...)
		default:
			logger.Fatal(largs...)
		}
	})
}

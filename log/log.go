package log

import (
	"go.uber.org/zap"
)

// Logger ...
func Logger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar
}

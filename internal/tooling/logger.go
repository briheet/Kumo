package tooling

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewDevelopmentLogger() *zap.Logger {

	logger, _ := zap.NewDevelopment()
	return logger
}

// TODO: Not in use currently
func NewProductionLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()

	// Customs
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

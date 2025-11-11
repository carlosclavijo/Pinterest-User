package services

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger(env string) {
	once.Do(func() {
		var err error
		if env == "production" {
			logger, err = zap.NewProduction()
		} else {
			logger, err = zap.NewDevelopment()
		}
		if err != nil {
			panic(err)
		}
	})
}

func Logger() *zap.Logger {
	if logger == nil {
		panic("logger not initialized - Call logging.InitLogger() first")
	}
	return logger
}

type ZapAdapter struct {
	logger *zap.Logger
}

func NewZapAdapter() *ZapAdapter {
	return &ZapAdapter{
		logger: Logger(),
	}
}

func (z *ZapAdapter) Debug(msg string, args ...any) {
	z.logger.Sugar().Debugf(msg, args...)
}

func (z *ZapAdapter) Info(msg string, args ...any) {
	z.logger.Sugar().Infof(msg, args...)
}

func (z *ZapAdapter) Warn(msg string, args ...any) {
	z.logger.Sugar().Warnf(msg, args...)
}

func (z *ZapAdapter) Error(msg string, args ...any) {
	z.logger.Sugar().Errorf(msg, args...)
}

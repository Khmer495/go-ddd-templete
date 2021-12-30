package logger

import (
	"github.com/Khmer495/go-templete/internal/pkg/config"
	"github.com/Khmer495/go-templete/internal/pkg/util"
	"go.uber.org/zap"
)

func NewLogger() {

	var logger *zap.Logger
	env := config.Env()
	if util.IsPrd(env) {
		logger = newProdLogger()
	} else {
		logger = newDevLogger()
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Complete setting up zap logger.")
}

func newProdLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Sugar().Panicf("zap.NewProduction: %+v", err)
	}
	return logger
}

func newDevLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Sugar().Panicf("zap.NewDevelopment: %+v", err)
	}
	return logger
}

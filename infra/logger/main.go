package logger

import (
	"github.com/gabrielmoura/estudo-api-go/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {

	config := zap.NewProductionConfig()
	if configs.Conf.AppEnv == "development" {
		config = zap.NewDevelopmentConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	Logger, _ = config.Build()
	defer Logger.Sync()
}

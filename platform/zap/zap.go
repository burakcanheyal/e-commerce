package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000")
	config.EncoderConfig.StacktraceKey = ""

	log, err := config.Build()
	if err != nil {
		log.Fatal(err.Error())
	}
	Logger = log.Sugar()
}

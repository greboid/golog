package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustCreateLogger(debug bool) *zap.SugaredLogger {
	logger, err := CreateLogger(debug)
	if err != nil {
		panic("Unable to create logger")
	}
	return logger
}

func CreateLogger(debug bool) (*zap.SugaredLogger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableCaller = !debug
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.DisableStacktrace = !debug
	zapConfig.OutputPaths = []string{"stdout"}
	if debug {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
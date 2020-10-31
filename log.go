package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Creates a new SugaredLogger and panics if it couldn't be created
func MustCreateLogger(debug bool) *zap.SugaredLogger {
	logger, err := CreateLogger(debug)
	if err != nil {
		panic("Unable to create logger")
	}
	return logger
}

//Creates a new SugaredLogger, an error is returned if the creation fails
//
//The debug parameters controls how much logging is displayed.  It also controls if the caller and stack trace are shown
//on error
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
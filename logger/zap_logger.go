package logger

import (
	"errors"
	"go.uber.org/zap"
)

func NewZapLoggerForEnv(env string, callerSkip int) (*zap.SugaredLogger, error) {
	if env == "" || env == "local" || env == "test" || env == "qa" || env == "dev" {
		logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
		return logger.Sugar(), err
	} else if env == "prod" {
		logger, err := zap.NewProduction(zap.AddCallerSkip(callerSkip), zap.AddStacktrace(zap.ErrorLevel))
		return logger.Sugar(), err
	}
	return nil, errors.New("not valid")
}

package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.SugaredLogger

func Init(levelStr string) {
	level := parseLevel(levelStr)
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	logger, _ := cfg.Build()
	global = logger.Sugar()
}

func parseLevel(levelStr string) zapcore.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func Get() *zap.SugaredLogger {
	return global
}

func Infof(template string, args ...interface{}) {
	global.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	global.Errorf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	global.Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	global.Warnf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	global.Fatalf(template, args...)
}

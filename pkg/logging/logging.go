package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"home-hap/internal/config"
	"strings"
)

var zl *zap.Logger

type Logger struct {
	*zap.Logger
}

func GetLogger() *Logger {
	return &Logger{
		zl,
	}
}

func GetLoggerWithField(k string, v string) *Logger {
	return &Logger{
		zl.With(zap.String(k, v)),
	}
}
func GetLoggerWithFields(fields ...zap.Field) *Logger {
	return &Logger{
		zl.With(fields...),
	}
}

func Init(cfg config.LogOpts) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapLogLevel(cfg.Level)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zl, _ = config.Build()
}

//func init() {
//	// TODO add configuration
//	zl, _ = zap.NewProduction()
//}

func zapLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	}

	return zap.InfoLevel
}

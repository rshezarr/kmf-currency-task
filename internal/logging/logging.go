package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoderConf = zapcore.EncoderConfig{
	TimeKey:        "timestamp",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func InitConfiguredLogger(externalConfig zap.Config) error {
	logger, err := externalConfig.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
	return nil
}

func InitDefaultLogger() error {
	var defaultJsonConfig = zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:   true,
		DisableCaller: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConf,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	zap.AddStacktrace(zap.ErrorLevel)
	return InitConfiguredLogger(defaultJsonConfig)
}

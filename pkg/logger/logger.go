package logger

import (
	"log"
	"os"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// note: this can be made a library and used in multiple codebases

func InitLogger() *zap.Logger {
	logLevelStr := "debug" // default set to debug log, will show all logs
	encCfg := zap.NewProductionEncoderConfig()
	if os.Getenv("LOG_LEVEL") != constants.EMPTY_STRING {
		logLevelStr = os.Getenv("LOG_LEVEL")
	}
	logLevel := getLogLevel(logLevelStr)
	core := zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), zapcore.AddSync(os.Stdout), logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	return logger
}

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
)

func getLogLevel(logLevel string) zapcore.Level {
	var level zapcore.Level
	switch logLevel {
	case DEBUG:
		level = zapcore.DebugLevel
	case INFO:
		level = zapcore.InfoLevel
	case WARN:
		level = zapcore.WarnLevel
	case ERROR:
		level = zapcore.ErrorLevel
	case FATAL:
		level = zapcore.FatalLevel
	default:
		log.Panic("invalid loglevel provided!")
	}
	return level
}

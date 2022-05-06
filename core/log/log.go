package log

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func Setup(path, servername string, debug bool) {
	encoder := getEncoder()

	var cores []zapcore.Core

	writeSyncer := getLogWriter(path, servername)
	fileCore := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	cores = append(cores, fileCore)

	if debug {
		consoleDebug := zapcore.Lock(os.Stdout)
		consoleCore := zapcore.NewCore(encoder, consoleDebug, zapcore.InfoLevel)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)
	logger = zap.
		New(core, zap.AddCaller(), zap.AddCallerSkip(1)).
		Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		TimeKey:      "time",
		LevelKey:     "level",
		NameKey:      "logger",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

//todo 调整日志切割方案
func getLogWriter(runtimePath, servername string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   runtimePath + "/" + servername + ".log",
		MaxSize:    2,
		MaxBackups: 10000,
		MaxAge:     0,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func format(label string, log interface{}) string {
	e := fmt.Sprintf("[%s] %s", label, log)
	return e
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

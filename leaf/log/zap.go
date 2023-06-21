package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strings"
)

type ZapLogger struct {
	level int
	log   *zap.SugaredLogger
}

func New(strLevel string, pathname string, flag int) (*ZapLogger, error) {
	logLevel := zap.NewAtomicLevel()
	// level
	var level int
	switch strings.ToLower(strLevel) {
	case "info":
		logLevel.SetLevel(zap.InfoLevel)
	case "warning":
		logLevel.SetLevel(zap.WarnLevel)
	case "debug":
		logLevel.SetLevel(zap.DebugLevel)
	case "error":
		logLevel.SetLevel(zap.ErrorLevel)
	case "fatal":
		logLevel.SetLevel(zap.FatalLevel)
	default:
		return nil, errors.New("unknown level: " + strLevel)
	}

	// logger
	logEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	logOutput := zapcore.Lock(os.Stdout)

	// var fileOutput *os.File
	var loglumFile *lumberjack.Logger
	if pathname != "" {
		// fileOutput, _ = os.Create()
		loglumFile = &lumberjack.Logger{
			Filename:   pathname + "/logfile.log", // 日志文件路径
			MaxSize:    1,                         // 单个日志文件的最大尺寸（单位：MB）
			MaxBackups: 3,                         // 保留的旧日志文件的最大个数
			MaxAge:     30,                        // 旧日志文件保留的最长时间（单位：天）
			Compress:   false,                     // 是否压缩旧日志文件
		}
	}

	// 创建核心日志对象
	consoleCore := zapcore.NewCore(logEncoder, logOutput, logLevel)

	var fileCore zapcore.Core
	if loglumFile != nil {
		fileCore = zapcore.NewCore(logEncoder, zapcore.AddSync(loglumFile), logLevel)
	}

	loggerZap := zap.New(zapcore.NewTee(consoleCore, fileCore))

	logger := new(ZapLogger)

	logger.level = level
	logger.log = loggerZap.Sugar()

	return logger, nil
}

// It's dangerous to call the method on logging
func (logger *ZapLogger) Close() {
	logger.log.Sync()
}

func (logger *ZapLogger) Debug(format string, a ...interface{}) {
	//logger.doPrintf(debugLevel, printDebugLevel, format, a...)
	logger.log.Debug(format, a)
}

func (logger *ZapLogger) Info(format string, a ...interface{}) {
	logger.log.Infof(format, a)
}

func (logger *ZapLogger) Warning(format string, a ...interface{}) {
	logger.log.Warnf(format, a)
}

func (logger *ZapLogger) Error(format string, a ...interface{}) {
	logger.log.Errorf(format, a)
}

func (logger *ZapLogger) Fatal(format string, a ...interface{}) {
	logger.log.Fatalf(format, a)
}

var gLogger, _ = New("debug", "", log.LstdFlags)

// It's dangerous to call the method on logging
func Export(logger *ZapLogger) {
	if logger != nil {
		gLogger = logger
	}
}

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a)
}

func Info(format string, a ...interface{}) {
	gLogger.Info(format, a)
}

func Warning(format string, a ...interface{}) {
	gLogger.Warning(format, a)
}

func Release(format string, a ...interface{}) {
	gLogger.Info(format, a)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a)
}

func Close() {
	gLogger.Close()
}

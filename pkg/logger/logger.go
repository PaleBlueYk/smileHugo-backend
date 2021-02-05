package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// logger
var Logger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 默认执行
func init() {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()
}

func Init() error {
	logLevel := getLoggerLevel(viper.GetString("logging.logLevel"))
	runMode := viper.GetString("runMode")

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	var core zapcore.Core

	var logger *zap.Logger
	consoleLog := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(os.Stdout), logLevel)
	if runMode == "release" {
		application := viper.GetString("application")
		filePath := viper.GetString("log.filePath")

		fileName := fmt.Sprintf("%s/%s.log", filePath, application)
		syncWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    100, // MB
			MaxBackups: 7,   // number of backups
			MaxAge:     28,  //day
			LocalTime:  true,
			Compress:   true,
		})
		jsonLog := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(logLevel))
		core = zapcore.NewTee(jsonLog, consoleLog)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		Logger = logger.Sugar()
		Infof("using log file: %s", fileName)
	} else {
		core = zapcore.NewTee(consoleLog)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		Logger = logger.Sugar()
		Info("using console logger")
	}
	return nil
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	Logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	Logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
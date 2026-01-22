package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Zero0719/go-api/pkg/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 设置日志等级
	logLevel := new(zapcore.Level)

	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	Logger = zap.New(
		core, 
		zap.AddCaller(), 
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	zap.ReplaceGlobals(Logger)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 设置日志存储格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
        FunctionKey:    zapcore.OmitKey,
        MessageKey:     "message",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
        EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
        EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
        EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
        EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	if app.IsLocal() {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		MaxSize: maxSize,
		MaxBackups: maxBackup,
		MaxAge: maxAge,
		Compress: compress,
	}

	if app.IsLocal() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.AddSync(lumberJackLogger)
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("json marshal error", err.Error()))
	}
	return string(b)
}

func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

func DebugString(moduleName string, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName string, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName string, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName string, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName string, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

func DebugJSON(moduleName string, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName string, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName string, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName string, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName string, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}
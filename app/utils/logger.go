package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger zerolog.Logger

// 定义级别写入器，实现zerolog.LevelWriter接口
type LevelLogger struct {
	debug *DynamicLogWriter
	info  *DynamicLogWriter
	warn  *DynamicLogWriter
	err   *DynamicLogWriter
}

// 动态日志写入器，支持按日期切换文件
type DynamicLogWriter struct {
	level    string
	baseDir  string
	writer   *lumberjack.Logger
	lastDate string
}

// WriteLevel 按级别写入不同文件
func (l *LevelLogger) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	switch level {
	case zerolog.DebugLevel:
		return l.debug.Write(p)
	case zerolog.InfoLevel:
		return l.info.Write(p)
	case zerolog.WarnLevel:
		return l.warn.Write(p)
	case zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel:
		return l.err.Write(p)
	default:
		return os.Stdout.Write(p) // 其他级别输出到控制台
	}
}

// Write 实现io.Writer接口
func (d *DynamicLogWriter) Write(p []byte) (int, error) {
	d.checkAndSwitchFile()
	return d.writer.Write(p)
}

// checkAndSwitchFile 检查日期是否变化，如果变化则切换文件
func (d *DynamicLogWriter) checkAndSwitchFile() {
	currentDate := time.Now().Format("2006-01-02")
	if d.lastDate != currentDate {
		d.lastDate = currentDate
		d.createNewWriter()
	}
}

// createNewWriter 创建新的日志写入器
func (d *DynamicLogWriter) createNewWriter() {
	// 目录格式: runtime/logs/年/月
	yearDir := filepath.Join(d.baseDir, time.Now().Format("2006"))
	monthDir := filepath.Join(yearDir, time.Now().Format("01"))
	if err := os.MkdirAll(monthDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 文件名格式: 日期-级别.log
	logFile := filepath.Join(monthDir, fmt.Sprintf("%s-%s.log", d.lastDate, d.level))

	d.writer = &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    50,   // 单个文件最大50MB
		MaxBackups: 10,   // 保留10个备份文件
		MaxAge:     30,   // 保留30天
		Compress:   true, // 压缩备份文件
		LocalTime:  true, // 使用本地时间
	}
}

// Write 实现io.Writer接口（默认写入）
func (l *LevelLogger) Write(p []byte) (int, error) {
	return l.WriteLevel(zerolog.NoLevel, p)
}

func InitLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// 创建根目录
	rootDir := "runtime/logs"
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 初始化各级别动态日志写入器
	levelLogger := &LevelLogger{
		debug: newDynamicLogWriter("debug", rootDir),
		info:  newDynamicLogWriter("info", rootDir),
		warn:  newDynamicLogWriter("warn", rootDir),
		err:   newDynamicLogWriter("error", rootDir),
	}

	// 组合控制台和文件输出
	multiWriter := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stdout}, // 控制台输出
		levelLogger,                           // 文件按级别输出
	)

	Logger = zerolog.New(multiWriter).With().Timestamp().Logger()

	return Logger
}

// newDynamicLogWriter 创建动态日志写入器
func newDynamicLogWriter(level, baseDir string) *DynamicLogWriter {
	writer := &DynamicLogWriter{
		level:    level,
		baseDir:  baseDir,
		lastDate: time.Now().Format("2006-01-02"),
	}
	writer.createNewWriter()
	return writer
}

package log

import (
	"adolesce/internal/conf"
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
	"sync"
)

type Logger interface {
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	// uses a message with some additional context.pairs of keys and values
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	//  uses fmt.Sprintf to log a templated message
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	// Flush calls the underlying Core's Sync method, flushing any buffered
	// log entries. Applications should take care to call Sync before exiting.
	Flush()
}

type zapLog struct {
	zapLogger *zap.Logger
}

var (
	std = initLogger(DefaultOptions())
	mu  sync.Mutex
)

func NewLogger(config *conf.Configs) (Logger, error) {
	mu.Lock()
	defer mu.Unlock()
	// copier
	opts := DefaultOptions()
	if err := copier.Copy(opts, config.Log); err != nil {
		return nil, err
	}
	if err := validate(opts); err != nil {
		return nil, err
	}
	std = initLogger(opts)
	return std, nil
}

func validate(opts *Options) error {
	if opts == nil {
		return fmt.Errorf("log config must not nil")
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		return err
	}

	format := strings.ToLower(opts.Format)
	if format != consoleFormat && format != jsonFormat {
		return fmt.Errorf("log format not valid: %q", format)
	}
	return nil
}

func initLogger(opts *Options) *zapLog {
	core := opts.initCore()
	logger := zap.New(core,
		zap.AddStacktrace(zapcore.PanicLevel), // 配置超过panic级别，不输出栈堆信息
		zap.AddCallerSkip(1))
	zap.RedirectStdLog(logger)
	return &zapLog{logger.Named(opts.Env)}
}

func (lg *zapLog) Flush() {
	_ = lg.zapLogger.Sync()
}

// StdErrLogger 转成标准库logger,级别是error
func StdErrLogger() *log.Logger {
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.ErrorLevel); err == nil {
		std.Error("StdErrLogger-->NewStdLogAt", zap.Error(err))
		return l
	}
	return nil
}

// StdInfoLogger 转成标准库logger,级别是info
func StdInfoLogger() *log.Logger {
	l, err := zap.NewStdLogAt(std.zapLogger, zapcore.InfoLevel)
	if err != nil {
		std.Error("StdInfoLogger-->NewStdLogAt", zap.Error(err))
		return nil
	}
	return l
}

// implement logger
func (z *zapLog) Info(msg string, fields ...Field) {
	z.zapLogger.Info(msg, fields...)
}

func (z *zapLog) Debug(msg string, fields ...Field) {
	z.zapLogger.Debug(msg, fields...)
}

func (z *zapLog) Warn(msg string, fields ...Field) {
	z.zapLogger.Warn(msg, fields...)
}

func (z *zapLog) Error(msg string, fields ...Field) {
	z.zapLogger.Error(msg, fields...)
}

func (z *zapLog) Panic(msg string, fields ...Field) {
	z.zapLogger.Panic(msg, fields...)
}

func (z *zapLog) Fatal(msg string, fields ...Field) {
	z.zapLogger.Fatal(msg, fields...)
}

func (z *zapLog) Debugw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Debugw(msg, keysAndValues)
}

func (z *zapLog) Infow(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Infow(msg, keysAndValues)
}

func (z *zapLog) Warnw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Warnw(msg, keysAndValues)
}

func (z *zapLog) Errorw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Errorw(msg, keysAndValues)
}

func (z *zapLog) DPanicw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().DPanicw(msg, keysAndValues)
}

func (z *zapLog) Panicw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Panicw(msg, keysAndValues)
}

func (z *zapLog) Fatalw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Fatalw(msg, keysAndValues)
}

func (z *zapLog) Debugf(template string, args ...interface{}) {
	z.zapLogger.Sugar().Debugf(template, args)
}

func (z *zapLog) Infof(template string, args ...interface{}) {
	z.zapLogger.Sugar().Infof(template, args)
}

func (z *zapLog) Warnf(template string, args ...interface{}) {
	z.zapLogger.Sugar().Warnf(template, args)
}

func (z *zapLog) Errorf(template string, args ...interface{}) {
	z.zapLogger.Sugar().Errorf(template, args)
}

func (z *zapLog) DPanicf(template string, args ...interface{}) {
	z.zapLogger.Sugar().DPanicf(template, args)
}

func (z *zapLog) Panicf(template string, args ...interface{}) {
	z.zapLogger.Sugar().Panicf(template, args)
}

func (z *zapLog) Fatalf(template string, args ...interface{}) {
	z.zapLogger.Sugar().Fatalf(template, args)
}

// Global logger method

// Info method output info level log.
func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	std.Debug(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.Infow(msg, keysAndValues)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.Warnw(msg, keysAndValues)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.Errorw(msg, keysAndValues)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	std.DPanicw(msg, keysAndValues)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.Panicw(msg, keysAndValues)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.Fatalw(msg, keysAndValues)
}

func Debugf(template string, args ...interface{}) {
	std.Debugf(template, args)
}

func Infof(template string, args ...interface{}) {
	std.Infof(template, args)
}

func Warnf(template string, args ...interface{}) {
	std.Warnf(template, args)
}

func Errorf(template string, args ...interface{}) {
	std.Errorf(template, args)
}

func DPanicf(template string, args ...interface{}) {
	std.DPanicf(template, args)
}

func Panicf(template string, args ...interface{}) {
	std.Panicf(template, args)
}

func Fatalf(template string, args ...interface{}) {
	std.Fatalf(template, args)
}
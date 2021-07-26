package log

import (
	"cloud_callback/internal/conf"
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

	// Flush calls the underlying Core's Sync method, flushing any buffered
	// log entries. Applications should take care to call Sync before exiting.
	Flush()
}

type zapLogger struct {
	*zap.Logger
}

var (
	std = initLogger(DefaultOptions())
	mu  sync.Mutex
)

func NewLogger(config *conf.Configs) (Logger, error) {
	mu.Lock()
	defer mu.Unlock()
	// copier
	var opts Options
	if err := copier.Copy(&opts, config.LogConf); err != nil {
		return nil, err
	}
	if err := validate(&opts); err != nil {
		return nil, err
	}
	std = initLogger(&opts)
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

func initLogger(opts *Options) *zapLogger {
	core := opts.initCore()
	logger := zap.New(core,
		zap.AddStacktrace(zapcore.PanicLevel), // 配置超过panic级别，不输出栈堆信息
		zap.AddCallerSkip(1))
	zap.RedirectStdLog(logger)
	return &zapLogger{logger}
}

func (lg *zapLogger) Flush() {
	_ = lg.Sync()
}

// StdErrLogger 转成标准库logger,级别是error
func StdErrLogger() *log.Logger {
	if l, err := zap.NewStdLogAt(std.Logger, zapcore.ErrorLevel); err == nil {
		std.Error("StdErrLogger-->NewStdLogAt", zap.Error(err))
		return l
	}
	return nil
}

// StdInfoLogger 转成标准库logger,级别是info
func StdInfoLogger() *log.Logger {
	l, err := zap.NewStdLogAt(std.Logger, zapcore.InfoLevel)
	if err != nil {
		std.Error("StdInfoLogger-->NewStdLogAt", zap.Error(err))
		return nil
	}
	return l
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

/*
old function
func validate(config *conf.Configs) error {
	if config == nil {
		return fmt.Errorf("log config must not nil")
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(config.LogConf.Level)); err != nil {
		return err
	}

	format := strings.ToLower(config.LogConf.Format)
	if format != consoleFormat && format != jsonFormat {
		return fmt.Errorf("log format not valid: %q", format)
	}
	return nil
}

func initLogger(config *conf.Configs) *Logger {
	// 配置level
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(config.LogConf.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 输出level格式
	encodeLevel := zapcore.CapitalLevelEncoder // 默认编码为大写level
	if config.LogConf.Format == consoleFormat && config.LogConf.EnableColor {
		// 给Level输出颜色
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// 配置输出格式
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       config.LogConf.Development,
		DisableCaller:     config.LogConf.DisableCaller,
		DisableStacktrace: config.LogConf.DisableStacktrace,
		Sampling: &zap.SamplingConfig{ // 控制日志输出，减少CPU消耗
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         config.LogConf.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{config.LogConf.OutputPath,"stdout"},
		ErrorOutputPaths: []string{config.LogConf.ErrorOutputPath,"stderr"},
	}
	var err error
	// 配置超过panic级别，不输出栈堆信息
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	zap.RedirectStdLog(l)
	return &Logger{l}
}
*/

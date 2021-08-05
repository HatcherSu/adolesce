package log

import (
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	consoleFormat = "console"
	jsonFormat    = "json"
)

// init zapcore encoder
func (o *Options) initEncoder() zapcore.Encoder {
	cfg := o.defaultEncoderConfig()
	switch o.Format {
	case consoleFormat:
		return zapcore.NewConsoleEncoder(cfg)
	case jsonFormat:
		return zapcore.NewJSONEncoder(cfg)
	default:
	}
	return nil
}

func (o *Options) defaultEncoderConfig() zapcore.EncoderConfig {
	// 输出level格式
	encodeLevel := zapcore.CapitalLevelEncoder // 默认编码为大写level
	if o.Format == consoleFormat && o.EnableColor {
		// 给Level输出颜色
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return zapcore.EncoderConfig{
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
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

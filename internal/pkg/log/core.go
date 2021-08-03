package log

import (
	"fmt"
	"github.com/fatih/color"
	rotate "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

func (o *Options) initCore() zapcore.Core {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// log file writer
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	infoWriter := getCutLogWriter(o.OutputPath)
	errorWriter := getCutLogWriter(o.ErrorOutputPath)

	// encoder
	encoder := o.initEncoder()

	return zapcore.NewTee(
		zapcore.NewCore(encoder, consoleDebugging, infoLevel),
		zapcore.NewCore(encoder, consoleErrors, errorLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel))
}

// 获取日志分割
func getCutLogWriter(filePath string) io.Writer {
	hook, err := rotate.New(
		filePath+".%Y-%m-%d",
		rotate.WithLinkName(filePath),
		rotate.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotate.WithRotationTime(time.Duration(24*7)*time.Hour),
	)
	if err != nil {
		fmt.Printf("rotate New error :%v %v\n", color.RedString("Error:"), err)
		return nil
	}
	return hook
}

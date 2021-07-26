package log

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

// Options logs available for external configuration
type Options struct {
	Level             string `json:"level"`
	Development       bool   `json:"development"`
	Format            string `json:"format"`
	EnableColor       bool   `json:"enable_color"`
	DisableCaller     bool   `json:"disable_caller"`
	DisableStacktrace bool   `json:"disable_stacktrace"`
	OutputPath        string `json:"output_path"`
	ErrorOutputPath   string `json:"error_output_path"`
}

// AddFlags todo 添加flag,从程序中输入
func (o *Options) AddFlags() {

}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func parseLevel(level string) zapcore.Level {
	l := strings.ToLower(level)
	return levelMap[l]
}

func DefaultOptions() *Options {
	// 配置level
	return &Options{
		Level:             "info",
		Development:       true,
		Format:            consoleFormat,
		EnableColor:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPath:        "stdout",
		ErrorOutputPath:   "stderr",
	}
}

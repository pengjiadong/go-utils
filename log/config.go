package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Option 日志参数
type Option struct {
	// 文件路径
	FilePath string
	// 文件最大保存时间
	MaxAge time.Duration
	// 日志切割时间间隔
	RotationTime time.Duration
	// 日志级别
	Level logrus.Level
	// 日志分割
	SplitLevel []logrus.Level
}

// NewOption 创建日志配置
func NewOption(opts ...func(*Option)) Option {
	option := &Option{
		FilePath:     "%Y%m%d%H.log",
		MaxAge:       time.Hour * 24 * 7,
		RotationTime: time.Hour * 24,
		Level:        logrus.TraceLevel,
	}
	for _, f := range opts {
		f(option)
	}
	return *option
}

// WithFilePath WithFilePath
func WithFilePath(filepath string) func(*Option) {
	return func(o *Option) {
		o.FilePath = filepath
	}
}

// WithMaxAge WithMaxAge
func WithMaxAge(d time.Duration) func(*Option) {
	return func(o *Option) {
		o.MaxAge = d
	}
}

// WithRotationTime WithRotationTime
func WithRotationTime(d time.Duration) func(*Option) {
	return func(o *Option) {
		o.RotationTime = d
	}
}

// WithLevel WithLevel
func WithLevel(level logrus.Level) func(*Option) {
	return func(o *Option) {
		o.Level = level
	}
}

// WithSplitLevel WithSplitLevel
func WithSplitLevel(levels ...logrus.Level) func(*Option) {
	return func(o *Option) {
		o.SplitLevel = levels
	}
}

package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Config 日志参数
type Config struct {
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
func NewOption(opts ...func(*Config)) Config {
	option := &Config{
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
func WithFilePath(filepath string) func(*Config) {
	return func(o *Config) {
		o.FilePath = filepath
	}
}

// WithMaxAge WithMaxAge
func WithMaxAge(d time.Duration) func(*Config) {
	return func(o *Config) {
		o.MaxAge = d
	}
}

// WithRotationTime WithRotationTime
func WithRotationTime(d time.Duration) func(*Config) {
	return func(o *Config) {
		o.RotationTime = d
	}
}

// WithLevel WithLevel
func WithLevel(level logrus.Level) func(*Config) {
	return func(o *Config) {
		o.Level = level
	}
}

// WithSplitLevel WithSplitLevel
func WithSplitLevel(levels ...logrus.Level) func(*Config) {
	return func(o *Config) {
		o.SplitLevel = levels
	}
}

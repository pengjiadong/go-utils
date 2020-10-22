package log

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Dir          string `mapstructure:"dir"`          // 日志目录
	MaxAge       uint   `mapstructure:"maxAge"`       // 保存时间
	RotationTime uint   `mapstructure:"rotationTime"` // 滚动时间
	Level        string `mapstructure:"level"`        // 日志级别
	SplitLevel   string `mapstructure:"splitLevel"`   // 分割日志级别
}

func LoadOption() (opt Option) {
	cfg := Config{}
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}
	opt = cfg.Convert()
	return
}

func (c *Config) Convert() Option {
	opt := NewOption()
	if c.Dir != "" {
		opt.DirPath = c.Dir
	}
	if c.MaxAge > 0 {
		opt.MaxAge = time.Hour * 24 * time.Duration(c.MaxAge)
	}
	if c.RotationTime > 0 {
		opt.RotationTime = time.Hour * time.Duration(c.RotationTime)
	}
	if c.Level != "" {
		l, err := logrus.ParseLevel(c.Level)
		if err != nil {
			l = logrus.InfoLevel
		}
		opt.Level = l
	}
	if c.SplitLevel != "" {
		levels := strings.Split(c.SplitLevel, ",")
		for _, level := range levels {
			l, err := logrus.ParseLevel(level)
			if err == nil {
				opt.SplitLevel = append(opt.SplitLevel, l)
			}
		}
	}
	return *opt
}

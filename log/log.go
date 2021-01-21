package log

import (
	"fmt"
	"os"
	"path"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	extionName string = ".log"
)

// Init 初始化日志配置
func Init(option ...Option) {
	var opt Option
	if len(option) == 0 {
		opt = LoadOption()
	} else {
		opt = option[0]
	}
	ConfigLocalFilesystemLogger(opt)
}

// ConfigLocalFilesystemLogger 切割日志和清理过期日志
func ConfigLocalFilesystemLogger(option Option) {
	// 创建目录
	CreateLogDir(path.Join(option.DirPath, option.FilePath))
	writerMap := lfshook.WriterMap{}
	for _, level := range option.SplitLevel {
		if level <= option.Level {
			writer, err := rotatelogs.New(
				path.Join(option.DirPath, fmt.Sprintf("%s_%s%s", option.FilePath, level.String(), extionName)),
				rotatelogs.WithMaxAge(option.MaxAge),
				rotatelogs.WithRotationTime(option.RotationTime),
			)
			if err != nil {
				logrus.Fatal("init split log failed, err:", err)
			}
			writerMap[level] = writer
		}
	}

	writerAll, err := rotatelogs.New(
		path.Join(option.DirPath, fmt.Sprintf("%s_%s%s", option.FilePath, "all", extionName)),
		rotatelogs.WithMaxAge(option.MaxAge),
		rotatelogs.WithRotationTime(option.RotationTime),
	)
	if err != nil {
		logrus.Fatal("init log failed, err:", err)
	}
	for _, level := range logrus.AllLevels {
		if _, ok := writerMap[level]; ok {
			continue
		}
		if level <= option.Level {
			writerMap[level] = writerAll
		}
	}

	logrus.SetLevel(option.Level)
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"})

	lfHook := lfshook.NewHook(writerMap, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"})
	logrus.AddHook(lfHook)
}

// CreateLogDir 创建目录
func CreateLogDir(filePath string) {
	dirPath := path.Dir(filePath)
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

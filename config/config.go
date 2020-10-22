package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(name string, path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// UnmarshalKey .
func UnmarshalKey(key string, rawVal interface{}) {
	err := viper.UnmarshalKey(key, rawVal)
	HandleError(err)
	logrus.Errorf("config %s: %v", key, rawVal)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

package config

import "fmt"

// RedisConfig 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

// MySQLConfig 配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func (c MySQLConfig) String() string {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Db)
	return dns
}

type MongoConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func (c MongoConfig) String() string {
	// mongodb://username:password@127.0.0.1:8888/db
	dns := fmt.Sprintf("mongodb://%s:%s@%s/%s", c.Username, c.Password, c.Host, c.Db)
	return dns
}

func GetMySQL(name string) MySQLConfig {
	c := MySQLConfig{}
	UnmarshalKey(name, &c)
	return c
}

func GetRedis(name string) RedisConfig {
	c := RedisConfig{}
	UnmarshalKey(name, &c)
	return c
}

func GetMongo(name string) MongoConfig {
	c := MongoConfig{}
	UnmarshalKey(name, &c)
	return c
}

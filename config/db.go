package config

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

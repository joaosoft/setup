package gosetup

import (
	"github.com/joaosoft/go-log/service"
)

var global = make(map[string]interface{})
var log = golog.NewLogDefault("go-setup", golog.InfoLevel)

func getDefaultNsqConfig() *NsqConfig {
	if value, exists := global["nsq"]; exists {
		return value.(*NsqConfig)
	}
	return nil
}

func getDefaultSqlConfig() *SqlConfig {
	if value, exists := global["sql"]; exists {
		return value.(*SqlConfig)
	}
	return nil
}

func getDefaultRedisConfig() *RedisConfig {
	if value, exists := global["redis"]; exists {
		return value.(*RedisConfig)
	}
	return nil
}

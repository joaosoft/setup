package gosetup

import (
	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

var global = make(map[string]interface{})
var log = golog.NewLogDefault("go-setup", golog.InfoLevel)

func init() {
	global["path"] = defaultPath
}

func getDefaultNsqConfig() *gomanager.NSQConfig {
	if value, exists := global["nsq"]; exists {
		return value.(*gomanager.NSQConfig)
	}
	return nil
}

func getDefaultSqlConfig() *gomanager.DBConfig {
	if value, exists := global["sql"]; exists {
		return value.(*gomanager.DBConfig)
	}
	return nil
}

func getDefaultRedisConfig() *gomanager.RedisConfig {
	if value, exists := global["redis"]; exists {
		return value.(*gomanager.RedisConfig)
	}
	return nil
}

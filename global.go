package setup

import (
	"github.com/joaosoft/manager"
)

var global = make(map[string]interface{})

func init() {
	global[path_key] = defaultPath
}

func getDefaultNsqConfig() *manager.NSQConfig {
	if value, exists := global["nsq"]; exists {
		return value.(*manager.NSQConfig)
	}
	return nil
}

func getDefaultSqlConfig() *manager.DBConfig {
	if value, exists := global["sql"]; exists {
		return value.(*manager.DBConfig)
	}
	return nil
}

func getDefaultRedisConfig() *manager.RedisConfig {
	if value, exists := global["redis"]; exists {
		return value.(*manager.RedisConfig)
	}
	return nil
}

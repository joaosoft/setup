package gosetup

import (
	"os"

	logger "github.com/joaosoft/go-log/service"
	writer "github.com/joaosoft/go-writer/service"

)

var global = make(map[string]interface{})
var log = logger.NewLog(
	logger.WithLevel(logger.InfoLevel),
	logger.WithFormatHandler(writer.JsonFormatHandler),
	logger.WithWriter(os.Stdout)).WithPrefixes(map[string]interface{}{
	"level":   logger.LEVEL,
	"time":    logger.TIME,
	"service": "go-setup"})

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

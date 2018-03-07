package gomock

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

var global = make(map[string]interface{})
var log = logger.WithFields(logger.Fields{
	"application": "go-mock",
})

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logger.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logger.DebugLevel)
}

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

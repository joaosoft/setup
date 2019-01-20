package setup

import (
	"fmt"

	manager "github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Setup SetupConfig `json:"setup"`
}

// goSetupConfig ...
type SetupConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

// Configurations ...
type Configurations struct {
	Connections Connections `json:"connections"`
}

// Connections ...
type Connections struct {
	NsqConfig   manager.NSQConfig   `json:"nsq"`
	SqlConfig   manager.DBConfig    `json:"sql"`
	RedisConfig manager.RedisConfig `json:"redis"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}

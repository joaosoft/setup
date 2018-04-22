package gosetup

import (
	"github.com/joaosoft/go-manager/service"
)

// appConfig ...
type appConfig struct {
	GoSetup GoSetupConfig `json:"gosetup"`
}

// goSetupConfig ...
type GoSetupConfig struct {
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
	NsqConfig   gomanager.NSQConfig   `json:"nsq"`
	SqlConfig   gomanager.DBConfig    `json:"sql"`
	RedisConfig gomanager.RedisConfig `json:"redis"`
}

package gosetup

import (
	"strings"

	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

// GoSetupOption ...
type GoSetupOption func(gosetup *GoSetup)

// Reconfigure ...
func (gosetup *GoSetup) Reconfigure(options ...GoSetupOption) {
	for _, option := range options {
		option(gosetup)
	}
}

// WithPath ...
func WithPath(path string) GoSetupOption {
	return func(gosetup *GoSetup) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global["path"] = path
		}
	}
}

// WithServices ...
func WithServices(services []*Services) GoSetupOption {
	return func(gosetup *GoSetup) {
		gosetup.services = services
	}
}

// WithRunInBackground ...
func WithRunInBackground(runInBackground bool) GoSetupOption {
	return func(gosetup *GoSetup) {
		gosetup.runInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) GoSetupOption {
	return func(gosetup *GoSetup) {
		config := &Configurations{}
		if _, err := readFile(file, config); err != nil {
			panic(err)
		}
		gosetup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *gomanager.RedisConfig) GoSetupOption {
	return func(gosetup *GoSetup) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *gomanager.DBConfig) GoSetupOption {
	return func(gosetup *GoSetup) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *gomanager.NSQConfig) GoSetupOption {
	return func(gosetup *GoSetup) {
		global["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) GoSetupOption {
	return func(gosetup *GoSetup) {
		gosetup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithLogger ...
func WithLogger(logger golog.ILog) GoSetupOption {
	return func(gosetup *GoSetup) {
		log = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level golog.Level) GoSetupOption {
	return func(gosetup *GoSetup) {
		log.SetLevel(level)
	}
}

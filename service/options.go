package gosetup

import (
	"strings"

	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

// SetupOption ...
type SetupOption func(setup *Setup)

// Reconfigure ...
func (setup *Setup) Reconfigure(options ...SetupOption) {
	for _, option := range options {
		option(setup)
	}
}

// WithPath ...
func WithPath(path string) SetupOption {
	return func(setup *Setup) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global[path_key] = path
		}
	}
}

// WithServices ...
func WithServices(services []*Services) SetupOption {
	return func(setup *Setup) {
		setup.services = services
	}
}

// WithRunInBackground ...
func WithRunInBackground(runInBackground bool) SetupOption {
	return func(setup *Setup) {
		setup.runInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) SetupOption {
	return func(setup *Setup) {
		config := &Configurations{}
		if _, err := readFile(file, config); err != nil {
			panic(err)
		}
		setup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *gomanager.RedisConfig) SetupOption {
	return func(setup *Setup) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *gomanager.DBConfig) SetupOption {
	return func(setup *Setup) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *gomanager.NSQConfig) SetupOption {
	return func(setup *Setup) {
		global["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) SetupOption {
	return func(setup *Setup) {
		setup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithLogger ...
func WithLogger(logger golog.ILog) SetupOption {
	return func(setup *Setup) {
		log = logger
		setup.logIsExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level golog.Level) SetupOption {
	return func(setup *Setup) {
		log.SetLevel(level)
	}
}

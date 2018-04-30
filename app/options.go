package gosetup

import (
	"strings"

	golog "github.com/joaosoft/go-log/app"
	gomanager "github.com/joaosoft/go-manager/app"
)

// setupOption ...
type setupOption func(setup *Setup)

// Reconfigure ...
func (setup *Setup) Reconfigure(options ...setupOption) {
	for _, option := range options {
		option(setup)
	}
}

// WithPath ...
func WithPath(path string) setupOption {
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
func WithServices(services []*Services) setupOption {
	return func(setup *Setup) {
		setup.services = services
	}
}

// WithRunInBackground ...
func WithRunInBackground(runInBackground bool) setupOption {
	return func(setup *Setup) {
		setup.isToRunInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) setupOption {
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
func WithRedisConfiguration(config *gomanager.RedisConfig) setupOption {
	return func(setup *Setup) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *gomanager.DBConfig) setupOption {
	return func(setup *Setup) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *gomanager.NSQConfig) setupOption {
	return func(setup *Setup) {
		global["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) setupOption {
	return func(setup *Setup) {
		setup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithLogger ...
func WithLogger(logger golog.ILog) setupOption {
	return func(setup *Setup) {
		log = logger
		setup.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level golog.Level) setupOption {
	return func(setup *Setup) {
		log.SetLevel(level)
	}
}

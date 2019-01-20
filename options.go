package setup

import (
	"strings"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
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
		setup.isToRunInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) SetupOption {
	return func(setup *Setup) {
		config := &Configurations{}
		if _, err := ReadFile(file, config); err != nil {
			panic(err)
		}
		setup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *manager.RedisConfig) SetupOption {
	return func(setup *Setup) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *manager.DBConfig) SetupOption {
	return func(setup *Setup) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *manager.NSQConfig) SetupOption {
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
func WithLogger(logger logger.ILogger) SetupOption {
	return func(setup *Setup) {
		setup.logger = logger
		setup.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) SetupOption {
	return func(setup *Setup) {
		setup.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) SetupOption {
	return func(setup *Setup) {
		setup.pm = mgr
	}
}

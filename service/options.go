package gosetup

import (
	"strings"

	logger "github.com/joaosoft/go-log/service"
)

// GoSetupOption ...
type GoSetupOption func(gosetup *gosetup)

// Reconfigure ...
func (gosetup *gosetup) Reconfigure(options ...GoSetupOption) {
	for _, option := range options {
		option(gosetup)
	}
}

// WithPath ...
func WithPath(path string) GoSetupOption {
	return func(gosetup *gosetup) {
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
	return func(gosetup *gosetup) {
		gosetup.services = services
	}
}

// WithRunInBackground ...
func WithRunInBackground(runInBackground bool) GoSetupOption {
	return func(gosetup *gosetup) {
		gosetup.runInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) GoSetupOption {
	return func(gosetup *gosetup) {
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
func WithRedisConfiguration(config *RedisConfig) GoSetupOption {
	return func(gosetup *gosetup) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *SqlConfig) GoSetupOption {
	return func(gosetup *gosetup) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *NsqConfig) GoSetupOption {
	return func(gosetup *gosetup) {
		global["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) GoSetupOption {
	return func(gosetup *gosetup) {
		gosetup.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithLogger ...
func WithLogger(logger logger.ILog) GoSetupOption {
	return func(gosetup *gosetup) {
		log = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GoSetupOption {
	return func(gosetup *gosetup) {
		log.SetLevel(level)
	}
}

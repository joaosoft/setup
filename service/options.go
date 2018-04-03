package gomock

import (
	"strings"

	logger "github.com/joaosoft/go-log/service"
)

// GoMockOption ...
type GoMockOption func(gomock *GoMock)

// Reconfigure ...
func (gomock *GoMock) Reconfigure(options ...GoMockOption) {
	for _, option := range options {
		option(gomock)
	}
}

// WithPath ...
func WithPath(path string) GoMockOption {
	return func(gomock *GoMock) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global["path"] = path
		}
	}
}

// WithServices ...
func WithServices(services []*Services) GoMockOption {
	return func(gomock *GoMock) {
		gomock.services = services
	}
}

// WithRunInBackground ...
func WithRunInBackground(runInBackground bool) GoMockOption {
	return func(gomock *GoMock) {
		gomock.runInBackground = runInBackground
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) GoMockOption {
	return func(gomock *GoMock) {
		config := &Configurations{}
		if _, err := readFile(file, config); err != nil {
			panic(err)
		}
		gomock.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *RedisConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["redis"] = config
	}
}

// WithSqlConfiguration ...
func WithSqlConfiguration(config *SqlConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["sql"] = config
	}
}

// WithNsqConfiguration ...
func WithNsqConfiguration(config *NsqConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) GoMockOption {
	return func(gomock *GoMock) {
		gomock.Reconfigure(
			WithSqlConfiguration(&config.Connections.SqlConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNsqConfiguration(&config.Connections.NsqConfig))
	}
}

// WithLogger ...
func WithLogger(logger logger.Log) GoMockOption {
	return func(gomock *GoMock) {
		log = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GoMockOption {
	return func(gomock *GoMock) {
		log.SetLevel(level)
	}
}

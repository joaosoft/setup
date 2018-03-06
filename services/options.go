package gomock

import (
	"strings"

	"github.com/sirupsen/logrus"
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

// WithRunInBackground ...
func WithRunInBackground(background bool) GoMockOption {
	return func(gomock *GoMock) {
		gomock.background = background
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
			WithSQLConfiguration(&config.Connections.SQLConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNSQConfiguration(&config.Connections.NSQConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *RedisConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["redis"] = config
	}
}

// WithSQLConfiguration ...
func WithSQLConfiguration(config *SQLConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["sql"] = config
	}
}

// WithNSQConfiguration ...
func WithNSQConfiguration(config *NSQConfig) GoMockOption {
	return func(gomock *GoMock) {
		global["nsq"] = config
	}
}

// WithLogLevel ...
func WithLogLevel(level logrus.Level) GoMockOption {
	return func(gomock *GoMock) {
		logrus.SetLevel(level)
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) GoMockOption {
	return func(gomock *GoMock) {
		gomock.Reconfigure(
			WithSQLConfiguration(&config.Connections.SQLConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNSQConfiguration(&config.Connections.NSQConfig))
	}
}

package gomock

import (
	"fmt"
	"strings"
)

// MockOption ...
type MockOption func(mock *Mock)

// WithPath ...
func WithPath(path string) MockOption {
	return func(mock *Mock) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global["path"] = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) MockOption {
	return func(mock *Mock) {
		mock.background = background
	}
}

// WithConfiguration ...
func WithConfiguration(file string) MockOption {
	return func(mock *Mock) {
		app := &App{}
		if _, err := readFile(file, app); err != nil {
			panic(err)
		}
		fmt.Println(app)
		mock.Reconfigure(
			WithConfigurationSQL(&app.Config.SQL),
			WithConfigurationRedis(&app.Config.Redis),
			WithConfigurationNSQ(&app.Config.NSQ))
	}
}

// WithConfigurationRedis ...
func WithConfigurationRedis(config *ConfigRedis) MockOption {
	return func(mock *Mock) {
		mock.defaults["redis"] = config
	}
}

// WithConfigurationSQL ...
func WithConfigurationSQL(config *ConfigSQL) MockOption {
	return func(mock *Mock) {
		mock.defaults["sql"] = config
	}
}

// WithConfigurationNSQ ...
func WithConfigurationNSQ(config *ConfigNSQ) MockOption {
	return func(mock *Mock) {
		mock.defaults["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) MockOption {
	return func(mock *Mock) {
		mock.Reconfigure(
			WithConfigurationSQL(&config.SQL),
			WithConfigurationRedis(&config.Redis),
			WithConfigurationNSQ(&config.NSQ))
	}
}

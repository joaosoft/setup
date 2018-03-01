package gomock

import (
	"fmt"
	"strings"
)

// GoMockOption ...
type GoMockOption func(gomock *GoMock)

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

// WithConfiguration ...
func WithConfiguration(file string) GoMockOption {
	return func(gomock *GoMock) {
		app := &App{}
		if _, err := readFile(file, app); err != nil {
			panic(err)
		}
		fmt.Println(app)
		gomock.Reconfigure(
			WithConfigurationSQL(&app.Config.SQL),
			WithConfigurationRedis(&app.Config.Redis),
			WithConfigurationNSQ(&app.Config.NSQ))
	}
}

// WithConfigurationRedis ...
func WithConfigurationRedis(config *ConfigRedis) GoMockOption {
	return func(gomock *GoMock) {
		gomock.defaults["redis"] = config
	}
}

// WithConfigurationSQL ...
func WithConfigurationSQL(config *ConfigSQL) GoMockOption {
	return func(gomock *GoMock) {
		gomock.defaults["sql"] = config
	}
}

// WithConfigurationNSQ ...
func WithConfigurationNSQ(config *ConfigNSQ) GoMockOption {
	return func(gomock *GoMock) {
		gomock.defaults["nsq"] = config
	}
}

// WithConfigurations ...
func WithConfigurations(config *Configurations) GoMockOption {
	return func(gomock *GoMock) {
		gomock.Reconfigure(
			WithConfigurationSQL(&config.SQL),
			WithConfigurationRedis(&config.Redis),
			WithConfigurationNSQ(&config.NSQ))
	}
}

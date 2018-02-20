package gomock

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
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
			mock.path = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) MockOption {
	return func(mock *Mock) {
		mock.background = background
	}
}

// Reconfigure ...
func (mock *Mock) Reconfigure(options ...MockOption) {
	for _, option := range options {
		option(mock)
	}
}

// Mock ...
type Mock struct {
	services   []*Services
	path       string
	background bool
}

// NewGoMock ...
func NewGoMock(options ...MockOption) *Mock {
	mock := &Mock{
		path:       "",
		background: background,
	}

	mock.Reconfigure(options...)

	return mock
}

// RunSingle ...
func (mock *Mock) RunSingle(file string) error {
	fmt.Println(":: Initializing Mock Service")

	if mock.path != "" {
		file = mock.path + file
	}
	if err := mock.setup(file); err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println(":: Mock Services Started")

	if !mock.background {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		mock.Stop()
	}

	return nil
}

// Run ...
func (mock *Mock) Run() error {
	fmt.Println(":: Initializing Mock Service")

	files, err := filepath.Glob(mock.path + "*.json")
	if err != nil {
		return err
	}
	if err := mock.setup(files...); err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println(":: Mock Services Started")

	mock.wait()

	return nil
}

// wait ...
func (mock *Mock) wait() {
	if !mock.background {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		mock.Stop()
	}
}

// Stop ...
func (mock *Mock) Stop() error {

	fmt.Println(":: Stopping Mock Service")
	if err := mock.teardown(); err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println(":: Stoped Mock Service")

	return nil
}

func (mock *Mock) setup(files ...string) error {
	for _, file := range files {
		config := &Services{}
		if err := config.fromFile(file); err != nil {
			return err
		}

		if err := config.setup(); err != nil {
			return err
		}
		mock.services = append(mock.services, config)
	}

	return nil
}

func (mock *Mock) teardown() error {
	for _, service := range mock.services {
		service.teardown()
	}
	return nil
}

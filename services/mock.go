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
	return func(gomock *Mock) {
		if path != "" {
			gomock.path = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) MockOption {
	return func(gomock *Mock) {
		gomock.background = background
	}
}

// Reconfigure ...
func (gomock *Mock) Reconfigure(options ...MockOption) {
	for _, option := range options {
		option(gomock)
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
		path:       path,
		background: background,
	}

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (mock *Mock) Run() error {
	fmt.Println("---------- STARTING ----------")
	fmt.Println(":: Initializing Mock Service")

	if err := mock.setup(); err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println(":: Mock Services Started")

	if !mock.background {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		fmt.Println("\n---------- SHUTTING DOWN ----------")
		mock.Stop()
	}

	return nil
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

func (gomock *Mock) setup() error {
	err := filepath.Walk(gomock.path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
			config := &Services{}
			if err := config.fromFile(path); err != nil {
				return err
			}

			config.setup()
			gomock.services = append(gomock.services, config)
		}

		return nil
	})

	return err
}

func (gomock *Mock) teardown() error {
	for _, service := range gomock.services {
		service.teardown()
	}
	return nil
}

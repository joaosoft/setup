package gomock

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// Mock ...
type Mock struct {
	services   []*ServicesConfig
	background bool
	defaults   map[string]interface{}
}

// Reconfigure ...
func (mock *Mock) Reconfigure(options ...MockOption) {
	for _, option := range options {
		option(mock)
	}
}

// NewGoMock ...
func NewGoMock(options ...MockOption) *Mock {
	fmt.Println(":: Starting Mock Service")
	mock := &Mock{
		background: background,
		defaults:   make(map[string]interface{}),
	}

	global["path"] = defaultPath

	mock.Reconfigure(options...)

	return mock
}

// RunSingleNoWait ...
func (mock *Mock) RunSingleNoWait(file string) error {
	if err := mock.setup(file); err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

// RunSingle ...
func (mock *Mock) RunSingleWait(file string) error {
	mock.RunSingleNoWait(file)
	mock.wait(true)

	return nil
}

// Run ...
func (mock *Mock) Run() error {
	files, err := filepath.Glob(global["path"] + "*.json")
	if err != nil {
		return err
	}
	if err := mock.setup(files...); err != nil {
		log.Panic(err)
		return err
	}
	mock.wait(false)

	return nil
}

// wait ...
func (mock *Mock) wait(force bool) {
	if !mock.background || force {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}

	mock.Stop()
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
		fmt.Println(fmt.Sprintf("\nSTARTING: setup [ %s ]", file))

		config := &ServicesConfig{File: file}
		if err := config.fromFile(file); err != nil {
			return err
		}

		if err := config.setup(mock.defaults); err != nil {
			return err
		}
		mock.services = append(mock.services, config)

		fmt.Println(fmt.Sprintf("FINISHED: setup [ %s ]", file))
	}

	return nil
}

func (mock *Mock) teardown() error {
	for _, service := range mock.services {
		fmt.Println(fmt.Sprintf("\nSTARTING: teardown [ %s ]", service.File))
		service.teardown(mock.defaults)
		fmt.Println(fmt.Sprintf("FINISHED: teardown [ %s ]", service.File))
	}
	return nil
}

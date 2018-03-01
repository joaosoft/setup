package gomock

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// GoMock ...
type GoMock struct {
	services   []*ServicesConfig
	background bool
	defaults   map[string]interface{}
	started    bool
}

// Reconfigure ...
func (gomock *GoMock) Reconfigure(options ...GoMockOption) {
	for _, option := range options {
		option(gomock)
	}
}

// NewGoMock ...
func NewGoMock(options ...GoMockOption) *GoMock {
	fmt.Println(":: Starting GoMock Service")
	mock := &GoMock{
		background: background,
		defaults:   make(map[string]interface{}),
	}

	global["path"] = defaultPath

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gomock *GoMock) Run() error {
	files, err := filepath.Glob(global["path"] + "*.json")
	if err != nil {
		return err
	}
	if err := gomock.setup(files...); err != nil {
		log.Panic(err)
		return err
	}
	gomock.wait(false)

	return nil
}

// RunSingle ...
func (gomock *GoMock) RunSingle(file string) error {
	if err := gomock.setup(file); err != nil {
		log.Panic(err)
		return err
	}
	gomock.wait(false)

	return nil
}

// wait ...
func (gomock *GoMock) wait(force bool) {
	if !gomock.started && (gomock.background || (!gomock.background && force)) {
		gomock.started = true
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}

	gomock.Stop()
}

// Stop ...
func (gomock *GoMock) Stop() error {

	fmt.Println(":: Stopping GoMock Service")
	if err := gomock.teardown(); err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println(":: Stoped GoMock Service")

	return nil
}

func (gomock *GoMock) setup(files ...string) error {
	for _, file := range files {
		fmt.Println(fmt.Sprintf("\nSTARTING: setup [ %s ]", file))

		config := &ServicesConfig{}
		if err := config.fromFile(file); err != nil {
			return err
		}

		if err := config.setup(gomock.defaults); err != nil {
			return err
		}
		gomock.services = append(gomock.services, config)

		fmt.Println(fmt.Sprintf("FINISHED: setup [ %s ]", file))
	}

	return nil
}

func (gomock *GoMock) teardown() error {
	for _, service := range gomock.services {
		fmt.Println(fmt.Sprintf("\nSTARTING: teardown [ %s ]", service.File))
		service.teardown(gomock.defaults)
		fmt.Println(fmt.Sprintf("FINISHED: teardown [ %s ]", service.File))
	}
	return nil
}

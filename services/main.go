package gomock

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"
)

// GoMock ...
type GoMock struct {
	services   []*Services
	runner     IRunner
	background bool
}

// NewGoMock ...make
func NewGoMock(options ...GoMockOption) *GoMock {
	log.Info("starting GoMock Service")
	mock := &GoMock{
		background: background,
		services:   make([]*Services, 0),
	}

	global["path"] = defaultPath

	// load configuration file
	app := &App{}
	if _, err := readFile("config/app.json", app); err != nil {
		log.Error(err)
	} else {
		level, _ := logrus.ParseLevel(app.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gomock *GoMock) Run() error {
	files, err := filepath.Glob(global["path"].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := gomock.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (gomock *GoMock) RunSingle(file string) error {
	if err := gomock.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (gomock *GoMock) Stop() error {
	if err := gomock.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (gomock *GoMock) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := readFile(file, servicesOnFile); err != nil {
			return err
		}
		gomock.services = append(gomock.services, servicesOnFile)
	}

	gomock.runner = NewRunner(gomock.services)
	if err := gomock.runner.Setup(); err != nil {
		return err
	}

	log.Info("started all services")

	if !gomock.background {
		gomock.Wait()
	}

	return nil
}

// Wait ...
func (gomock *GoMock) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

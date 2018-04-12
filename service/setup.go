package gosetup

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joaosoft/go-log/service"
)

// GoSetup ...
type GoSetup struct {
	services        []*Services
	runner          IRunner
	runInBackground bool
	config          *AppConfig
}

// NewGoSetup ...make
func NewGoSetup(options ...GoSetupOption) *GoSetup {
	log.Info("starting GoSetup Service")

	// load configuration file
	configApp := &AppConfig{}
	if _, err := readFile("./config/app.json", configApp); err != nil {
		log.Error(err)
	} else {
		level, _ := golog.ParseLevel(configApp.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	mock := &GoSetup{
		runInBackground: background,
		services:        make([]*Services, 0),
		config:          configApp,
	}

	global["path"] = defaultPath

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gosetup *GoSetup) Run() error {
	files, err := filepath.Glob(global["path"].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := gosetup.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (gosetup *GoSetup) RunSingle(file string) error {
	if err := gosetup.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (gosetup *GoSetup) Stop() error {
	if err := gosetup.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (gosetup *GoSetup) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := readFile(file, servicesOnFile); err != nil {
			return err
		}

		array, err := load(servicesOnFile)
		if err != nil {
			return err
		}
		gosetup.services = append(gosetup.services, array...)
	}

	gosetup.runner = NewRunner(gosetup.services)
	if err := gosetup.runner.Setup(); err != nil {
		return err
	}

	log.Info("started all services")

	if !gosetup.runInBackground {
		gosetup.Wait()
	}

	return nil
}

// load recursive load services files inside every service
func load(service *Services) ([]*Services, error) {
	log.Info("loading service...")
	array := make([]*Services, 0)

	for _, file := range service.Files {
		log.Infof("loading service file %s", file)
		nextService := &Services{}
		if _, err := readFile(file, nextService); err != nil {
			return nil, err
		}

		log.Infof("getting next service...")
		if nextArray, err := load(nextService); err != nil {
			return nil, err
		} else {
			array = append(array, nextArray...)
		}
	}

	return append(array, service), nil
}

// Wait ...
func (gosetup *GoSetup) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

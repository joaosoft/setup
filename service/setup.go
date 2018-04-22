package gosetup

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"fmt"

	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

// Setup ...
type Setup struct {
	services        []*Services
	runner          IRunner
	runInBackground bool
	config          *SetupConfig
	pm              *gomanager.Manager
	logIsExternal   bool
}

// NewGoSetup ...make
func NewGoSetup(options ...SetupOption) *Setup {
	pm := gomanager.NewManager(gomanager.WithRunInBackground(false))

	log.Info("starting Setup Service")

	// load configuration file
	appConfig := &appConfig{}
	if simpleConfig, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig); err != nil {
		log.Error(err.Error())
	} else {
		pm.AddConfig("config_app", simpleConfig)
		level, _ := golog.ParseLevel(appConfig.GoSetup.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	mock := &Setup{
		runInBackground: background,
		services:        make([]*Services, 0),
		config:          &appConfig.GoSetup,
	}

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gosetup *Setup) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
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
func (gosetup *Setup) RunSingle(file string) error {
	if err := gosetup.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (gosetup *Setup) Stop() error {
	if err := gosetup.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (gosetup *Setup) execute(files []string) error {
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
func (gosetup *Setup) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

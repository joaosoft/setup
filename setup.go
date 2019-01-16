package setup

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"fmt"

	logger "github.com/joaosoft/logger"
	manager "github.com/joaosoft/manager"
)

// Setup ...
type Setup struct {
	services            []*Services
	runner              IRunner
	isToRunInBackground bool
	config              *SetupConfig
	pm                  *manager.Manager
	isLogExternal       bool
}

// NewSetup ...make
func NewSetup(options ...SetupOption) *Setup {
	pm := manager.NewManager(manager.WithRunInBackground(false))

	log.Info("starting Setup Service")

	setup := &Setup{
		isToRunInBackground: background,
		services:            make([]*Services, 0),
	}

	setup.Reconfigure(options...)

	if setup.isLogExternal {
		pm.Reconfigure(manager.WithLogger(log))
	}

	// load configuration file
	appConfig := &AppConfig{}
	if simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config.%s.json", getEnv()), appConfig); err != nil {
		log.Error(err.Error())
	} else if appConfig.Setup != nil {
		pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(appConfig.Setup.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
	}

	setup.config = appConfig.Setup

	return setup
}

// Run ...
func (setup *Setup) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := setup.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (setup *Setup) RunSingle(file string) error {
	if err := setup.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (setup *Setup) Stop() error {
	if err := setup.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (setup *Setup) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := readFile(file, servicesOnFile); err != nil {
			return err
		}

		array, err := load(servicesOnFile)
		if err != nil {
			return err
		}
		setup.services = append(setup.services, array...)
	}

	setup.runner = NewRunner(setup.services)
	if err := setup.runner.Setup(); err != nil {
		return err
	}

	log.Info("started all services")

	if !setup.isToRunInBackground {
		setup.Wait()
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
func (setup *Setup) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

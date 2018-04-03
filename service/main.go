package gosetup

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joaosoft/go-log/service"
)

// gosetup ...
type gosetup struct {
	services        []*Services
	runner          IRunner
	runInBackground bool
}

// NewGoSetup ...make
func NewGoSetup(options ...GoSetupOption) *gosetup {
	log.Info("starting gosetup Service")
	mock := &gosetup{
		runInBackground: background,
		services:        make([]*Services, 0),
	}

	global["path"] = defaultPath

	// load configuration file
	app := &App{}
	if _, err := readFile("config/app.json", app); err != nil {
		log.Error(err)
	} else {
		level, _ := golog.ParseLevel(app.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gosetup *gosetup) Run() error {
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
func (gosetup *gosetup) RunSingle(file string) error {
	if err := gosetup.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (gosetup *gosetup) Stop() error {
	if err := gosetup.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (gosetup *gosetup) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := readFile(file, servicesOnFile); err != nil {
			return err
		}

		for _, fileName := range servicesOnFile.Files {
			servicesByFile := &Services{}
			if _, err := readFile(fileName, servicesByFile); err != nil {
				return err
			}
			gosetup.services = append(gosetup.services, servicesByFile)
		}

		gosetup.services = append(gosetup.services, servicesOnFile)
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

// Wait ...
func (gosetup *gosetup) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

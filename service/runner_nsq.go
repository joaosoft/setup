package gosetup

import (
	"fmt"

	nsqlib "github.com/nsqio/go-nsq"
)

type NsqRunner struct {
	services      []NsqService
	configuration *NsqConfig
}

func NewNsqRunner(services []NsqService, config *NsqConfig) *NsqRunner {
	return &NsqRunner{
		services:      services,
		configuration: config,
	}
}

func (runner *NsqRunner) Setup() error {
	for _, service := range runner.services {
		log.Infof("creating service [ %s ] with description [ %s] ", service.Name, service.Description)

		var conn *nsqlib.Producer
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.connect(); err != nil {
				return fmt.Errorf("failed to create nsq connection")
			}
		}

		if service.Run.Setup != nil {
			for _, setup := range service.Run.Setup {
				if err := runner.runCommands(conn, &setup); err != nil {
					return err
				}

				if err := runner.runCommandsFromFile(conn, &setup); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (runner *NsqRunner) Teardown() error {
	for _, service := range runner.services {
		log.Infof("teardown service [ %s ]", service.Name)

		var conn *nsqlib.Producer
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.connect(); err != nil {
				return fmt.Errorf("failed to create nsq connection")
			}
		}

		if service.Run.Teardown != nil {
			for _, teardown := range service.Run.Teardown {
				if err := runner.runCommands(conn, &teardown); err != nil {
					return err
				}

				if err := runner.runCommandsFromFile(conn, &teardown); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (runner *NsqRunner) loadConfiguration(test NsqService) (*NsqConfig, error) {
	if test.Configuration != nil {
		return test.Configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid nsq configuration")
	}
}

func (runner *NsqRunner) runCommands(conn *nsqlib.Producer, run *NsqRun) error {
	if run.Message != nil && string(run.Message) != "" {
		log.Infof("executing nsq [ %s ] message: %s", run.Description, string(run.Message))
		if err := conn.Publish(run.Topic, run.Message); err != nil {
			return err
		}
	}

	return nil
}

func (runner *NsqRunner) runCommandsFromFile(conn *nsqlib.Producer, run *NsqRun) error {

	if run.File != "" {
		log.Infof("executing nsq commands by file [ %s ]", run.File)
		command, err := readFile(run.File, nil)
		if err != nil {
			return err
		}

		if err := conn.Publish(run.Topic, command); err != nil {
			return err
		}
	}

	return nil
}

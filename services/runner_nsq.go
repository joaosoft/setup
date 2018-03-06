package gomock

import (
	"fmt"

	nsqlib "github.com/nsqio/go-nsq"
)

type NSQRunner struct {
	services      []NSQService
	configuration *NSQConfig
}

func NewNSQRunner(services []NSQService, config *NSQConfig) *NSQRunner {
	return &NSQRunner{
		services:      services,
		configuration: config,
	}
}

func (runner *NSQRunner) Setup() error {
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

func (runner *NSQRunner) Teardown() error {
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

func (runner *NSQRunner) loadConfiguration(test NSQService) (*NSQConfig, error) {
	if test.Configuration != nil {
		return runner.configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid nsq configuration")
	}
}

func (runner *NSQRunner) runCommands(conn *nsqlib.Producer, run *NSQRun) error {
	if run.Message != nil && string(run.Message) != "" {
		log.Infof("executing nsq [ %s ] message: %s", run.Description, string(run.Message))
		if err := conn.Publish(run.Topic, run.Message); err != nil {
			return err
		}
	}

	return nil
}

func (runner *NSQRunner) runCommandsFromFile(conn *nsqlib.Producer, run *NSQRun) error {

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

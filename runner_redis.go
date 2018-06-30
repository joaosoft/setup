package setup

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alphazero/Go-Redis"
	manager "github.com/joaosoft/manager"
)

type RedisRunner struct {
	services      []RedisService
	configuration *manager.RedisConfig
}

func NewRedisRunner(services []RedisService, config *manager.RedisConfig) *RedisRunner {
	return &RedisRunner{
		services:      services,
		configuration: config,
	}
}

func (runner *RedisRunner) Setup() error {
	for _, service := range runner.services {
		log.Infof("creating service [ %s ] with description [ %s] ", service.Name, service.Description)

		var conn redis.Client
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create redis connection")
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

func (runner *RedisRunner) Teardown() error {
	for _, service := range runner.services {
		log.Infof("teardown service [ %s ]", service.Name)

		var conn redis.Client
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create redis connection")
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

func (runner *RedisRunner) loadConfiguration(test RedisService) (*manager.RedisConfig, error) {
	if test.Configuration != nil {
		return test.Configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid redis configuration")
	}
}

func (runner *RedisRunner) runCommands(conn redis.Client, run *RedisRun) error {
	for _, command := range run.Commands {
		log.Infof("executing redis command [ %s ] arguments [ %s ]", command.Command, command.Arguments)

		inputs := make([]reflect.Value, len(command.Arguments))
		for i, arg := range command.Arguments {
			inputs[i] = reflect.ValueOf(arg)
		}

		result := reflect.ValueOf(conn).MethodByName(command.Command).Call(inputs)
		if result != nil && len(result) > 0 && !result[0].IsNil() {
			return fmt.Errorf(result[0].String())
		}
	}
	return nil
}

func (runner *RedisRunner) runCommandsFromFile(conn redis.Client, run *RedisRun) error {
	for _, file := range run.Files {
		log.Infof("executing redis commands by file [ %s ]", file)

		if lines, err := readFileLines(file); err != nil {
			for _, line := range lines {
				command := strings.SplitN(line, " ", 1)
				arguments := strings.Split(command[1], " ")
				log.Infof("executing redis command [ %s ] arguments [ %s ]", command[0], command[1])
				inputs := make([]reflect.Value, len(arguments))
				for i, arg := range arguments {
					inputs[i] = reflect.ValueOf(arg)
				}

				result := reflect.ValueOf(conn).MethodByName(command[0]).Call(inputs)
				if result != nil && len(result) > 0 && !result[0].IsNil() {
					return fmt.Errorf(result[0].String())
				}
			}
		}
	}
	return nil
}

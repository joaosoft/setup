package gomock

import (
	"fmt"
	"strings"

	redis "github.com/mediocregopher/radix.v3"
)

type RedisRunner struct {
	services      []RedisService
	configuration *RedisConfig
}

func NewRedisRunner(services []RedisService, config *RedisConfig) *RedisRunner {
	return &RedisRunner{
		services:      services,
		configuration: config,
	}
}

func (runner *RedisRunner) Setup() error {
	for _, service := range runner.services {
		log.Infof("creating service [ %s ] with description [ %s] ", service.Name, service.Description)

		var conn *redis.Pool
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.connect(); err != nil {
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

		var conn *redis.Pool
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.connect(); err != nil {
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

func (runner *RedisRunner) loadConfiguration(test RedisService) (*RedisConfig, error) {
	if test.Configuration != nil {
		return runner.configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid redis configuration")
	}
}

func (runner *RedisRunner) runCommands(conn *redis.Pool, run *RedisRun) error {
	for _, command := range run.Commands {
		log.Infof("executing redis command [ %s ] arguments [ %s ]", command.Command, command.Arguments)
		if err := conn.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
			return err
		}
	}
	return nil
}

func (runner *RedisRunner) runCommandsFromFile(conn *redis.Pool, run *RedisRun) error {
	for _, file := range run.Files {
		log.Infof("executing nsq commands by file [ %s ]", file)

		if lines, err := readFileLines(file); err != nil {
			for _, line := range lines {
				command := strings.SplitN(line, " ", 1)
				arguments := strings.Split(command[1], " ")
				log.Infof("executing redis command [ %s ] arguments [ %s ]", command[0], command[1])
				if err := conn.Do(redis.Cmd(nil, command[0], arguments...)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

package setup

import (
	"database/sql"
	"fmt"

	manager "github.com/joaosoft/manager"
)

type SqlRunner struct {
	services      []SqlService
	configuration *manager.DBConfig
}

func NewSqlRunner(services []SqlService, config *manager.DBConfig) *SqlRunner {
	return &SqlRunner{
		services:      services,
		configuration: config,
	}
}

func (runner *SqlRunner) Setup() error {
	for _, service := range runner.services {
		log.Infof("creating service [ %s ] with description [ %s ]", service.Name, service.Description)

		var conn *sql.DB
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create sql connection")
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

func (runner *SqlRunner) Teardown() error {
	for _, service := range runner.services {
		log.Infof("teardown service [ %s ] with description [ %s ]", service.Name, service.Description)

		var conn *sql.DB
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create sql connection")
			}
		}

		if service.Run.Setup != nil {
			for _, setup := range service.Run.Teardown {
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

func (runner *SqlRunner) loadConfiguration(test SqlService) (*manager.DBConfig, error) {
	if test.Configuration != nil {
		return test.Configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid redis configuration")
	}
}

func (runner *SqlRunner) runCommands(conn *sql.DB, run *SqlRun) error {
	for _, command := range run.Queries {
		log.Infof("executing sql command [ %s ]", command)

		if _, err := conn.Exec(command); err != nil {
			return fmt.Errorf("failed to execute sql command [ %s ]", err)
		}
	}
	return nil
}

func (runner *SqlRunner) runCommandsFromFile(conn *sql.DB, run *SqlRun) error {
	for _, file := range run.Files {
		log.Infof("executing nsq commands by file [ %s ]", file)

		var query string
		if bytes, err := ReadFile(file, nil); err != nil {
			return fmt.Errorf("failed to read sql file [ %s ] with error [ %s ]", file, err)
		} else {
			query = string(bytes)
		}

		if _, err := conn.Exec(query); err != nil {
			return fmt.Errorf("failed to execute sql file %s : %s", file, err)
		}
	}
	return nil
}

package gomock

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

func (services *ServicesConfig) setupSQL(defaultConfig *ConfigSQL) error {
	for _, service := range services.SQL {
		fmt.Println(fmt.Sprintf("\n creating service [ %s ] with description [ %s ]", service.Name, service.Description))

		var conn *sql.DB
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create sql connection")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create sql connection")
			}
		} else {
			panic("invalid sql configuration")
		}

		if service.Run.Setup != nil {
			for _, file := range service.Run.Setup.Files {
				fmt.Println(fmt.Sprintf(" executing SQL file [ %s ]", file))

				var query string
				if bytes, err := readFile(file, nil); err != nil {
					return fmt.Errorf("failed to read SQL file [ %s ] with error [ %s ]", file, err)
				} else {
					query = string(bytes)
				}

				if _, err := conn.Exec(query); err != nil {
					return fmt.Errorf("failed to execute SQL file [ %s] with error [ %s ]", file, err)
				}
			}

			for _, query := range service.Run.Setup.Queries {
				fmt.Println(fmt.Sprintf(" executing SQL query [ %s ]", query))

				if _, err := conn.Exec(query); err != nil {
					return fmt.Errorf("failed to execute SQL query [ %s ]", err)
				}
			}
		}
	}
	return nil
}

func (services *ServicesConfig) teardownSQL(defaultConfig *ConfigSQL) error {
	for _, service := range services.SQL {
		fmt.Println(fmt.Sprintf("\n teardown service [ %s ]", service.Name))

		var conn *sql.DB
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create sql connection")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create sql connection")
			}
		} else {
			panic("invalid sql configuration")
		}

		if service.Run.Teardown != nil {
			for _, file := range service.Run.Teardown.Files {
				fmt.Println(fmt.Sprintf(" executing SQL file [ %s ]", file))

				var query string
				if bytes, err := readFile(file, nil); err != nil {
					return fmt.Errorf("failed to read SQL file [ %s ] with error [ %s ]", file, err)
				} else {
					query = string(bytes)
				}

				if _, err := conn.Exec(query); err != nil {
					return fmt.Errorf("failed to execute SQL file %s : %s", file, err)
				}
			}

			for _, query := range service.Run.Teardown.Queries {
				fmt.Println(fmt.Sprintf(" executing SQL query [ %s ]", query))

				if _, err := conn.Exec(query); err != nil {
					return fmt.Errorf("failed to execute SQL query [ %s ]", err)
				}
			}
		}
	}
	return nil
}

// createConnection ...
func (config *ConfigSQL) createConnection() (*sql.DB, error) {
	fmt.Println(fmt.Sprintf(" connecting with driver [ %s ] and data source [ %s ]", config.Driver, config.DataSource))
	return sql.Open(config.Driver, config.DataSource)
}

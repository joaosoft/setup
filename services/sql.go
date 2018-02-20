package gomock

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

func (services *Services) setupSQL() error {
	for _, service := range services.SQL {
		fmt.Println(fmt.Sprintf(" creating service %s", service.Name))

		for _, command := range service.Commands.Setup {
			fmt.Println(fmt.Sprintf(" executing SQL command: %s", command))
			conn, err := service.getConnection()
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed to create connection %s", service.DataSource)
			}
			if _, err := conn.Exec(command); err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed to execute SQL comand %s", err)
			}
		}
	}
	return nil
}

func (services *Services) teardownSQL() error {
	for _, service := range services.SQL {
		fmt.Println(fmt.Sprintf(" teardown service %s", service.Name))

		for _, command := range service.Commands.Teardown {
			fmt.Println(fmt.Sprintf(" executing SQL command: %s", command))
			conn, err := service.getConnection()
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed to create connection %s", service.DataSource)
			}
			if _, err := conn.Exec(command); err != nil {
				return fmt.Errorf("failed to execute SQL comand %s", err)
			}
		}
	}
	return nil
}

// conn ...
func (service *SQL) getConnection() (*sql.DB, error) {
	conn, err := sql.Open(service.Driver, service.DataSource)

	if err != nil {
		return nil, fmt.Errorf("could not instantiate sql connection %s", err)
	}

	return conn, nil
}

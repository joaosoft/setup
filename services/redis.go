package gomock

import (
	"fmt"
	"strings"

	redis "github.com/mediocregopher/radix.v3"
)

func (services *ServicesConfig) setupRedis(defaultConfig *ConfigRedis) error {
	for _, service := range services.Redis {
		fmt.Println(fmt.Sprintf("\n creating service [ %s ] with description [ %s] ", service.Name, service.Description))

		var conn *redis.Pool
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create redis connection")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create redis connection")
			}
		} else {
			panic("invalid redis configuration")
		}

		if service.Run.Setup != nil {
			for _, setup := range service.Run.Setup {
				for _, command := range setup.Commands {
					fmt.Println(fmt.Sprintf(" executing redis command [ %s ] arguments [ %s ]", command.Command, command.Arguments))
					if err := conn.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
						return err
					}
				}

				for _, file := range setup.Files {
					fmt.Println(" executing redis commands...")

					if lines, err := readFileLines(file); err != nil {
						for _, line := range lines {
							command := strings.SplitN(line, " ", 1)
							fmt.Println(fmt.Sprintf(" executing redis command [ %s ] arguments [ %s ]", command[0], command[1]))
							if err := conn.Do(redis.Cmd(nil, command[0], command[1])); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (services *ServicesConfig) teardownRedis(defaultConfig *ConfigRedis) error {
	for _, service := range services.Redis {
		fmt.Println(fmt.Sprintf("\n teardown service [ %s ]", service.Name))

		var conn *redis.Pool
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create redis connection")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create redis connection")
			}
		} else {
			panic("invalid redis configuration")
		}

		if service.Run.Teardown != nil {
			for _, teardown := range service.Run.Teardown {
				for _, command := range teardown.Commands {
					fmt.Println(fmt.Sprintf(" executing redis command [ %s ] arguments [ %s ]", command.Command, command.Arguments))
					if err := conn.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
						return err
					}
				}

				for _, file := range teardown.Files {
					fmt.Println(" executing redis commands...")

					if lines, err := readFileLines(file); err != nil {
						for _, line := range lines {
							command := strings.SplitN(line, " ", 1)
							fmt.Println(fmt.Sprintf(" executing redis command [ %s ] arguments [ %s ]", command[0], command[1]))
							if err := conn.Do(redis.Cmd(nil, command[0], command[1])); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// createConnection ...
func (config *ConfigRedis) createConnection() (*redis.Pool, error) {
	fmt.Println(fmt.Sprintf(" connecting with protocol [ %s ], address [ %s ] and size [ %d ]", config.Protocol, config.Address, config.Size))
	return redis.NewPool(config.Protocol, config.Address, config.Size, nil)
}

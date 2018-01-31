package gomock

import (
	"fmt"

	redis "github.com/mediocregopher/radix.v3"
)

func (services *Services) setupRedis() error {
	for _, service := range services.Redis {
		fmt.Println(fmt.Sprintf(" Creating service %s", service.Name))

		pool, err := redis.NewPool(service.Configuration.Protocol, service.Configuration.Addr, service.Configuration.Size, nil)
		if err != nil {
			return fmt.Errorf("Failed to create redis pool")
		}

		for _, command := range service.Commands.Setup {
			fmt.Println(fmt.Sprintf(" Executing redis command: %s arguments:%s", command.Command, command.Arguments))
			if err := pool.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
				return fmt.Errorf(err.Error())
			}
		}
	}
	return nil
}

func (services *Services) teardownRedis() error {
	for _, service := range services.Redis {
		fmt.Println(fmt.Sprintf(" Teardown service %s", service.Name))

		pool, err := redis.NewPool(service.Configuration.Protocol, service.Configuration.Addr, service.Configuration.Size, nil)
		if err != nil {
			return fmt.Errorf("Failed to create redis pool")
		}

		for _, command := range service.Commands.Teardown {
			fmt.Println(fmt.Sprintf(" Executing redis command: %s arguments:%s", command.Command, command.Arguments))
			if err := pool.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
				return fmt.Errorf(err.Error())
			}
		}
	}
	return nil
}

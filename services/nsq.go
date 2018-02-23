package gomock

import (
	"fmt"
	"time"

	nsqlib "github.com/nsqio/go-nsq"
)

func (services *ServicesConfig) setupNSQ(defaultConfig *ConfigNSQ) error {
	for _, service := range services.NSQ {
		fmt.Println(fmt.Sprintf("\n creating service [ %s ] with description [ %s] ", service.Name, service.Description))

		var conn *nsqlib.Producer
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create nsq pool")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create nsq pool")
			}
		} else {
			panic("invalid nsq configuration")
		}

		for _, command := range service.Run.Setup {
			var message []byte

			if command.Body != nil {
				message = *command.Body
			} else if command.File != nil {
				var err error
				if message, err = readFile(*command.File, nil); err != nil {
					return err
				}
			}

			fmt.Println(fmt.Sprintf(" executing nsq [ %s ] message: %s", command.Description, message))
			if err := conn.Publish(command.Topic, message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (services *ServicesConfig) teardownNSQ(defaultConfig *ConfigNSQ) error {
	for _, service := range services.NSQ {
		fmt.Println(fmt.Sprintf("\n teardown service %s", service.Name))

		var conn *nsqlib.Producer
		var err error
		if service.Configuration != nil {
			if conn, err = service.Configuration.createConnection(); err != nil {
				return fmt.Errorf("failed to create nsq pool")
			}
		} else if defaultConfig != nil {
			if conn, err = defaultConfig.createConnection(); err != nil {
				return fmt.Errorf("failed to create nsq pool")
			}
		} else {
			panic("invalid nsq configuration")
		}

		for _, command := range service.Run.Teardown {
			var message []byte

			if command.Body != nil {
				message = *command.Body
			} else if command.File != nil {
				var err error
				if message, err = readFile(*command.File, nil); err != nil {
					return err
				}
			}

			fmt.Println(fmt.Sprintf(" executing nsq [ %s ] message: %s", command.Description, message))
			if err := conn.Publish(command.Topic, message); err != nil {
				return err
			}
		}
	}
	return nil
}

// createConnection ...
func (config *ConfigNSQ) createConnection() (*nsqlib.Producer, error) {
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	fmt.Println(fmt.Sprintf(" connecting with max attempts [ %d ]", config.MaxAttempts))

	return nsqlib.NewProducer(config.Lookupd, nsqConfig)
}

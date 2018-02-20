package gomock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	nsqlib "github.com/nsqio/go-nsq"
)

func (services *Services) setupNSQ() error {
	for _, service := range services.NSQ {
		fmt.Println(fmt.Sprintf(" creating service %s", service.Name))

		producer, err := services.createProducer(&service.Configuration)
		if err != nil {
			return fmt.Errorf("failed to create nsq connection")
		}

		for _, command := range service.Messages.Setup {
			var message json.RawMessage

			if command.Message != nil {
				message = command.Message
			} else {
				file, err := os.Open(command.File)
				if err != nil {
					return err
				}

				data, err := ioutil.ReadAll(file)
				if err != nil {
					return err
				}

				message = data
			}

			fmt.Println(fmt.Sprintf(" executing nsq [ %s ] message: %s", command.Description, message))
			if err := producer.Publish(command.Topic, message); err != nil {
				return fmt.Errorf(err.Error())
			}
		}
	}
	return nil
}

func (services *Services) teardownNSQ() error {
	for _, service := range services.NSQ {
		fmt.Println(fmt.Sprintf(" creating service %s", service.Name))

		producer, err := services.createProducer(&service.Configuration)
		if err != nil {
			return fmt.Errorf("failed to create nsq connection")
		}

		for _, command := range service.Messages.Teardown {
			var message json.RawMessage

			if command.Message != nil {
				message = command.Message
			} else {
				file, err := os.Open(command.File)
				if err != nil {
					return err
				}

				data, err := ioutil.ReadAll(file)
				if err != nil {
					return err
				}

				message = data
			}

			fmt.Println(fmt.Sprintf(" executing nsq [ %s ] message: %s", command.Description, message))
			if err := producer.Publish(command.Topic, message); err != nil {
				return fmt.Errorf(err.Error())
			}
		}
	}
	return nil
}

// NSQConfig ...
type NSQConfig struct {
	Lookupd      string `json:"lookupd"`
	RequeueDelay int64  `json:"requeue_delay"`
	MaxInFlight  int    `json:"max_in_flight"`
	MaxAttempts  uint16 `json:"max_attempts"`
}

func (services *Services) createProducer(config *NSQConfig) (*nsqlib.Producer, error) {
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	return nsqlib.NewProducer(config.Lookupd, nsqConfig)
}

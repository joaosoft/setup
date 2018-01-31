package gomock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (service *Services) fromFile(fileName string) error {
	var err error

	fmt.Println(fmt.Sprintf(":: Loading file %s", fileName))

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, service); err != nil {
		return err
	}

	return nil
}

func (config *Services) setup() error {
	// web services
	if err := config.setupWebServices(); err != nil {
		return err
	}

	// redis
	if err := config.setupRedis(); err != nil {
		return err
	}

	// sql
	if err := config.setupSQL(); err != nil {
		return err
	}

	return nil
}

func (config *Services) teardown() error {
	// web services
	if err := config.teardownWebServices(); err != nil {
		return err
	}

	// redis
	if err := config.teardownRedis(); err != nil {
		return err
	}

	// sql
	if err := config.teardownSQL(); err != nil {
		return err
	}

	return nil
}

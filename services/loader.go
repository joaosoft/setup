package gomock

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (service *ServicesConfig) fromFile(fileName string) error {
	service.File = fileName
	var err error

	if !exists(fileName) {
		fileName = global["path"] + fileName
	}

	fmt.Println(fmt.Sprintf(":: Loading file [ %s ]", fileName))

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

func exists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readFile(fileName string, obj interface{}) ([]byte, error) {
	var err error

	if !exists(fileName) {
		fileName = global["path"] + fileName
	}

	fmt.Println(fmt.Sprintf(":: Loading file [ %s ]", fileName))
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if obj != nil {
		fmt.Println(fmt.Sprintf("Unmarshalling file [ %s ] to struct", fileName))
		if err := json.Unmarshal(data, obj); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func readFileLines(fileName string) ([]string, error) {
	lines := make([]string, 0)

	if !exists(fileName) {
		fileName = global["path"] + fileName
	}

	fmt.Println(fmt.Sprintf(":: Loading file [ %s ]", fileName))
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (config *ServicesConfig) setup(defaults map[string]interface{}) error {
	// web services
	if err := config.setupWebServices(); err != nil {
		return err
	}

	// redis
	if err := config.setupRedis(getDefaultConfigRedis(defaults)); err != nil {
		return err
	}

	// nsq
	if err := config.setupNSQ(getDefaultConfigNSQ(defaults)); err != nil {
		return err
	}

	// sql
	if err := config.setupSQL(getDefaultConfigSQL(defaults)); err != nil {
		return err
	}

	return nil
}

func (config *ServicesConfig) teardown(defaults map[string]interface{}) error {
	// web services
	if err := config.teardownWebServices(); err != nil {
		return err
	}

	// redis
	if err := config.teardownRedis(getDefaultConfigRedis(defaults)); err != nil {
		return err
	}

	// nsq
	if err := config.teardownNSQ(getDefaultConfigNSQ(defaults)); err != nil {
		return err
	}

	// sql
	if err := config.teardownSQL(getDefaultConfigSQL(defaults)); err != nil {
		return err
	}

	return nil
}

func getDefaultConfigNSQ(defaults map[string]interface{}) *ConfigNSQ {
	if value, exists := defaults["nsq"]; exists {
		return value.(*ConfigNSQ)
	}
	return nil
}

func getDefaultConfigSQL(defaults map[string]interface{}) *ConfigSQL {
	if value, exists := defaults["sql"]; exists {
		return value.(*ConfigSQL)
	}
	return nil
}

func getDefaultConfigRedis(defaults map[string]interface{}) *ConfigRedis {
	if value, exists := defaults["redis"]; exists {
		return value.(*ConfigRedis)
	}
	return nil
}

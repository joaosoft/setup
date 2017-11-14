package gomock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

// GoMock ...
type GoMock struct {
	processes map[string]*echo.Echo
}

// NewGoMock ...
func NewGoMock() *GoMock {
	return &GoMock{
		processes: make(map[string]*echo.Echo),
	}
}

// Handle ...
func (instance Response) Handle(ctx echo.Context) error {
	data, _ := json.Marshal(instance.Body)
	fmt.Println("RESPONSE: " + string(data))

	return ctx.JSON(instance.Status, instance.Body)
}

// Run ...
func (gomock *GoMock) Run(path string) error {

	fmt.Println(":: Initializing Mock Service")

	if err := gomock.setup(path); err != nil {
		log.Panic(err)
		return err
	}

	fmt.Println(":: Mock Services Started")

	return nil
}

// Stop ...
func (gomock *GoMock) Stop() error {

	fmt.Println(":: Shutting down Mock Service")

	if err := gomock.teardown(); err != nil {
		log.Panic(err)
		return err
	}

	fmt.Println(":: Shutted down Mock Service")

	return nil
}

func (gomock *GoMock) setup(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
			if err := gomock.loadFile(path); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (gomock *GoMock) loadFile(fileName string) error {
	var err error
	config := &ServicesLoader{}

	fmt.Println(fmt.Sprintf("Loading file %s", fileName))

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	for _, service := range config.WebServices {
		fmt.Println(fmt.Sprintf("Creating service %s", service.Name))

		e := echo.New()
		e.HideBanner = true
		for _, route := range service.Routes {
			fmt.Println(fmt.Sprintf("Creating route %s", route.Route))

			e.Add(route.Method, route.Route, route.Response.Handle)
		}

		go e.Start(service.Host)

		key := "webservice::" + service.Name
		fmt.Println(fmt.Sprintf(":: Started service: %s at %s", key, service.Host))

		gomock.processes[key] = e
	}

	return nil
}

func (gomock *GoMock) teardown() error {

	for key, process := range gomock.processes {
		process.Close()

		fmt.Println(fmt.Sprintf(":: TearDown %s", key))
	}

	return nil
}

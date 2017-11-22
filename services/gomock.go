package gomock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/labstack/echo"
	redis "github.com/mediocregopher/radix.v3"
)

const (
	path       = "."
	background = true
)

// GoMockOption ...
type GoMockOption func(gomock *GoMock)

// WithPath ...
func WithPath(path string) GoMockOption {
	return func(gomock *GoMock) {
		if path != "" {
			gomock.path = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) GoMockOption {
	return func(gomock *GoMock) {
		gomock.background = background
	}
}

// WithMaxRetries ...
func WithMaxRetries(maxRetries int) GoMockOption {
	return func(gomock *GoMock) {
		gomock.maxRetries = maxRetries
	}
}

// Reconfigure ...
func (gomock *GoMock) Reconfigure(options ...GoMockOption) {
	for _, option := range options {
		option(gomock)
	}
}

// GoMock ...
type GoMock struct {
	path       string
	background bool
	maxRetries int
	processes  map[string]*echo.Echo
}

// NewGoMock ...
func NewGoMock(options ...GoMockOption) *GoMock {
	gomock := &GoMock{
		path:       path,
		background: background,
		processes:  make(map[string]*echo.Echo),
	}

	gomock.Reconfigure(options...)

	return gomock
}

// Handle ...
func (instance Response) Handle(ctx echo.Context) error {
	data, _ := json.Marshal(instance.Body)
	fmt.Println("RESPONSE: " + string(data))

	return ctx.JSON(instance.Status, instance.Body)
}

// Run ...
func (gomock *GoMock) Run() error {

	fmt.Println(":: Initializing Mock Service")

	if err := gomock.setup(); err != nil {
		log.Panic(err)
		return err
	}

	if !gomock.background {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		gomock.Stop()
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

func (gomock *GoMock) setup() error {
	err := filepath.Walk(gomock.path, func(path string, f os.FileInfo, err error) error {
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

	fmt.Println(fmt.Sprintf(":: Loading file %s", fileName))

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

	// Web Services
	for _, service := range config.WebServices {
		fmt.Println(fmt.Sprintf(":: Creating service %s", service.Name))

		e := echo.New()
		e.HideBanner = true
		for _, route := range service.Routes {
			fmt.Println(fmt.Sprintf(":: Creating route %s", route.Route))

			e.Add(route.Method, route.Route, route.Response.Handle)
		}

		go e.Start(service.Host)

		key := "webservice::" + service.Name
		fmt.Println(fmt.Sprintf(":: Started service: %s at %s", key, service.Host))

		gomock.processes[key] = e
	}

	// Redis
	for _, service := range config.Redis {
		fmt.Println(fmt.Sprintf(":: Creating service %s", service.Name))

		pool, err := redis.NewPool(service.Configuration.Protocol, service.Configuration.Addr, service.Configuration.Size, nil)
		if err != nil {
			return fmt.Errorf("Failed to create redis pool")
		}
		for _, command := range service.Commands {
			fmt.Println(fmt.Sprintf(":: Executing redis command: %s arguments:%s", command.Command, command.Arguments))
			if err := pool.Do(redis.Cmd(nil, command.Command, command.Arguments...)); err != nil {
				return fmt.Errorf(err.Error())
			}
		}
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

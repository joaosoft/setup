package gomock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	expandedMatchers "github.com/Benjamintf1/Expanded-Unmarshalled-Matchers"
	gomega "github.com/onsi/gomega"
)

var processes = make(map[string]*echo.Echo)

type HttpRunner struct {
	services []HttpService
}

func NewHttpRunner(services []HttpService) *HttpRunner {
	return &HttpRunner{
		services: services,
	}
}

func (runner *HttpRunner) Setup() error {
	for _, service := range runner.services {
		log.Infof("creating service [ %s ] with description [ %s ]", service.Name, service.Description)

		e := echo.New()
		e.HideBanner = true

		if err := runner.runRoutes(e, &service); err != nil {
			return fmt.Errorf("error adding service routes [service: %s]", service.Name)
		}

		// shutdown service on allocated port
		//if listener, err := net.Listen("tcp", service.Host); err != nil {
		//	log.Info(err)
		//	log.Infof("closing connection to %s", service.Host)
		//	listener.Close()
		//}

		go e.Start(service.Host)

		key := "http" + service.Name
		log.Infof("started service [ %s ] at [ %s ]", service.Name, service.Host)

		processes[key] = e
	}

	return nil
}

func (runner *HttpRunner) Teardown() error {
	for _, service := range runner.services {
		log.Infof("teardown service [ %s ]", service.Name)
		key := "http" + service.Name
		processes[key].Close()
	}

	return nil
}

func failHandler(message string, callerSkip ...int) {
	log.Infof("failed with message [ %s ]", message)
}

// Handle ...
func (instance Route) handle(ctx echo.Context) error {
	gomega.RegisterFailHandler(failHandler)

	log.Infof("calling [ %s ] URL [ %s ]", ctx.Request().Method, ctx.Request().URL)

	var requestBody json.RawMessage
	ctx.Bind(&requestBody)

	// what to expect
	var expectedBody string
	if instance.Body != nil {
		expectedBody = string(instance.Body)
	} else if instance.File != nil {
		if bytes, err := readFile(*instance.File, nil); err != nil {
			return err
		} else {
			expectedBody = string(bytes)
		}
	}
	if instance.Body != nil || instance.File != nil {
		if gomega.Expect(string(requestBody)).To(expandedMatchers.MatchUnorderedJSON(string(expectedBody))) {
		} else {
			log.Infof("expect [ %s ] to be equal to [ %s ]", string(requestBody), expectedBody)
			return ctx.NoContent(http.StatusNotFound)
		}
	}

	// what to return
	var response json.RawMessage
	if instance.Response.Body != nil {
		response = instance.Response.Body
	} else if instance.Response.File != nil {
		if bytes, err := readFile(*instance.Response.File, nil); err != nil {
			return err
		} else {
			response = bytes
		}
	} else {
		log.Info("there is no body to process")
	}

	log.Infof("response [ %s ]", string(response))

	return ctx.JSON(instance.Response.Status, response)
}

func (runner *HttpRunner) runRoutes(e *echo.Echo, run *HttpService) error {
	for _, route := range run.Routes {
		log.Infof("creating route [ %s ] method [ %s ]", route.Route, route.Method)

		e.Add(route.Method, route.Route, route.handle)
	}
	return nil
}

package setup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/joaosoft/logger"

	"github.com/labstack/echo"

	expandedMatchers "github.com/Benjamintf1/Expanded-Unmarshalled-Matchers"
	gomega "github.com/onsi/gomega"
)

var processes = make(map[string]*echo.Echo)

type HttpRunner struct {
	services []HttpService
	logger   logger.ILogger
}

func (setup *Runner) NewHttpRunner(services []HttpService) *HttpRunner {
	return &HttpRunner{
		services: services,
		logger: setup.logger,
	}
}

func (runner *HttpRunner) Setup() error {
	for _, service := range runner.services {
		runner.logger.Infof("creating service [ %s ] with description [ %s ]", service.Name, service.Description)

		e := echo.New()
		e.HideBanner = true

		if err := runner.runRoutes(e, &service); err != nil {
			return fmt.Errorf("error adding service routes [service: %s]", service.Name)
		}

		// shutdown service on allocated port
		//if listener, err := net.Listen("tcp", service.Host); err != nil {
		//	runner.logger.Infof("closing connection to %s", service.Host)
		//	listener.Close()
		//}

		go e.Start(service.Host)

		key := "http" + service.Name
		runner.logger.Infof("started service [ %s ] at [ %s ]", service.Name, service.Host)

		processes[key] = e
	}

	return nil
}

func (runner *HttpRunner) Teardown() error {
	for _, service := range runner.services {
		runner.logger.Infof("teardown service [ %s ]", service.Name)
		key := "http" + service.Name
		processes[key].Close()
	}

	return nil
}

func (route Route) failHandler(message string, callerSkip ...int) {
	route.logger.Infof("failed with message [ %s ]", message)
}

// Handle ...
func (route Route) handle(ctx echo.Context) error {
	gomega.RegisterFailHandler(route.failHandler)

	route.logger.Infof("calling [ %s ] URL [ %s ]", ctx.Request().Method, ctx.Request().URL)

	var requestBody json.RawMessage
	ctx.Bind(&requestBody)

	// headers
	if route.Headers != nil {
		for expectedKey, expectedValue := range *route.Headers {
			if requestValue, ok := ctx.Request().Header[expectedKey]; !ok {
				return fmt.Errorf("the header [ %s: %s ] is not defined in the request", expectedKey, expectedValue)
			} else if !reflect.DeepEqual(expectedValue, requestValue) {
				return fmt.Errorf("the headers aren't the ones we expected [ key: %s, request: %+v, expected: %+v ]", expectedKey, requestValue, expectedValue)
			}
		}
	}

	// cookies
	if len(route.Cookies) > 0 {
		for _, expectedCookie := range route.Cookies {
			found := false
			for _, requestCookie := range ctx.Cookies() {
				if expectedCookie != nil && requestCookie != nil {
					if expectedCookie.Name != nil && *expectedCookie.Name == requestCookie.Name {
						found = true
						if expectedCookie.Value != nil && *expectedCookie.Value != requestCookie.Value ||
							expectedCookie.Domain != nil && *expectedCookie.Domain != requestCookie.Domain ||
							expectedCookie.Path != nil && *expectedCookie.Path != requestCookie.Path ||
							expectedCookie.Expires != nil && *expectedCookie.Expires != requestCookie.Expires {

							return fmt.Errorf("the cookie is diferent that we expected!git sa\n"+
								"actual: [ name: %s, value: %s, domain: %s, path: %s, expires: %s ]\n"+
								"expected: [ name: %s, value: %s, domain: %s, path: %s, expires: %s ]"+
								requestCookie.Name, requestCookie.Value, requestCookie.Domain, requestCookie.Path, requestCookie.Expires,
								*expectedCookie.Name, *expectedCookie.Value, *expectedCookie.Domain, *expectedCookie.Path, *expectedCookie.Expires)
						}
					}
				}
			}
			if !found {
				return fmt.Errorf("the cookie isn't in the request [ name: %s value: %s ]", *expectedCookie.Name, *expectedCookie.Value)
			}
		}
	}

	// what to expect
	var expectedBody string
	if route.Body != nil {
		expectedBody = string(route.Body)
	} else if route.File != nil {
		if bytes, err := ReadFile(*route.File, nil); err != nil {
			return err
		} else {
			expectedBody = string(bytes)
		}
	}
	if route.Body != nil || route.File != nil {
		if gomega.Expect(string(requestBody)).To(expandedMatchers.MatchUnorderedJSON(string(expectedBody))) {
		} else {
			route.logger.Infof("expect [ %s ] to be equal to [ %s ]", string(requestBody), expectedBody)
			return ctx.NoContent(http.StatusNotFound)
		}
	}

	// what to return
	var response json.RawMessage
	if route.Response.Body != nil {
		response = route.Response.Body
	} else if route.Response.File != nil {
		if bytes, err := ReadFile(*route.Response.File, nil); err != nil {
			return err
		} else {
			response = bytes
		}
	} else {
		route.logger.Info("there is no body to process")
	}

	route.logger.Infof("response [ %s ]", string(response))

	return ctx.JSON(route.Response.Status, response)
}

func (runner *HttpRunner) runRoutes(e *echo.Echo, run *HttpService) error {
	for _, route := range run.Routes {
		runner.logger.Infof("creating route [ %s ] method [ %s ]", route.Route, route.Method)

		e.Add(route.Method, route.Route, route.handle)
	}
	return nil
}

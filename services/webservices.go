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

func (services *ServicesConfig) setupWebServices() error {
	// Web Services
	for _, service := range services.WebServices {
		fmt.Println(fmt.Sprintf("\n creating service [ %s ] with description [ %s ]", service.Name, service.Description))

		e := echo.New()
		e.HideBanner = true
		for _, route := range service.Routes {
			fmt.Println(fmt.Sprintf(" creating route [ %s ] method [ %s ]", route.Route, route.Method))

			e.Add(route.Method, route.Route, route.handle)
		}

		go e.Start(service.Host)

		key := "webservice" + service.Name
		fmt.Println(fmt.Sprintf(" started service [ %s ] at [ %s ]", service.Name, service.Host))

		processes[key] = e

	}
	return nil
}

func (services *ServicesConfig) teardownWebServices() error {
	for _, service := range services.WebServices {
		fmt.Println(fmt.Sprintf("\n teardown service [ %s ]", service.Name))
		key := "webservice" + service.Name
		processes[key].Close()

	}
	return nil
}

func failHandler(message string, callerSkip ...int) {
	fmt.Println(fmt.Sprintf("failed with message [ %s ]", message))
}

// Handle ...
func (instance Route) handle(ctx echo.Context) error {
	gomega.RegisterFailHandler(failHandler)

	fmt.Print(fmt.Sprintf(" calling [ %s ] URL [ %s ]", ctx.Request().Method, ctx.Request().URL))

	var requestBody json.RawMessage
	ctx.Bind(&requestBody)

	// what to expect
	var expectedBody string
	if instance.Body != nil {
		expectedBody = string(instance.Body)
	} else if instance.File != nil {
		if bytes, err := readFile(*instance.File, nil); err != nil {
			expectedBody = string(bytes)
		}
	}
	if (instance.Body != nil || instance.File != nil) &&
		gomega.Expect(string(requestBody)).ToNot(expandedMatchers.MatchUnorderedJSON(string(expectedBody))) {
		fmt.Println(" with invalid payload")
		fmt.Println(fmt.Sprint(" expect [ %s ] to be equal to [ %s ]", string(requestBody), expectedBody))
		return ctx.NoContent(http.StatusNotFound)
	}
	fmt.Println(" with valid payload")

	// what to return
	var response string
	if instance.Response.Body != nil {
		response = string(instance.Body)
	} else if instance.Response.File != nil {
		if bytes, err := readFile(*instance.Response.File, nil); err != nil {
			return err
		} else {
			response = string(bytes)
		}
	} else {
		fmt.Println(" there is no body to process")
	}

	data, _ := json.Marshal(response)
	fmt.Println(fmt.Sprintf(" response [ %s ]", string(data)))

	return ctx.JSON(instance.Response.Status, instance.Response.Body)
}

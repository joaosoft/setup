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

func (services *Services) setupWebServices() error {
	// Web Services
	for _, service := range services.WebServices {
		fmt.Println(fmt.Sprintf(" Creating service %s", service.Name))

		e := echo.New()
		e.HideBanner = true
		for _, route := range service.Routes {
			fmt.Println(fmt.Sprintf(" Creating route %s method %s", route.Route, route.Method))

			e.Add(route.Method, route.Route, route.handle)
		}

		go e.Start(service.Host)

		key := "webservice" + service.Name
		fmt.Println(fmt.Sprintf(" Started service: %s at %s", service.Name, service.Host))

		processes[key] = e

	}
	return nil
}

func (services *Services) teardownWebServices() error {
	for _, service := range services.WebServices {
		fmt.Println(fmt.Sprintf(" Teardown service %s", service.Name))
		key := "webservice" + service.Name
		processes[key].Close()

	}
	return nil
}

func failHandler(message string, callerSkip ...int) {
	fmt.Println(fmt.Sprintf("Failed %s", message))
}

// Handle ...
func (instance Route) handle(ctx echo.Context) error {
	gomega.RegisterFailHandler(failHandler)

	fmt.Print(fmt.Sprintf("Calling POST URL: %s", ctx.Request().URL))

	var body json.RawMessage
	ctx.Bind(&body)

	fmt.Println(fmt.Sprintf("\n %s \n %s", string(body), string(instance.Payload)))

	if instance.Payload != nil &&
		gomega.Expect(string(body)).ToNot(expandedMatchers.MatchUnorderedJSON(string(instance.Payload))) {
		fmt.Println(" with invalid payload")

		return ctx.NoContent(http.StatusNotFound)
	}
	fmt.Println(" with valid payload")

	data, _ := json.Marshal(instance.Response.Body)
	fmt.Println("RESPONSE: " + string(data))

	return ctx.JSON(instance.Response.Status, instance.Response.Body)
}

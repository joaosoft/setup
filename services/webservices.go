package gomock

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
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

// Handle ...
func (instance Route) handle(ctx echo.Context) error {
	//fmt.Print(fmt.Sprintf("Calling POST URL: %s", ctx.Request().URL))
	//fmt.Println(string(instance.Payload))
	//aa, _ := json.Marshal(ctx.Request().Body)
	//fmt.Println(string(aa))
	//if instance.Payload != nil &&
	//	gomega.Expect(ctx.Request().Body).ToNot(gomega.Equal(instance.Payload)) {
	//	fmt.Println(" with invalid payload")
	//
	//	return ctx.NoContent(http.StatusNotFound)
	//}
	//fmt.Println(" with valid payload")

	data, _ := json.Marshal(instance.Response.Body)
	fmt.Println("RESPONSE: " + string(data))

	return ctx.JSON(instance.Response.Status, instance.Response.Body)
}

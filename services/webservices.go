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
			fmt.Println(fmt.Sprintf(" Creating route %s", route.Route))

			e.Add(route.Method, route.Route, route.Response.handle)
		}

		go e.Start(service.Host)

		key := "webservice" + service.Name
		fmt.Println(fmt.Sprintf(" Started service: %s at %s", key, service.Host))

		processes[key] = e

	}
	return nil
}

func (services *Services) teardownWebServices() error {
	for _, service := range services.WebServices {
		fmt.Println(fmt.Sprintf(" Teardown service %s", service.Name))

		key := "webservice" + service.Name
		processes[key].Close()
		fmt.Println(fmt.Sprintf(" TearDown %s", key))

	}
	return nil
}

// Handle ...
func (instance Response) handle(ctx echo.Context) error {
	data, _ := json.Marshal(instance.Body)
	fmt.Println("RESPONSE: " + string(data))

	return ctx.JSON(instance.Status, instance.Body)
}

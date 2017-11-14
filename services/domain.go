package gomock

type ServicesLoader struct {
	WebServices []WebService `json:"webservices,omitempty"`
}

type WebService struct {
	Name   string  `json:"name"`
	Host   string  `json:"host"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Description string   `json:"description"`
	Route       string   `json"route"`
	Method      string   `json:"method"`
	Response    Response `json:"response"`
}

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

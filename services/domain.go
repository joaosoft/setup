package gomock

type ServicesLoader struct {
	WebServices []WebService `json:"webservices,omitempty"`
	Redis       []Redis      `json:"redis,omitempty"`
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

type Redis struct {
	Name          string        `json:"name"`
	Configuration Configuration `json:"configuration"`
	Commands      []Command     `json:"commands"`
}

type Configuration struct {
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`
	Size     int    `json:"size"`
}

type Command struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

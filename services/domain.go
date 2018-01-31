package gomock

// Services
type Services struct {
	WebServices []WebService `json:"webservices,omitempty"`
	Redis       []Redis      `json:"redis,omitempty"`
	SQL         []SQL        `json:"sql,omitempty"`
}

// WebService
type WebService struct {
	Name   string  `json:"name"`
	Host   string  `json:"host"`
	Routes []Route `json:"routes"`
}

// Route
type Route struct {
	Description string   `json:"description"`
	Route       string   `json"route"`
	Method      string   `json:"method"`
	Response    Response `json:"response"`
}

// Response
type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

// Redis
type Redis struct {
	Name          string        `json:"name"`
	Configuration Configuration `json:"configuration"`
	Commands      struct {
		Setup    []Command `json:"setup"`
		Teardown []Command `json:"teardown"`
	} `json:"commands"`
}

// Configuration
type Configuration struct {
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`
	Size     int    `json:"size"`
}

// Command
type Command struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// SQL
type SQL struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	DataSource  string `json:"datasource"`
	Commands    struct {
		Setup    []string `json:"setup"`
		Teardown []string `json:"teardown"`
	} `json:"commands"`
}

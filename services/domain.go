package gomock

import "encoding/json"

// Services
type Services struct {
	WebServices []WebService `json:"webservices,omitempty"`
	Redis       []Redis      `json:"redis,omitempty"`
	NSQ         []NSQ        `json:"nsq,omitempty"`
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
	Description string          `json:"description"`
	Route       string          `json"route"`
	Method      string          `json:"method"`
	Payload     json.RawMessage `json:"payload"`
	Response    Response        `json:"response"`
}

// Response
type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

// Redis
type Redis struct {
	Name          string      `json:"name"`
	Configuration RedisConfig `json:"configuration"`
	Commands      struct {
		Setup    []RedisCommand `json:"setup"`
		Teardown []RedisCommand `json:"teardown"`
	} `json:"commands"`
}

// NSQ
type NSQ struct {
	Name          string    `json:"name"`
	Configuration NSQConfig `json:"configuration"`
	Messages      struct {
		Setup    []NSQMessage `json:"setup"`
		Teardown []NSQMessage `json:"teardown"`
	} `json:"messages"`
}

// RedisCommand
type RedisCommand struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// NSQMessage
type NSQMessage struct {
	Description string          `json:"description"`
	Topic       string          `json:"topic"`
	Message     json.RawMessage `json:"message"`
	File        string          `json:"file"`
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

package gomock

import (
	"encoding/json"
)

// Services
type Services struct {
	HttpServices  []HttpService  `json:"http,omitempty"`
	RedisServices []RedisService `json:"redis,omitempty"`
	NSQServices   []NSQService   `json:"nsq,omitempty"`
	SQLServices   []SQLService   `json:"sql,omitempty"`
}

// HttpService
type HttpService struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Host        string  `json:"host"`
	Routes      []Route `json:"routes"`
}

// Route
type Route struct {
	Description string          `json:"description"`
	Route       string          `json"route"`
	Method      string          `json:"method"`
	Body        json.RawMessage `json:"body"`
	File        *string         `json:"file"`
	Response    Response        `json:"response"`
}

// Response
type Response struct {
	Status int             `json:"status"`
	Body   json.RawMessage `json:"body"`
	File   *string         `json:"file"`
}

// RedisService
type RedisService struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Configuration *RedisConfig `json:"configuration"`
	Connection    *string      `json:"connection"`
	Run           struct {
		Setup    []RedisRun `json:"setup"`
		Teardown []RedisRun `json:"teardown"`
	} `json:"run"`
}

// RedisRun
type RedisRun struct {
	Commands []RedisCommand `json:"commands"`
	Files    []string       `json:"files"`
}

type RedisCommand struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// NSQService
type NSQService struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *NSQConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    []NSQRun `json:"setup"`
		Teardown []NSQRun `json:"teardown"`
	} `json:"run"`
}

// NSQRun
type NSQRun struct {
	Description string          `json:"description"`
	Topic       string          `json:"topic"`
	Message     json.RawMessage `json:"message"`
	File        string          `json:"file"`
}

// SQLService
type SQLService struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *SQLConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    []SQLRun `json:"setup"`
		Teardown []SQLRun `json:"teardown"`
	} `json:"run"`
}

type SQLRun struct {
	Files   []string `json:"file"`
	Queries []string `json:"query"`
}

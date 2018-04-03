package gosetup

import (
	"encoding/json"
	"time"
)

// Services
type Services struct {
	Files         []string       `json:"files,omitempty"`
	HttpServices  []HttpService  `json:"http,omitempty"`
	RedisServices []RedisService `json:"redis,omitempty"`
	NsqServices   []NsqService   `json:"nsq,omitempty"`
	SqlServices   []SqlService   `json:"sql,omitempty"`
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
	Headers     *Headers        `json:"headers"`
	Cookies     []*Cookies      `json:"cookies"`
	Body        json.RawMessage `json:"body"`
	File        *string         `json:"file"`
	Response    Response        `json:"response"`
}

// Headers
type Headers map[string][]string

// Cookies
type Cookies struct {
	Name    *string    `json:"name"`
	Value   *string    `json:"value"`
	Path    *string    `json:"path"`    // optional
	Domain  *string    `json:"domain"`  // optional
	Expires *time.Time `json:"expires"` // optional
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

// NsqService
type NsqService struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *NsqConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    []NsqRun `json:"setup"`
		Teardown []NsqRun `json:"teardown"`
	} `json:"run"`
}

// NsqRun
type NsqRun struct {
	Description string          `json:"description"`
	Topic       string          `json:"topic"`
	Message     json.RawMessage `json:"message"`
	File        string          `json:"file"`
}

// SqlService
type SqlService struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *SqlConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    []SqlRun `json:"setup"`
		Teardown []SqlRun `json:"teardown"`
	} `json:"run"`
}

type SqlRun struct {
	Files   []string `json:"file"`
	Queries []string `json:"query"`
}

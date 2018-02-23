package gomock

import "encoding/json"

// ServicesConfig
type ServicesConfig struct {
	File        string       `json:"file,omitempty"`
	WebServices []WebService `json:"webservices,omitempty"`
	Redis       []Redis      `json:"redis,omitempty"`
	NSQ         []NSQ        `json:"nsq,omitempty"`
	SQL         []SQL        `json:"sql,omitempty"`
}

// WebService
type WebService struct {
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

// Redis
type Redis struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Configuration *ConfigRedis `json:"configuration"`
	Connection    *string      `json:"connection"`
	Run           struct {
		Setup    []RedisCommand `json:"setup"`
		Teardown []RedisCommand `json:"teardown"`
	} `json:"run"`
}

// RedisCommand
type RedisCommand struct {
	Commands []struct {
		Command   string   `json:"command"`
		Arguments []string `json:"arguments"`
	} `json:"commands"`
	Files []string `json:"files"`
}

// NSQ
type NSQ struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *ConfigNSQ `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    []NSQMessage `json:"setup"`
		Teardown []NSQMessage `json:"teardown"`
	} `json:"run"`
}

// NSQMessage
type NSQMessage struct {
	Description string           `json:"description"`
	Topic       string           `json:"topic"`
	Body        *json.RawMessage `json:"body"`
	File        *string          `json:"file"`
}

// SQL
type SQL struct {
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Configuration *ConfigSQL `json:"configuration"`
	Connection    *string    `json:"connection"`
	Run           struct {
		Setup    *SQLData `json:"setup"`
		Teardown *SQLData `json:"teardown"`
	} `json:"run"`
}

type SQLData struct {
	Files   []string `json:"files"`
	Queries []string `json:"queries"`
}

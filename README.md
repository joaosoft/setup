# go-mock
[![Build Status](https://travis-ci.org/joaosoft/go-mock.svg?branch=master)](https://travis-ci.org/joaosoft/go-mock) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-mock/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-mock)

A package framework to create mock services. At the moment it has support for web services, redis, postgres, mysql and nsq services. 
## Dependency Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/go-mock
```

## Docker
>### Start Environment 
* Redis / Postgres / MySQL / NSQ
```
make env
```

>### Start Application
```
make start
```

## Usage 
This example is available in the project at [go-mock/bin/launcher](https://github.com/joaosoft/go-mock/tree/master/bin/launcher)

```go
package main

import (
	gomock "github.com/joaosoft/go-mock/services"
)

func main() {
	// WEB SERVICES
	test := gomock.NewGoMock(
		gomock.WithPath("./config"))
	test.RunSingleNoWait("001_webservices.json")

	// SQL
	configSQL := &gomock.ConfigSQL{
		Driver:     "postgres",
		DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
	}
	test.Reconfigure(gomock.WithConfigurationSQL(configSQL))
	test.RunSingleNoWait("002_sql.json")

	// NSQ
	configNSQ := &gomock.ConfigNSQ{
		Lookupd:      "localhost:4150",
		RequeueDelay: 30,
		MaxInFlight:  5,
		MaxAttempts:  5,
	}
	test.Reconfigure(gomock.WithConfigurationNSQ(configNSQ))
	test.RunSingleNoWait("003_nsq.json")

	// REDIS
	configRedis := &gomock.ConfigRedis{
		Protocol: "tcp",
		Address:  "localhost:6379",
		Size:     10,
	}
	test.Reconfigure(gomock.WithConfigurationRedis(configRedis))
	test.RunSingleNoWait("004_redis.json")

	// ALL
	test.Reconfigure(
		gomock.WithRunInBackground(false),
		gomock.WithConfiguration("data/app.json"))
	test.Run()
}
```

>## Configurations

>### WebServices [ see 001_webservices.json ]

```javascript
{
  "webservices": [
    {
      "name": "hello",
      "description": "test hello",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/hello",
          "response": {
            "status": 200,
            "body": {
              "message": "Hello friend!"
            }
          }
        }
      ]
    },
    {
      "name": "something",
      "description": "testing payload of a post",
      "host": ":8003",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/something",
          "body": {
            "name": "joao",
            "age": 29
          },
          "response": {
            "status": 200,
            "body": {
              "message": "Goodbye friend!"
            }
          }
        }
      ]
    },
    {
      "name": "loading",
      "description": "loading the payload from a file",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/loading",
          "response": {
            "status": 200,
            "file": "data/webservice_body.json"
          }
        }
      ]
    }
  ]
}

```

>### SQL [ see 002_sql.json ]
```javascript
{
  "sql": [
    {
      "name": "postgres",
      "description": "add users information",
      "configuration": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7001?sslmode=disable"
      },
      "run": {
        "setup": {
          "queries": [
            "DROP TABLE IF EXISTS USERS",
            "CREATE TABLE USERS(name varchar(255), description varchar(255))",
            "INSERT INTO USERS(name, description) VALUES('joao', 'administrator')",
            "INSERT INTO USERS(name, description) VALUES('tiago', 'user')"
          ]
        },
        "teardown": {
          "queries": [
            "DROP TABLE IF EXISTS USERS"
          ]
        }
      }
    },
    {
      "name": "postgres",
      "description": "add users information from files",
      "run": {
        "setup": {
          "files": ["data/sql_setup_file.sql"]
        },
        "teardown": {
          "files": ["data/sql_teardown_file.sql"]
        }
      }
    },
    {
      "name": "mysql",
      "description": "add clients information",
      "configuration": {
        "driver": "mysql",
        "datasource": "root:password@tcp(127.0.0.1:7002)/mysql"
      },
      "run": {
        "setup": {
          "queries": [
            "DROP TABLE IF EXISTS CLIENTS",
            "CREATE TABLE CLIENTS(name varchar(255), description varchar(255))",
            "INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator')",
            "INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user')"
          ]
        },
        "teardown": {
          "queries": [
            "DROP TABLE IF EXISTS CLIENTS"
          ]
        }
      }
    }
  ]
}
```

>### NSQ [ see 003_nsq.json ]
```javascript
{
  "nsq": [
    {
      "name": "nsq",
      "description": "loading a script from file and from body",
      "configuration": {
        "lookupd": "localhost:4150",
        "requeue_delay": 30,
        "max_in_flight": 5,
        "max_attempts": 5
      },
      "run": {
        "setup": [
          {
            "description": "ADD PERSON ONE",
            "topic": "topic.example.lo",
            "body": {
              "name": "joao",
              "age": 29
            }
          },
          {
            "description": "ADD PERSON ONE",
            "topic": "topic.example.hi",
            "file": "data/xml_file.txt"
          }
        ],
        "teardown": []
      }
    },
    {
      "name": "nsq",
      "description": "",
      "configuration": {
        "lookupd": "localhost:4150",
        "requeue_delay": 30,
        "max_in_flight": 5,
        "max_attempts": 5
      },
      "run": {
        "setup": [
          {
            "description": "ADD PERSON TWO",
            "topic": "topic.example.lo",
            "body": {
              "name": "pedro",
              "age": 30
            }
          },
          {
            "description": "ADD PERSON TWO",
            "topic": "topic.example.hi",
            "file": "data/xml_file.txt"
          }
        ],
        "teardown": []
      }
    }
  ]
}
```

>### REDIS [ see 004_redis.json ]
```javascript
{
  "redis": [
    {
      "name": "redis",
      "description": "loading redis commands from file",
      "configuration": {
        "protocol": "tcp",
        "address": "localhost:6379",
        "size": 10
      },
      "run": {
        "setup": [
          {
            "files": ["data/redis_setup_file.txt"]
          }
        ],
        "teardown": [
          {
            "commands": [
              {
                "command": "DEL",
                "arguments": [
                  "id"
                ]
              },
              {
                "command": "DEL",
                "arguments": [
                  "name"
                ]
              }
            ]
          }
        ]
      }
    },
    {
      "name": "redis",
      "description": "adding by commands",
      "run": {
        "setup": [
          {
            "commands": [
              {
                "command": "APPEND",
                "arguments": [
                  "id",
                  "1"
                ]
              },
              {
                "command": "APPEND",
                "arguments": [
                  "name",
                  "JOAO RIBEIRO"
                ]
              }
            ]
          }
        ],
        "teardown": [
          {
            "commands": [
              {
                "command": "APPEND",
                "arguments": [
                  "id",
                  "2"
                ]
              },
              {
                "command": "APPEND",
                "arguments": [
                  "name",
                  "PEDRO RIBEIRO"
                ]
              }
            ]
          }
        ]
      }
    }
  ]
}
```

>### ALL [ see 005_all.json ]
This example have all previous mocks, just to show you that you can config them all together at [go-mock/config/data/005_all.json](https://github.com/joaosoft/go-mock/tree/master/config/data/005_all.json)

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

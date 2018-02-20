# go-mock
[![Build Status](https://travis-ci.org/joaosoft/go-mock.svg?branch=master)](https://travis-ci.org/joaosoft/go-mock) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-mock/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-mock)

A package framework to create mock services. At the moment it has support for web services, redis, postgres, mysql and nsq. 
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

>### Configuration [services_001.json]
```javascript
{
  "webservices": [
    {
      "name": "hello",
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
      "name": "goodbye",
      "host": ":8002",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/goodbye",
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
      "name": "something",
      "host": ":8003",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/something",
          "payload": {
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
    }
  ],
  "sql": [
    {
      "name": "postgres",
      "description": "add users information",
      "driver": "postgres",
      "datasource": "postgres://user:password@localhost:7001?sslmode=disable",
      "commands": {
        "setup": [
          "CREATE TABLE USERS(name varchar(255), description varchar(255))",
          "INSERT INTO USERS(name, description) VALUES('joao', 'administrator')",
          "INSERT INTO USERS(name, description) VALUES('tiago', 'user')"
        ],
        "teardown": [
            "DROP TABLE USERS"
        ]
      }
    }
  ],
  "redis": [
    {
      "name": "redis",
      "configuration": {
        "protocol": "tcp",
        "addr": "localhost:6379",
        "size": 10
      },
      "commands": {
        "setup": [
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
        ],
        "teardown": [
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
    }
  ],
  "nsq": [
    {
      "name": "nsq",
      "configuration": {
        "lookupd": "localhost:4150",
        "requeue_delay": 30,
        "max_in_flight": 5,
        "max_attempts": 5
      },
      "messages": {
        "setup": [
          {
            "description": "ADD PERSON ONE",
            "topic": "topic.example.lo",
            "message": {
              "name": "joao",
              "age": 29
            }
          },
          {
            "description": "ADD PERSON TWO",
            "topic": "topic.example.hi",
            "file": "./config/xml_file.txt"
          }
        ],
        "teardown": []
      }
    }
  ]
}
```

>### Configuration [services_002.json]
```javascript
{
  "webservices": [
    {
      "name": "hello",
      "host": ":9001",
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
      "name": "goodbye",
      "host": ":9002",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/goodbye",
          "response": {
            "status": 200,
            "body": {
              "message": "Goodbye friend!"
            }
          }
        }
      ]
    }
  ],
  "sql": [
    {
      "name": "mysql",
      "description": "add clients information",
      "driver": "mysql",
      "datasource": "root:password@tcp(127.0.0.1:7002)/mysql",
      "commands": {
        "setup": [
          "CREATE TABLE CLIENTS(name varchar(255), description varchar(255))",
          "INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator')",
          "INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user')"
        ],
        "teardown": [
            "DROP TABLE CLIENTS"
        ]
      }
    }
  ],
  "redis": [
    {
      "name": "redis",
      "configuration": {
        "protocol": "tcp",
        "addr": "localhost:6379",
        "size": 10
      },
      "commands": {
        "setup": [
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
              "LUIS RIBEIRO"
            ]
          }
        ],
        "teardown": [
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
    }
  ]
}
```

>### Run
```go
import "github.com/joaosoft/go-mock"

func main() {
	gomock := gomock.NewGoMock(
        gomock.WithPath("./config"),
        gomock.WithRunInBackground(false))
    gomock.Run()
}
```

## Run example
```
make run
```

You can see that you have created the following...

##### Web services on 
> [service_001.json]
* http://localhost:8001/hello
* http://localhost:8002/goodbye
* http://localhost:8003/something

> [service_002.json]
* http://localhost:9001/hello
* http://localhost:9002/goodbye

##### Redis information on
> [service_001.json]
id: 1
name: JOAO RIBEIRO

> [service_002.json]
id: 2
name: LUIS RIBEIRO

##### Postgres on
> [service_001.json]
Created table USERS with two inserted users

##### MySQL on
> [service_002.json]
Created table CLIENTS with two inserted clients

## Logging
```
GOROOT=/usr/local/Cellar/go/1.9/libexec #gosetup
GOPATH=/Users/joaoribeiro/workspace/go/personal:/Users/joaoribeiro/workspace/go/sonae:/Users/joaoribeiro/workspace/go/others #gosetup
/usr/local/Cellar/go/1.9/libexec/bin/go build -i -o /private/var/folders/m9/qq7btgzx76l2qsqt1w64kld00000gn/T/___Run_Mock /Users/joaoribeiro/workspace/go/personal/src/go-mock/bin/launcher/main.go #gosetup
/private/var/folders/m9/qq7btgzx76l2qsqt1w64kld00000gn/T/___Run_Mock #gosetup
:: Initializing Mock Service
:: Loading file config/service_001.json
 creating service hello
 creating route /hello method GET
 started service: hello at :8001
 creating service goodbye
 creating route /goodbye method GET
 started service: goodbye at :8002
 creating service something
 creating route /something method POST
 started service: something at :8003
 creating service redis
 executing redis command: APPEND arguments:[id 1]
 executing redis command: APPEND arguments:[name JOAO RIBEIRO]
2018/02/20 16:52:47 INF    1 (localhost:4150) connecting to nsqd
 creating service nsq
 executing nsq [ ADD PERSON ONE ] message: {
              "name": "joao",
              "age": 29
            }
 executing nsq [ ADD PERSON TWO ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>
 creating service postgres
 executing SQL command: CREATE TABLE USERS(name varchar(255), description varchar(255))
 executing SQL command: INSERT INTO USERS(name, description) VALUES('joao', 'administrator')
 executing SQL command: INSERT INTO USERS(name, description) VALUES('tiago', 'user')
:: Loading file config/service_002.json
 creating service hello
 creating route /hello method GET
 started service: hello at :9001
 creating service goodbye
 creating route /goodbye method GET
 started service: goodbye at :9002
 creating service redis
 executing redis command: APPEND arguments:[id 2]
 executing redis command: APPEND arguments:[name LUIS RIBEIRO]
 creating service mysql
 executing SQL command: CREATE TABLE CLIENTS(name varchar(255), description varchar(255))
 executing SQL command: INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator')
 executing SQL command: INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user')
:: Mock Services Started
```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

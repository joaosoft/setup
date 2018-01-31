# go-mock
[![Build Status](https://travis-ci.org/joaosoft/go-mock.svg?branch=master)](https://travis-ci.org/joaosoft/go-mock) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-mock/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-mock)

A package framework to create mock services. At the moment have support to web services, redis, postgres and mysql. 
## Dependecy Management 
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
* Redis / Postgres / MySQL
```
make env
```

>### Start Application
```
make start
```

## Usage 
This example is available in the project at [go-mock/bin](https://github.com/joaosoft/go-mock/tree/master/bin)

>### Configuration services_001.json
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
    }
  ],
  "sql": [
    {
      "name": "postgres",
      "description": "add users information",
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
  ]
}
```

>### Configuration services_002.json
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
> [service_001]
* http://localhost:8001/hello
* http://localhost:8002/goodbye

> [service_002]
* http://localhost:9001/hello
* http://localhost:9002/goodbye

##### Redis information on
> [service_001]
id: 1
name: JOAO RIBEIRO

> [service_002]
id: 2
name: LUIS RIBEIRO

##### Postgres on
> [service_001]
Created table USERS with two inserted users

##### MySQL on
> [service_002]
Created table CLIENTS with two inserted clients

## Logging
```
---------- STARTING ----------
:: Initializing Mock Service
:: Loading file config/service_001.json
 Creating service hello
 Creating route /hello
 Started service: webservicehello at :8001
 Creating service goodbye
 Creating route /goodbye
 Started service: webservicegoodbye at :8002
 Creating service redis
 Executing redis command: APPEND arguments:[id 1]
 Executing redis command: APPEND arguments:[name JOAO RIBEIRO]
 Creating service postgres
 Executing SQL command: CREATE TABLE USERS(name varchar(255), description varchar(255))
 Executing SQL command: INSERT INTO USERS(name, description) VALUES('joao', 'administrator')
 Executing SQL command: INSERT INTO USERS(name, description) VALUES('tiago', 'user')
:: Loading file config/service_002.json
 Creating service hello
 Creating route /hello
 Started service: webservicehello at :9001
 Creating service goodbye
 Creating route /goodbye
 Started service: webservicegoodbye at :9002
 Creating service redis
 Executing redis command: APPEND arguments:[id 2]
 Executing redis command: APPEND arguments:[name LUIS RIBEIRO]
 Creating service mysql
 Executing SQL command: CREATE TABLE CLIENTS(name varchar(255), description varchar(255))
 Executing SQL command: INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator')
 Executing SQL command: INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user')
:: Mock Services Started
^C
---------- SHUTTING DOWN ----------
:: Stopping down Mock Service
 Teardown service hello
 TearDown webservicehello
 Teardown service goodbye
 TearDown webservicegoodbye
 Teardown service redis
 Executing redis command: DEL arguments:[id]
 Executing redis command: DEL arguments:[name]
 Teardown service postgres
 Executing SQL command: DROP TABLE USERS
 Teardown service hello
 TearDown webservicehello
 Teardown service goodbye
 TearDown webservicegoodbye
 Teardown service redis
 Executing redis command: DEL arguments:[id]
 Executing redis command: DEL arguments:[name]
 Teardown service mysql
 Executing SQL command: DROP TABLE CLIENTS
:: Stoped down Mock Service
```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

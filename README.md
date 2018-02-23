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

>## Configuration

>### WebServices [ 001_webservices.json ]

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
      "name": "goodbye",
      "description": "test goodbye",
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
      "description": "testing payload of a post",
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
    },
    {
      "name": "loading",
      "description": "loading the payload from a file",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/hello",
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
>## Configuration

>### SQL [ 002_sql.json ]
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

>### NSQ [ 003_nsq.json ]
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

>### REDIS [ 004_redis.json ]
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

>### ALL [ 005_all.json ]
This example have all previous mocks, just to show you that you can config them all together
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
      "name": "goodbye",
      "description": "test goodbye",
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
      "description": "testing payload of a post",
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
    },
    {
      "name": "loading",
      "description": "loading the payload from a file",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/hello",
          "response": {
            "status": 200,
            "file": "data/webservice_body.json"
          }
        }
      ]
    }
  ],
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
  ],
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
  ],
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

>### Run
```go
import "github.com/joaosoft/go-mock"

func main() {
func main() {
	// Run all tests at once and blocks waiting for a term command
	test1 := gomock.NewGoMock(
		gomock.WithPath("./config/all"),
		gomock.WithRunInBackground(false))
	test1.Run()

	// Run with with SQL default configuration without blocking
	configSQL := &gomock.ConfigSQL{
		DataSource: "postgres",
		Driver:     "postgres",
	}
	test2 := gomock.NewGoMock(
		gomock.WithPath("./config/once"),
		gomock.WithConfigurationSQL(configSQL))
	test2.RunSingle("service_001.json")

	// Loads setup configurations from file and run the single mock without blocking
	test3 := gomock.NewGoMock(
		gomock.WithPath("./config/once"),
		gomock.WithConfigurationFile("app.json"))
	test3.RunSingle("service_001.json")
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

### Starting
```
:: Starting Mock Service

STARTING: setup [ 001_webservices.json ]
:: Loading file [ ./config/001_webservices.json ]

 creating service [ hello ] with description [ test hello ]
 creating route [ /hello ] method [ GET ]
 started service [ hello ] at [ :8001 ]

 creating service [ goodbye ] with description [ test goodbye ]
 creating route [ /goodbye ] method [ GET ]
 started service [ goodbye ] at [ :8002 ]

 creating service [ something ] with description [ testing payload of a post ]
 creating route [ /something ] method [ POST ]
 started service [ something ] at [ :8003 ]

 creating service [ loading ] with description [ loading the payload from a file ]
 creating route [ /hello ] method [ GET ]
 started service [ loading ] at [ :8001 ]
FINISHED: setup [ 001_webservices.json ]

STARTING: setup [ 002_sql.json ]
:: Loading file [ ./config/002_sql.json ]

 creating service [ postgres ] with description [ add users information ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]
 executing SQL query [ CREATE TABLE USERS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('tiago', 'user') ]

 creating service [ postgres ] with description [ add users information from files ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_setup_file.sql ]
:: Loading file [ ./config/data/sql_setup_file.sql ]

 creating service [ mysql ] with description [ add clients information ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
 executing SQL query [ CREATE TABLE CLIENTS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user') ]
FINISHED: setup [ 002_sql.json ]

STARTING: setup [ 003_nsq.json ]
:: Loading file [ ./config/003_nsq.json ]

 creating service [ nsq ] with description [ loading a script from file and from body] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON ONE ] message: {
              "name": "joao",
              "age": 29
            }
2018/02/23 00:48:52 INF    1 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON ONE ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>

 creating service [ nsq ] with description [ ] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON TWO ] message: {
              "name": "pedro",
              "age": 30
            }
2018/02/23 00:48:52 INF    2 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON TWO ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>
FINISHED: setup [ 003_nsq.json ]

STARTING: setup [ 004_redis.json ]
:: Loading file [ ./config/004_redis.json ]

 creating service [ redis ] with description [ loading redis commands from file] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis commands...
:: Loading file [ ./config/data/redis_setup_file.txt ]

 creating service [ redis ] with description [ adding by commands] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 1] ]
 executing redis command [ APPEND ] arguments [ [name JOAO RIBEIRO] ]
FINISHED: setup [ 004_redis.json ]
:: Loading file [ ./config/data/app.json ]
Unmarshalling file [ ./config/data/app.json ] to struct
&{{{localhost:4150 30 5 5} {postgres postgres://user:password@localhost:7001?sslmode=disable} {tcp localhost:6379 10}}}

STARTING: setup [ config/001_webservices.json ]
:: Loading file [ config/001_webservices.json ]

 creating service [ hello ] with description [ test hello ]
 creating route [ /hello ] method [ GET ]
 started service [ hello ] at [ :8001 ]

 creating service [ goodbye ] with description [ test goodbye ]
 creating route [ /goodbye ] method [ GET ]
 started service [ goodbye ] at [ :8002 ]

 creating service [ something ] with description [ testing payload of a post ]
 creating route [ /something ] method [ POST ]
 started service [ something ] at [ :8003 ]

 creating service [ loading ] with description [ loading the payload from a file ]
 creating route [ /hello ] method [ GET ]
 started service [ loading ] at [ :8001 ]
FINISHED: setup [ config/001_webservices.json ]

STARTING: setup [ config/002_sql.json ]
:: Loading file [ config/002_sql.json ]

 creating service [ postgres ] with description [ add users information ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]
 executing SQL query [ CREATE TABLE USERS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('tiago', 'user') ]

 creating service [ postgres ] with description [ add users information from files ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_setup_file.sql ]
:: Loading file [ ./config/data/sql_setup_file.sql ]

 creating service [ mysql ] with description [ add clients information ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
 executing SQL query [ CREATE TABLE CLIENTS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user') ]
FINISHED: setup [ config/002_sql.json ]

STARTING: setup [ config/003_nsq.json ]
:: Loading file [ config/003_nsq.json ]

 creating service [ nsq ] with description [ loading a script from file and from body] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON ONE ] message: {
              "name": "joao",
              "age": 29
            }
2018/02/23 00:48:53 INF    3 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON ONE ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>

 creating service [ nsq ] with description [ ] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON TWO ] message: {
              "name": "pedro",
              "age": 30
            }
2018/02/23 00:48:53 INF    4 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON TWO ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>
FINISHED: setup [ config/003_nsq.json ]

STARTING: setup [ config/004_redis.json ]
:: Loading file [ config/004_redis.json ]

 creating service [ redis ] with description [ loading redis commands from file] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis commands...
:: Loading file [ ./config/data/redis_setup_file.txt ]

 creating service [ redis ] with description [ adding by commands] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 1] ]
 executing redis command [ APPEND ] arguments [ [name JOAO RIBEIRO] ]
FINISHED: setup [ config/004_redis.json ]

STARTING: setup [ config/005_all.json ]
:: Loading file [ config/005_all.json ]

 creating service [ hello ] with description [ test hello ]
 creating route [ /hello ] method [ GET ]
 started service [ hello ] at [ :8001 ]

 creating service [ goodbye ] with description [ test goodbye ]
 creating route [ /goodbye ] method [ GET ]
 started service [ goodbye ] at [ :8002 ]

 creating service [ something ] with description [ testing payload of a post ]
 creating route [ /something ] method [ POST ]
 started service [ something ] at [ :8003 ]

 creating service [ loading ] with description [ loading the payload from a file ]
 creating route [ /hello ] method [ GET ]
 started service [ loading ] at [ :8001 ]

 creating service [ redis ] with description [ loading redis commands from file] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis commands...
:: Loading file [ ./config/data/redis_setup_file.txt ]

 creating service [ redis ] with description [ adding by commands] 
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 1] ]
 executing redis command [ APPEND ] arguments [ [name JOAO RIBEIRO] ]

 creating service [ nsq ] with description [ loading a script from file and from body] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON ONE ] message: {
              "name": "joao",
              "age": 29
            }
2018/02/23 00:48:53 INF    5 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON ONE ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>

 creating service [ nsq ] with description [ ] 
 connecting with max attempts [ 5 ]
 executing nsq [ ADD PERSON TWO ] message: {
              "name": "pedro",
              "age": 30
            }
2018/02/23 00:48:53 INF    6 (localhost:4150) connecting to nsqd
:: Loading file [ ./config/data/xml_file.txt ]
 executing nsq [ ADD PERSON TWO ] message: <TEST>
    <TITLE>HELLO, THIS IS A TEST</TITLE>
</TEST>

 creating service [ postgres ] with description [ add users information ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]
 executing SQL query [ CREATE TABLE USERS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO USERS(name, description) VALUES('tiago', 'user') ]

 creating service [ postgres ] with description [ add users information from files ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_setup_file.sql ]
:: Loading file [ ./config/data/sql_setup_file.sql ]

 creating service [ mysql ] with description [ add clients information ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
 executing SQL query [ CREATE TABLE CLIENTS(name varchar(255), description varchar(255)) ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator') ]
 executing SQL query [ INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user') ]
FINISHED: setup [ config/005_all.json ]
```

### Stopping
```
:: Stopping Mock Service

STARTING: teardown [ 001_we`bservices.json ]

 teardown service [ hello ]

 teardown service [ goodbye ]

 teardown service [ something ]

 teardown service [ loading ]
FINISHED: teardown [ 001_webservices.json ]

STARTING: teardown [ 002_sql.json ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_teardown_file.sql ]
:: Loading file [ ./config/data/sql_teardown_file.sql ]

 teardown service [ mysql ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
FINISHED: teardown [ 002_sql.json ]

STARTING: teardown [ 003_nsq.json ]

 teardown service nsq
 connecting with max attempts [ 5 ]

 teardown service nsq
 connecting with max attempts [ 5 ]
FINISHED: teardown [ 003_nsq.json ]

STARTING: teardown [ 004_redis.json ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ DEL ] arguments [ [id] ]
 executing redis command [ DEL ] arguments [ [name] ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 2] ]
 executing redis command [ APPEND ] arguments [ [name PEDRO RIBEIRO] ]
FINISHED: teardown [ 004_redis.json ]

STARTING: teardown [ config/001_webservices.json ]

 teardown service [ hello ]

 teardown service [ goodbye ]

 teardown service [ something ]

 teardown service [ loading ]
FINISHED: teardown [ config/001_webservices.json ]

STARTING: teardown [ config/002_sql.json ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_teardown_file.sql ]
:: Loading file [ ./config/data/sql_teardown_file.sql ]

 teardown service [ mysql ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
FINISHED: teardown [ config/002_sql.json ]

STARTING: teardown [ config/003_nsq.json ]

 teardown service nsq
 connecting with max attempts [ 5 ]

 teardown service nsq
 connecting with max attempts [ 5 ]
FINISHED: teardown [ config/003_nsq.json ]

STARTING: teardown [ config/004_redis.json ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ DEL ] arguments [ [id] ]
 executing redis command [ DEL ] arguments [ [name] ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 2] ]
 executing redis command [ APPEND ] arguments [ [name PEDRO RIBEIRO] ]
FINISHED: teardown [ config/004_redis.json ]

STARTING: teardown [ config/005_all.json ]

 teardown service [ hello ]

 teardown service [ goodbye ]

 teardown service [ something ]

 teardown service [ loading ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ DEL ] arguments [ [id] ]
 executing redis command [ DEL ] arguments [ [name] ]

 teardown service [ redis ]
 connecting with protocol [ tcp ], address [ localhost:6379 ] and size [ 10 ]
 executing redis command [ APPEND ] arguments [ [id 2] ]
 executing redis command [ APPEND ] arguments [ [name PEDRO RIBEIRO] ]

 teardown service nsq
 connecting with max attempts [ 5 ]

 teardown service nsq
 connecting with max attempts [ 5 ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL query [ DROP TABLE IF EXISTS USERS ]

 teardown service [ postgres ]
 connecting with driver [ postgres ] and data source [ postgres://user:password@localhost:7001?sslmode=disable ]
 executing SQL file [ data/sql_teardown_file.sql ]
:: Loading file [ ./config/data/sql_teardown_file.sql ]

 teardown service [ mysql ]
 connecting with driver [ mysql ] and data source [ root:password@tcp(127.0.0.1:7002)/mysql ]
 executing SQL query [ DROP TABLE IF EXISTS CLIENTS ]
FINISHED: teardown [ config/005_all.json ]
:: Stoped Mock Service

Process finished with exit code 0
```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

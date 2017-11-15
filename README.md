# go-mock
[![Build Status](https://travis-ci.org/joaosoft/go-mock.svg?branch=master)](https://travis-ci.org/joaosoft/go-mock) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-mock/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-mock)

A package framework to create mock web services. 
## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/go-mock
```

## Usage 
This example is available in the project at [go-mock/getting_started](https://github.com/joaosoft/go-mock/tree/master/getting_started)

>### Configuration services.json
```javascript
{
  "webservices": [
    {
      "name": "hello",
      "host": "localhost:8001",
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
      "host": "localhost:8002",
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
  ]
}
```

>### Run
```go
import "github.com/joaosoft/go-mock"

func main() {
	gomock := NewGoMock(WithPath("./getting_started/config"), WithRunInBackground(false))
	gomock.Run()
}
```

## Run example
```
make getting-started
```

You can see that you have created two web services:
* http://localhost:8001/hello
* http://localhost:8002/goodbye

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

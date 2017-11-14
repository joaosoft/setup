# go-mock
[![Build Status](https://travis-ci.org/joaosoft/go-mock.svg?branch=master)](https://travis-ci.org/joaosoft/go-mock) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-mock/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-mock)

A package framework with application support. 
###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## with support to
* mock web services

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
This examples are available in the project at [go-mock/getting_started](https://github.com/joaosoft/go-mock/tree/master/getting_started)

```go
import "github.com/joaosoft/go-mock"
mock, _ := mock.NewGoMock()
mock.Run("./config/")
```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

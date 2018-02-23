package main

import (
	gomock "go-mock/services"
)

func main() {
	// WEB SERVICES
	test := gomock.NewGoMock(
		gomock.WithPath("./config"))
	test.RunSingleWait("001_webservices.json")

}

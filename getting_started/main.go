package main

import (
	gomock "go-mock/services"
)

func main() {
	gomock := gomock.NewGoMock(gomock.WithPath("./getting_started/config"), gomock.WithRunInBackground(false))
	gomock.Run()
}

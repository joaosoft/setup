package main

import (
	gomock "go-mock/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gomock := gomock.NewGoMock()
	gomock.Run("./getting_started/config")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	gomock.Stop()
}

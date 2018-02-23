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

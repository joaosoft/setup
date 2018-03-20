package main

import (
	gomock "go-mock/service"
)

func main() {
	test := gomock.NewGoMock(
		gomock.WithPath("./examples"),
		gomock.WithRunInBackground(true))

	//// web
	//test.RunSingle("001_webservices.json")
	//
	//// sql
	//configSql := &gomock.SqlConfig{
	//	Driver:     "postgres",
	//	DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
	//}
	//test.Reconfigure(gomock.WithSqlConfiguration(configSql))
	//test.RunSingle("002_sql.json")
	//
	//// nsq
	//configNSQ := &gomock.NsqConfig{
	//	Lookupd:      "localhost:4150",
	//	RequeueDelay: 30,
	//	MaxInFlight:  5,
	//	MaxAttempts:  5,
	//}
	//test.Reconfigure(gomock.WithNsqConfiguration(configNSQ))
	//test.RunSingle("003_nsq.json")
	//
	//// redis
	//configRedis := &gomock.RedisConfig{
	//	Protocol: "tcp",
	//	Address:  "localhost:6379",
	//	Size:     10,
	//}
	//test.Reconfigure(gomock.WithRedisConfiguration(configRedis))
	//test.RunSingle("004_redis.json")

	// all
	test.Reconfigure(
		gomock.WithConfigurationFile("data/config.json"))

	test.Run()
	test.Wait()
	test.Stop()
}

package main

import (
	"setup"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	logger "github.com/joaosoft/logger"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	log := logger.NewLogDefault("setup", logger.InfoLevel)
	start := time.Now()
	test := setup.NewSetup(
		setup.WithPath("./examples"),
		setup.WithRunInBackground(true))

	//// web
	//test.RunSingle("001_webservices.json")
	//
	//// sql
	//configSql := &setup.SqlConfig{
	//	Driver:     "postgres",
	//	DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
	//}
	//test.Reconfigure(setup.WithSqlConfiguration(configSql))
	//test.RunSingle("002_sql.json")
	//
	//// nsq
	//configNSQ := &setup.NsqConfig{
	//	Lookupd:      "localhost:4150",
	//	RequeueDelay: 30,
	//	MaxInFlight:  5,
	//	MaxAttempts:  5,
	//}
	//test.Reconfigure(setup.WithNsqConfiguration(configNSQ))
	//test.RunSingle("003_nsq.json")
	//
	//// redis
	//configRedis := &setup.RedisConfig{
	//	Protocol: "tcp",
	//	Address:  "localhost:6379",
	//	Size:     10,
	//}
	//test.Reconfigure(setup.WithRedisConfiguration(configRedis))
	//test.RunSingle("004_redis.json")

	//// files
	//test.RunSingle("005_files.json")

	// all
	test.Reconfigure(
		setup.WithConfigurationFile("examples/data/config.json"))

	test.Run()
	test.Wait()
	test.Stop()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed)
}

package main

import (
	gosetup "go-setup/service"
	"time"

	"github.com/labstack/gommon/log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

func main() {
	start := time.Now()
	test := gosetup.NewGoSetup(
		gosetup.WithPath("./example"),
		gosetup.WithRunInBackground(true))

	//// web
	//test.RunSingle("001_webservices.json")
	//
	//// sql
	//configSql := &gosetup.SqlConfig{
	//	Driver:     "postgres",
	//	DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
	//}
	//test.Reconfigure(gosetup.WithSqlConfiguration(configSql))
	//test.RunSingle("002_sql.json")
	//
	//// nsq
	//configNSQ := &gosetup.NsqConfig{
	//	Lookupd:      "localhost:4150",
	//	RequeueDelay: 30,
	//	MaxInFlight:  5,
	//	MaxAttempts:  5,
	//}
	//test.Reconfigure(gosetup.WithNsqConfiguration(configNSQ))
	//test.RunSingle("003_nsq.json")
	//
	//// redis
	//configRedis := &gosetup.RedisConfig{
	//	Protocol: "tcp",
	//	Address:  "localhost:6379",
	//	Size:     10,
	//}
	//test.Reconfigure(gosetup.WithRedisConfiguration(configRedis))
	//test.RunSingle("004_redis.json")

	//// files
	//test.RunSingle("005_files.json")

	// all
	test.Reconfigure(
		gosetup.WithConfigurationFile("example/data/config.json"))

	test.Run()
	test.Wait()
	test.Stop()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed)
}

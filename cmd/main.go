package main

import (
	"fmt"

	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/redis"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlUser     = "test"
	mysqlPassword = "test"
	mysqlDB       = "test"

	redisHost = "localhost"
	redisPort = 6379
)

func main() {
	var err error
	if err = parseConfigFile(); err != nil {
		logrus.Fatal("解析配置文件错误", err)
	}
	logrus.Infoln(mysqlHost, mysqlPort, mysqlUser, mysqlPassword, mysqlDB, redisHost, redisPort)
	common.RedisClientService = redis.New(fmt.Sprintf("%s:%d", redisHost, redisPort))
	common.DB, err = model.NewDB(&model.Options{
		User:     mysqlUser,
		Password: mysqlPassword,
		DB:       mysqlDB,
		Addr:     fmt.Sprintf("%s:%d", mysqlHost, mysqlPort),
	})
	if err != nil {
		logrus.Fatal("启动mysql错误", err)
	}

}

func parseConfigFile() error {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs/")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	mysqlHost = viper.GetString("db.mysql.host")
	mysqlPort = viper.GetInt("db.mysql.port")
	mysqlUser = viper.GetString("db.mysql.user")
	mysqlPassword = viper.GetString("db.mysql.password")
	mysqlDB = viper.GetString("db.mysql.db")

	redisHost = viper.GetString("db.redis.host")
	redisPort = viper.GetInt("db.redis.port")
	return nil
}

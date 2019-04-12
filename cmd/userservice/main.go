package userservice

import (
	"fmt"
	"net"
	"os"
	"time"

	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/pb"
	"userService/pkg/redis"
	"userService/pkg/userservice"

	"github.com/go-kit/kit/sd/consul"
	consuld "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlUser     = "test"
	mysqlPassword = "test"
	mysqlDB       = "test"

	redisHost = "localhost"
	redisPort = 6379

	grpcHost = "localhost"
	grpcPort = 5001

	consulHost = "localhost"
	consulPort = 8500
)

func main() {
	var err error
	if err = parseConfigFile(); err != nil {
		logrus.Fatal("解析配置文件错误", err)
	}
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

	consulClient, err := newConsulClient(fmt.Sprintf("%s:%d", consulHost, consulPort))
	if err != nil {
		logrus.Fatal("连接consul失败", err)
	}

	go func() {
		if err := runGRPCServer(fmt.Sprintf("%s:%d", grpcHost, grpcPort)); err != nil {
			logrus.Fatal("grpc server shutdown", err)
		}
	}()

	err = registerConsulService(consulClient, "userService", grpcHost, grpcPort)
	if err != nil {
		logrus.Errorln("注册userService失败", err)
	}
	logrus.Infoln("启动成功")
	for {
		time.Sleep(time.Hour)
	}
}

func parseConfigFile() error {
	fileName := os.Getenv("CONFIG_FILE")

	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs/")

	if fileName != "" {
		viper.SetConfigFile(fileName)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	settings := viper.AllSettings()
	logrus.Infoln(settings)

	mysqlHost = viper.GetString("db.mysql.host")
	mysqlPort = viper.GetInt("db.mysql.port")
	mysqlUser = viper.GetString("db.mysql.user")
	mysqlPassword = viper.GetString("db.mysql.password")
	mysqlDB = viper.GetString("db.mysql.db")

	redisHost = viper.GetString("db.redis.host")
	redisPort = viper.GetInt("db.redis.port")

	grpcHost = viper.GetString("grpc.host")
	grpcPort = viper.GetInt("grpc.port")

	if host := os.Getenv("GRPC_HOST"); host != "" {
		grpcHost = host
	}

	consulHost = viper.GetString("consul.host")
	consulPort = viper.GetInt("consul.port")
	return nil
}

func newConsulClient(addr string) (consul.Client, error) {
	consulClient, err := consuld.NewClient(&consuld.Config{
		Address: addr,
	})
	if err != nil {
		return nil, err
	}
	return consul.NewClient(consulClient), nil
}

func registerConsulService(client consul.Client, name string, host string, port int) error {
	return client.Register(&consuld.AgentServiceRegistration{
		Address: host,
		Port:    port,
		Name:    name,
	})
}

func runGRPCServer(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	svr := grpc.NewServer()
	pb.RegisterUserServer(svr, userservice.New())
	return svr.Serve(l)
}

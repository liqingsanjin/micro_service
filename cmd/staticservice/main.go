package main

import (
	"fmt"
	"net"
	"os"
	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/pb"
	"userService/pkg/staticservice"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	chanHTTPErr := make(chan error)

	conf, err := ParseConfigFile()
	if err != nil {
		logrus.Fatal("解析配置文件错误", err)
	}

	fmt.Println("正在链接mysql...")
	common.DB, err = model.NewDB(&model.Options{
		User:     conf.MysqlUser,
		Password: conf.MysqlPassword,
		DB:       conf.MysqlDB,
		Addr:     fmt.Sprintf("%s:%d", conf.MysqlHost, conf.MysqlPort),
	})
	defer common.DB.Close()

	if err != nil {
		logrus.Fatal("启动mysql错误", err)
	}

	fmt.Println("启动consul watcher ...")
	go staticservice.StartServer(conf.WatcherAddr, chanHTTPErr)

	go func() {
		if err := runGRPCServer(fmt.Sprintf("%s:%d", conf.GrpcHost, conf.GrpcPort)); err != nil {
			logrus.Fatal("grpc server shutdown", err)
		}
	}()

	//register service.
	fmt.Println("正在链接consul...")
	consulClient, err := newConsulClient(fmt.Sprintf("%s:%d", conf.ConsulHost, conf.ConsulPort))

	if err != nil {
		logrus.Fatal("consul链接失败： ", err)
	}

	err = registerService(consulClient, conf.ServiceName, conf.GrpcHost, conf.GrpcPort)
	if err != nil {
		logrus.Fatal("服务注册失败:", err)
	}

	logrus.Info("启动成功")

	select {
	case err := <-chanHTTPErr:
		logrus.Fatal(err)
	}
}

func runGRPCServer(addr string) error {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	insHandle := staticservice.NewGRPCServer()

	svr := grpc.NewServer()
	pb.RegisterStaticServer(svr, insHandle)
	return svr.Serve(l)
}

func newConsulClient(addr string) (consul.Client, error) {
	client, err := api.NewClient(&api.Config{
		Address: addr,
	})
	if err != nil {
		return nil, err
	}
	common.ConsulClient = client

	return consul.NewClient(client), nil
}

func registerService(client consul.Client, name, host string, port int) error {
	return client.Register(&api.AgentServiceRegistration{
		Address: host,
		Port:    port,
		Name:    name,
	})
}

//Conf .
type Conf struct {
	MysqlHost     string
	MysqlPort     int
	MysqlUser     string
	MysqlPassword string
	MysqlDB       string
	RedisHost     string
	RedisPort     int
	GrpcHost      string
	GrpcPort      int
	ConsulHost    string
	ConsulPort    int
	ServiceName   string
	WatcherAddr   string
}

//ParseConfigFile .
func ParseConfigFile() (*Conf, error) {
	fileName := os.Getenv("CONFIG_FILE")

	viper.SetConfigType("toml")
	viper.AddConfigPath("./cmd/staticservice")

	if fileName != "" {
		viper.SetConfigFile(fileName)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	settings := viper.AllSettings()
	logrus.Infoln(settings)

	conf := Conf{}
	conf.MysqlHost = viper.GetString("db.mysql.host")
	conf.MysqlPort = viper.GetInt("db.mysql.port")
	conf.MysqlUser = viper.GetString("db.mysql.user")
	conf.MysqlPassword = viper.GetString("db.mysql.password")
	conf.MysqlDB = viper.GetString("db.mysql.db")

	conf.RedisHost = viper.GetString("db.redis.host")
	conf.RedisPort = viper.GetInt("db.redis.port")

	conf.GrpcHost = viper.GetString("grpc.host")
	conf.GrpcPort = viper.GetInt("grpc.port")

	conf.WatcherAddr = viper.GetString("watcher.addr")

	conf.ServiceName = viper.GetString("info.serviceName")

	if host := os.Getenv("GRPC_HOST"); host != "" {
		conf.GrpcHost = host
	}

	conf.ConsulHost = viper.GetString("consul.host")
	conf.ConsulPort = viper.GetInt("consul.port")

	return &conf, err
}

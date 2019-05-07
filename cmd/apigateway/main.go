package main

import (
	"net/http"
	"os"
	"userService/pkg/gateway"
	"userService/pkg/util"

	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func main() {
	level := os.Getenv("LOG_LEVEL")
	if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&util.LogFormatter{})
	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}
	consulAddr := os.Getenv("CONSUL_ADDR")

	consulConfig := api.DefaultConfig()
	if consulAddr != "" {
		consulConfig.Address = consulAddr
	}
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	client := consulsd.NewClient(consulClient)
	log := &logger{}

	var (
		tags                 = []string{}
		passingOnly          = true
		instancer            = consulsd.NewInstancer(client, log, "userService", tags, passingOnly)
		staticInstancer      = consulsd.NewInstancer(client, log, "staticService", tags, passingOnly)
		institutionInstancer = consulsd.NewInstancer(client, log, "institutionService", tags, passingOnly)
	)

	userEndpoint := gateway.GetUserEndpoints(instancer, log)
	staticEndpoint := gateway.GetStaticCliEndpoints(staticInstancer, log)
	institutionEndpoint := gateway.GetInstitutionCliEndpoints(institutionInstancer, log)

	userHandler := gateway.NewHttpHandler(userEndpoint, &staticEndpoint, &institutionEndpoint)
	http.ListenAndServe(":"+port, userHandler)
}

type logger struct{}

func (l *logger) Log(args ...interface{}) error {
	logrus.Infoln(args...)
	return nil
}

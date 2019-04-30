package main

import (
	"net/http"
	"userService/pkg/gateway"

	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	client := consulsd.NewClient(consulClient)
	log := &logger{}

	var (
		tags        = []string{}
		passingOnly = true
		instancer   = consulsd.NewInstancer(client, log, "userService", tags, passingOnly)
	)

	userEndpoint := gateway.GetUserEndpoints(instancer, log)
	staticEndpoint := GetStaticCliEndpoints()

	userHandler := gateway.NewHttpHandler(userEndpoint, &staticEndpoint)
	http.ListenAndServe(":8080", userHandler)
}

type logger struct{}

func (l *logger) Log(args ...interface{}) error {
	logrus.Infoln(args...)
	return nil
}

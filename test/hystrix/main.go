package main

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sirupsen/logrus"
)

func main() {
	hystrix.ConfigureCommand("my", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})
	err := hystrix.Do("my", func() error {
		time.Sleep(10 * time.Second)
		return nil
	}, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infoln("my success")
}

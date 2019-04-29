package main

import (
	"io"
	"time"
	"userService/pkg/gateway"
	"userService/pkg/pb"
	"userService/pkg/staticservice"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/hashicorp/consul/api"
	stdzipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	staticbreaker = "staticbreaker"
)

func GetStaticCliEndpoints() gateway.StaticEndpoints {
	log := &logger{}
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	client := consulsd.NewClient(consulClient)
	var (
		tags        = []string{}
		passingOnly = true
		instancer   = consulsd.NewInstancer(client, log, "staticService", tags, passingOnly)
		endpoints   gateway.StaticEndpoints
	)

	hystrix.ConfigureCommand(staticbreaker, hystrix.CommandConfig{
		MaxConcurrentRequests: 1000,
		Timeout:               10000,
		ErrorPercentThreshold: 25,
		SleepWindow:           10000,
	})
	staticBreaker := circuitbreaker.Hystrix(staticbreaker)

	{
		factory := staticserviceFactory(staticservice.MakeSyncDataEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = staticBreaker(retry)
		endpoints.SyncDataEndpoint = retry
	}

	{
		factory := staticserviceFactory(staticservice.MakeGetDictionaryItemEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = staticBreaker(retry)
		endpoints.GetDictionaryItemEndpoint = retry
	}

	{
		factory := staticserviceFactory(staticservice.MakeGetDicByProdAndBizEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = staticBreaker(retry)
		endpoints.GetDicByProdAndBizEndpoint = retry
	}

	{
		factory := staticserviceFactory(staticservice.MakeGetDicByInsCmpCdEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = staticBreaker(retry)
		endpoints.GetDicByInsCmpCdEndpoint = retry
	}

	{
		factory := staticserviceFactory(staticservice.MakeCheckValuesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = staticBreaker(retry)
		endpoints.CheckValuesEndpoint = retry
	}

	return endpoints
}

func staticserviceFactory(makeEndpoint func(pb.StaticServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		localEndpoint, _ := stdzipkin.NewEndpoint("static", "localhost:9411")
		reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
		stdTracer, err := stdzipkin.NewTracer(
			reporter,
			stdzipkin.WithLocalEndpoint(localEndpoint),
		)
		if err != nil {
			logrus.Errorln(err)
		}

		var service *gateway.StaticEndpoints
		if stdTracer == nil {
			service = gateway.NewStaticServiceGRPCClient(conn, nil)
		} else {
			tracer := zipkin.GRPCClientTrace(stdTracer)
			service = gateway.NewStaticServiceGRPCClient(conn, tracer)
		}

		return makeEndpoint(service), conn, nil
	}
}

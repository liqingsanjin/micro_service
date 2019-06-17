package gateway

import (
	"io"
	"time"
	"userService/pkg/institutionservice"
	"userService/pkg/pb"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	institutionbreaker = "institutionbreaker"
)

func GetInstitutionCliEndpoints(instancer sd.Instancer, log log.Logger) *InstitutionEndpoints {
	var endpoints InstitutionEndpoints

	hystrix.ConfigureCommand(institutionbreaker, hystrix.CommandConfig{
		MaxConcurrentRequests: 1000,
		Timeout:               10000,
		ErrorPercentThreshold: 25,
		SleepWindow:           10000,
	})
	institutionBreaker := circuitbreaker.Hystrix(institutionbreaker)

	{
		factory := institutionserviceFactory(institutionservice.MakeTnxHisDownloadEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.TnxHisDownloadEndpoint = retry
	}

	{
		factory := institutionserviceFactory(institutionservice.MakeGetTfrTrnLogsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.GetTfrTrnLogsEndpoint = retry
	}

	{
		factory := institutionserviceFactory(institutionservice.MakeGetTfrTrnLogEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.GetTfrTrnLogEndpoint = retry
	}

	{
		factory := institutionserviceFactory(institutionservice.MakeDownloadTfrTrnLogsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.DownloadTfrTrnLogsEndpoint = retry
	}

	{
		factory := institutionserviceFactory(institutionservice.MakeListGroupsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.ListGroupsEndpoint = retry
	}

	{
		factory := institutionserviceFactory(institutionservice.MakeListInstitutionsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = institutionBreaker(retry)
		endpoints.ListInstitutionsEndpoint = retry
	}
	return &endpoints
}

func institutionserviceFactory(makeEndpoint func(pb.InstitutionServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		localEndpoint, _ := stdzipkin.NewEndpoint("institution", "localhost:9411")
		reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
		stdTracer, err := stdzipkin.NewTracer(
			reporter,
			stdzipkin.WithLocalEndpoint(localEndpoint),
		)
		if err != nil {
			logrus.Errorln(err)
		}

		var service *InstitutionEndpoints
		if stdTracer == nil {
			service = NewInstitutionServiceGRPCClient(conn, nil)
		} else {
			tracer := zipkin.GRPCClientTrace(stdTracer)
			service = NewInstitutionServiceGRPCClient(conn, tracer)
		}

		return makeEndpoint(service), conn, nil
	}
}

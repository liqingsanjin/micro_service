package gateway

import (
	"io"
	"time"
	"userService/pkg/pb"
	"userService/pkg/productservice"

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

const productBreaker = "productbreaker"

func GetProductEndpoints(instancer sd.Instancer, log log.Logger) *ProductEndpoints {
	endpoints := new(ProductEndpoints)

	hystrix.ConfigureCommand(productBreaker, hystrix.CommandConfig{
		MaxConcurrentRequests: 1000,
		Timeout:               10000,
		ErrorPercentThreshold: 25,
		SleepWindow:           10000,
	})
	breaker := circuitbreaker.Hystrix(productBreaker)

	{
		factory := productServiceFactory(productservice.MakeListTransMapEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = breaker(retry)
		endpoints.ListTransMapEndpoint = retry
	}

	{
		factory := productServiceFactory(productservice.MakeListFeeMapEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = breaker(retry)
		endpoints.ListFeeMapEndpoint = retry
	}

	return endpoints
}

func productServiceFactory(makeEndpoint func(pb.ProductServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		localEndpoint, _ := stdzipkin.NewEndpoint("product", "localhost:9411")
		reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
		stdTracer, err := stdzipkin.NewTracer(
			reporter,
			stdzipkin.WithLocalEndpoint(localEndpoint),
		)
		if err != nil {
			logrus.Errorln(err)
		}

		var service *ProductEndpoints
		if stdTracer == nil {
			service = NewProductServiceClient(conn, nil)
		} else {
			tracer := zipkin.GRPCClientTrace(stdTracer)
			service = NewProductServiceClient(conn, tracer)
		}

		return makeEndpoint(service), conn, nil
	}
}

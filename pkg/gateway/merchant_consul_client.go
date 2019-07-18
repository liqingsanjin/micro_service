package gateway

import (
	"io"
	"os"
	"time"
	merchantservice "userService/pkg/merchantservice"
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

func GetMerchantEndpoints(instancer sd.Instancer, log log.Logger) *MerchantEndpoints {
	var endpoints MerchantEndpoints

	hystrix.ConfigureCommand(userbreaker, hystrix.CommandConfig{
		MaxConcurrentRequests: 1000,
		Timeout:               10000,
		ErrorPercentThreshold: 25,
		SleepWindow:           10000,
	})
	userBreaker := circuitbreaker.Hystrix(userbreaker)

	{
		factory := merchantServiceFactory(merchantservice.MakeListMerchantEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.ListMerchantEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeListGroupMerchantEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.ListGroupMerchantEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantBankAccountEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantBankAccountEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantBizDealEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantBizDealEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantBizFeeEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantBizFeeEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantBusinessEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantBusinessEndpoint = retry
	}

	{
		factory := merchantServiceFactory(merchantservice.MakeSaveMerchantPictureEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.SaveMerchantPictureEndpoint = retry
	}

	return &endpoints
}

func merchantServiceFactory(makeEndpoint func(pb.MerchantServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		addr := os.Getenv("ZIPKIN_ADDR")
		if addr == "" {
			addr = "127.0.0.1:9411"
		}

		localEndpoint, _ := stdzipkin.NewEndpoint("merchant", "localhost:9411")
		reporter := zipkinhttp.NewReporter("http://" + addr + "/api/v2/spans")
		stdTracer, err := stdzipkin.NewTracer(
			reporter,
			stdzipkin.WithLocalEndpoint(localEndpoint),
		)
		if err != nil {
			logrus.Errorln(err)
		}
		var service *MerchantEndpoints
		if stdTracer == nil {
			service = NewMerchantServiceClient(conn, nil)
		} else {
			tracer := zipkin.GRPCClientTrace(stdTracer)
			service = NewMerchantServiceClient(conn, tracer)
		}

		return makeEndpoint(service), conn, nil
	}
}

package main

import (
	"io"
	"net/http"
	"time"
	"userService/pkg/gateway"
	"userService/pkg/pb"
	"userService/pkg/userservice"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		endpoints   gateway.UserEndpoints
	)

	{
		factory := userserviceFactory(userservice.MakeLoginEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		endpoints.LoginEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeGetPermissionsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.GetPermissionsEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeCheckPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.CheckPermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRegisterEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RegisterEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddPermissionForRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddPermissionForRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeCreateRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.CreateRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddRoleForUserEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddRoleForUserEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddRoutesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddRoutesEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeListRoutesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.ListRoutesEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeCreatePermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.CreatePermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeUpdatePermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.UpdatePermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddRouteForPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddRouteForPermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemoveRouteForPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemoveRouteForPermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemovePermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemovePermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeListPermissionsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.ListPermissionsEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddPermissionForPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddPermissionForPermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemovePermissionForPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemovePermissionForPermissionEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeListRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.ListRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeUpdateRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.UpdateRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemovePermissionForRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemovePermissionForRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddRoleForRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddRoleForRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemoveRoleForRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemoveRoleForRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemoveRoleEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemoveRoleEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeListUsersEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.ListUsersEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeUpdateUserEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.UpdateUserEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeAddPermissionForUserEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddPermissionForUserEndpoint = retry
	}

	{
		factory := userserviceFactory(userservice.MakeRemovePermissionForUserEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RemovePermissionForUserEndpoint = retry
	}

	userHandler := gateway.NewHttpHandler(&endpoints)
	http.ListenAndServe(":8080", userHandler)
}

func userserviceFactory(makeEndpoint func(pb.UserServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := gateway.NewUserServiceGRPCClient(conn)
		return makeEndpoint(service), conn, nil
	}
}

type logger struct{}

func (l *logger) Log(args ...interface{}) error {
	logrus.Infoln(args...)
	return nil
}

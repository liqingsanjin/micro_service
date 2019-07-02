package camunda

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
)

type Service struct {
	ProcessDefinition pb.ProcessDefinitionServer
}

var defaultService *Service

func Load(client consul.Client, log log.Logger) {
	defaultService = new(Service)
	var (
		tags        []string
		passingOnly = true
	)
	processDefinitionInstancer := consul.NewInstancer(client, log, "camundaService", tags, passingOnly)
	defaultService.ProcessDefinition = GetProcessDefinitionConsulEndpoints(processDefinitionInstancer, log)
}

func Get() *Service {
	return defaultService
}

func encodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

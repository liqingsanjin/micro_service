package camunda

import (
	"userService/pkg/gateway"
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
	processDefinitionInstancer := consul.NewInstancer(client, log, "processDefinitionService", tags, passingOnly)
	defaultService.ProcessDefinition = gateway.GetProcessDefinitionConsulEndpoints(processDefinitionInstancer, log)
}

func Get() *Service {
	return defaultService
}

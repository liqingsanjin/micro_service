package workflow

import (
	"context"
	"fmt"
	"net/http"
	"userService/pkg/camunda"
	camundapb "userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/pb"
)

type service struct{}

func (s *service) ListTask(context.Context, *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	panic("implement me")
}

func (s *service) HandleTask(context.Context, *pb.HandleTaskRequest) (*pb.HandleTaskReply, error) {
	panic("implement me")
}

func (s *service) Start(ctx context.Context, in *pb.StartWorkflowRequest) (*pb.StartWorkflowReply, error) {
	reply := new(pb.StartWorkflowReply)
	if in.Name == "" || in.Id == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamError",
			Description: "工作流和id不能为空",
		}
		return reply, nil
	}
	db := common.DB

	processes, err := camundamodel.QueryProcessDefinition(db, &camundamodel.ProcessDefinition{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	if len(processes) == 0 {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamError",
			Description: "工作流不存在",
		}
		return reply, nil
	}

	client := camunda.Get()

	startProcessInstanceRes, err := client.ProcessDefinition.Start(ctx, &camundapb.StartProcessDefinitionReq{
		Id: processes[0].Id,
		Body: &camundapb.StartProcessDefinitionReqBody{
			BusinessKey: fmt.Sprintf("%s:%s", in.Name, in.Id),
		},
	})
	if err != nil {
		return nil, err
	}

	if camunda.CheckError(startProcessInstanceRes) {
		reply.Err = camunda.TransError(startProcessInstanceRes)
		return reply, nil
	}
	return reply, nil
}

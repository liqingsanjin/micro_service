package workflow

import (
	"context"
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
	if in.Name == "" || in.Type == "" || in.UserId == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamError",
			Description: "参数不能为空",
		}
		return reply, nil
	}
	db := common.DB

	processes, err := camundamodel.QueryProcessDefinition(db, &camundamodel.ProcessDefinition{
		Name: in.Type,
	})
	if err != nil {
		return nil, err
	}
	if len(processes) == 0 {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamError",
			Description: "工作流类型不存在",
		}
		return reply, nil
	}

	client := camunda.Get()

	startProcessInstanceRes, err := client.ProcessDefinition.Start(ctx, &camundapb.StartProcessDefinitionReq{
		Id: processes[0].Id,
		Body: &camundapb.StartProcessDefinitionReqBody{
			BusinessKey: in.Name,
		},
	})
	if err != nil {
		return nil, err
	}

	if camunda.CheckError(startProcessInstanceRes) {
		reply.Err = camunda.TransError(startProcessInstanceRes)
		return reply, nil
	}

	// 获取任务
	taskListRes, err := client.Task.GetList(ctx, &camundapb.GetListTaskReq{
		ProcessInstanceId: startProcessInstanceRes.Item.Id,
	})
	if err != nil {
		return nil, err
	}

	if camunda.CheckError(taskListRes) {
		reply.Err = camunda.TransError(taskListRes)
		return reply, nil
	}

	// todo 写入任务列表
	if len(taskListRes.Tasks) == 0 {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "WorkflowError",
			Description: "没有任务需要执行",
		}
		return reply, nil
	}
	task := taskListRes.Tasks[0]
	err = camundamodel.SaveTask(db, &camundamodel.Task{
		Title:         in.Name,
		UserId:        in.UserId,
		CurrentNode:   task.Name,
		CamundaTaskId: task.Id,
		InstanceId:    startProcessInstanceRes.Item.Id,
	})

	return reply, err
}

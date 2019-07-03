package workflow

import (
	"context"
	"net/http"
	"userService/pkg/camunda"
	camundapb "userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/pb"
	"userService/pkg/util"
)

type service struct{}

func (s *service) ListTask(ctx context.Context, in *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}
	db := common.DB
	query := new(camundamodel.Task)
	if in.Item != nil {
		query.TaskId = in.Item.TaskId
		query.Title = in.Item.Title
		query.UserId = in.Item.UserId
		query.CurrentNode = in.Item.CurrentNode
		query.CamundaTaskId = in.Item.CamundaTaskId
		query.InstanceId = in.Item.InstanceId
		query.EndFlag = &in.Item.EndFlag
	}
	tasks, count, err := camundamodel.QueryTask(db, query, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	pbTasks := make([]*pb.TaskField, len(tasks))
	for i := range tasks {
		pbTasks[i] = &pb.TaskField{
			TaskId:        tasks[i].TaskId,
			Title:         tasks[i].Title,
			UserId:        tasks[i].UserId,
			CurrentNode:   tasks[i].CurrentNode,
			CamundaTaskId: tasks[i].CamundaTaskId,
			InstanceId:    tasks[i].InstanceId,
			EndFlag:       *tasks[i].EndFlag,
			CreatedAt:     tasks[i].CreatedAt.Format("2006-01-02 15:03:04"),
			UpdatedAt:     tasks[i].UpdatedAt.Format("2006-01-02 15:03:04"),
		}
	}

	return &pb.ListTaskReply{
		Page:  in.Page,
		Size:  in.Size,
		Count: count,
		Items: pbTasks,
	}, nil
}

func (s *service) HandleTask(ctx context.Context, in *pb.HandleTaskRequest) (*pb.HandleTaskReply, error) {
	reply := new(pb.HandleTaskReply)
	if in.TaskId == 0 || in.Result == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamError",
			Description: "result和taskId不能为空",
		}
		return reply, nil
	}
	db := common.DB.Begin()
	defer db.Rollback()
	task, err := camundamodel.FindTaskById(db, in.TaskId)
	if err != nil {
		return nil, err
	}

	client := camunda.Get()
	values := make(map[string]*camundapb.Variable)
	values["result"] = &camundapb.Variable{
		Value: in.Result,
		Type:  "string",
	}

	// 保存备注
	err = camundamodel.SaveRemark(db, &camundamodel.Remark{
		Comment: in.Remark,
		TaskId:  in.TaskId,
	})
	if err != nil {
		return nil, err
	}
	// 完成任务
	completeTaskRes, err := client.Task.Complete(ctx, &camundapb.CompleteTaskReq{
		Id: task.CamundaTaskId,
		Body: &camundapb.CompleteTaskReqBody{
			Variables: values,
		},
	})
	if err != nil {
		return nil, err
	}

	if camunda.CheckError(completeTaskRes) {
		reply.Err = camunda.TransError(completeTaskRes)
		return reply, nil
	}

	// 查询下一个任务
	listTaskRes, err := client.Task.GetList(ctx, &camundapb.GetListTaskReq{
		ProcessInstanceId: task.InstanceId,
	})
	if err != nil {
		return nil, err
	}
	if camunda.CheckError(listTaskRes) {
		reply.Err = camunda.TransError(listTaskRes)
		return reply, nil
	}

	if len(listTaskRes.Tasks) != 0 {
		// 有下个任务，更新任务节点
		err = camundamodel.UpdateTask(db, &camundamodel.Task{
			TaskId: task.TaskId,
		}, &camundamodel.Task{
			CurrentNode:   listTaskRes.Tasks[0].Name,
			CamundaTaskId: listTaskRes.Tasks[0].Id,
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 修改状态为结束
		t := new(camundamodel.Task)
		*t.EndFlag = true
		err = camundamodel.UpdateTask(db, &camundamodel.Task{
			TaskId: task.TaskId,
		}, t)
	}
	db.Commit()

	return reply, nil
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

	// 写入任务列表
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

func (s *service) ListRemark(ctx context.Context, in *pb.ListRemarkRequest) (*pb.ListRemarkReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}
	query := new(camundamodel.Remark)
	if in.Item != nil {
		query.RemarkId = in.Item.RemarkId
		query.Comment = in.Item.Comment
		query.TaskId = in.Item.TaskId

	}
	db := common.DB
	items, count, err := camundamodel.QueryRemark(db, query, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	docs := make([]*pb.RemarkField, len(items))
	for i := range items {
		docs[i] = &pb.RemarkField{
			RemarkId:  items[i].RemarkId,
			Comment:   items[i].Comment,
			TaskId:    items[i].TaskId,
			CreatedAt: items[i].CreatedAt.Format(util.TimePattern),
			UpdatedAt: items[i].UpdatedAt.Format(util.TimePattern),
		}
	}

	return &pb.ListRemarkReply{
		Items: docs,
		Count: count,
		Page:  in.Page,
		Size:  in.Size,
	}, nil
}

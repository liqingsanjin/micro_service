package workflow

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"userService/pkg/camunda"
	camundapb "userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/user"
	"userService/pkg/pb"
	"userService/pkg/util"

	"google.golang.org/grpc/metadata"
)

const (
	InvalidParam  = "InvalidParamError"
	AlreadyExists = "AlreadyExistsError"
	NotFound      = "NotFoundError"
	INTERNAL      = "InternalServerError"
)

type service struct{}

func (s *service) ListTask(ctx context.Context, in *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}
	reply := new(pb.ListTaskReply)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "找不到用户信息",
		}
		return reply, nil
	}
	ids := md.Get("userid")
	if !ok || len(ids) == 0 {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "找不到用户信息",
		}
		return reply, nil
	}
	all := common.Enforcer.GetImplicitRolesForUser(fmt.Sprintf("user:%s", ids[0]))
	roleIds := make([]int64, 0)
	for _, a := range all {
		if strings.HasPrefix(a, "role") {
			ss := strings.Split(a, ":")
			i, _ := strconv.Atoi(ss[1])
			roleIds = append(roleIds, int64(i))
		}
	}

	db := common.DB

	tasks, count, err := camundamodel.FindTaskByRoles(db, roleIds, in.Page, in.Size)
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
		Comment:    in.Remark,
		TaskId:     in.TaskId,
		InstanceId: task.InstanceId,
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
	// 修改当前状态为结束
	t := new(camundamodel.Task)
	endFlag := true
	t.EndFlag = &endFlag
	err = camundamodel.UpdateTask(db, &camundamodel.Task{
		TaskId: task.TaskId,
	}, t)
	if err != nil {
		return nil, err
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
		// 有下个任务，保存新的任务节点
		for _, t := range listTaskRes.Tasks {
			endFlag := false
			role, err := user.FindRole(db, t.Assignee)
			if err != nil {
				return nil, err
			}
			if role == nil {
				reply.Err = &pb.Error{
					Code:        http.StatusInternalServerError,
					Message:     INTERNAL,
					Description: "角色不存在",
				}
				return reply, nil
			}

			err = camundamodel.SaveTask(db, &camundamodel.Task{
				Title:         task.Title,
				UserId:        task.UserId,
				Role:          role.ID,
				CurrentNode:   t.Name,
				CamundaTaskId: t.Id,
				InstanceId:    t.ProcessInstanceId,
				EndFlag:       &endFlag,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	db.Commit()

	return reply, nil
}

func (s *service) Start(ctx context.Context, in *pb.StartWorkflowRequest) (*pb.StartWorkflowReply, error) {
	reply := new(pb.StartWorkflowReply)

	if in.Name == "" || in.Type == "" || in.UserId == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "参数不能为空",
		}
		return reply, nil
	}
	db := common.DB.Begin()
	defer db.Rollback()

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

	err = camundamodel.SaveProcessInstance(db, &camundamodel.ProcessInstance{
		Id:     startProcessInstanceRes.Item.Id,
		Title:  in.Name,
		DataId: in.DataId,
	})
	if err != nil {
		return nil, err
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

	for _, task := range taskListRes.Tasks {
		var endFlag = false
		role, err := user.FindRole(db, task.Assignee)
		if err != nil {
			return nil, err
		}
		if role == nil {
			reply.Err = &pb.Error{
				Code:        http.StatusInternalServerError,
				Message:     INTERNAL,
				Description: "角色不存在",
			}
			return reply, nil
		}
		err = camundamodel.SaveTask(db, &camundamodel.Task{
			Title:         in.Name,
			UserId:        in.UserId,
			Role:          role.ID,
			CurrentNode:   task.Name,
			CamundaTaskId: task.Id,
			InstanceId:    startProcessInstanceRes.Item.Id,
			EndFlag:       &endFlag,
		})
		if err != nil {
			return nil, err
		}
	}
	db.Commit()
	return reply, nil
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

package task

import (
	"fmt"
	"strconv"
	"userService/pkg/camunda"
	"userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/institution"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var topics = []string{
	"add_ins",
	"add_mtch",
	"del_ins",
	"del_mtch",
}

func RunServiceTask(format string, workerNum int) {
	ch := make(chan int, workerNum)
	ctx := context.TODO()
	for i := 0; i < workerNum; i++ {
		go finishRegister(ctx, i, ch)
	}
	c := cron.New()
	_ = c.AddFunc(format, func() {
		ch <- 1
	})
	c.Start()
}

func finishRegister(ctx context.Context, workerId int, ch <-chan int) {
	id := strconv.Itoa(workerId)
	for {
		<-ch
		client := camunda.Get()
		for _, topic := range topics {
			resp, err := client.ExternalTask.FetchAndLock(ctx, &pb.FetchAndLockExternalTaskReq{
				WorkerId: id,
				MaxTasks: 1,
				Topics: []*pb.Topic{
					{
						TopicName:    topic,
						LockDuration: 10000,
					},
				},
			})
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			if camunda.CheckError(resp) {
				logrus.Errorln(camunda.TransError(resp))
				continue
			}
			if len(resp.Item) == 0 {
				continue
			}
			switch topic {
			case "add_ins":
				// todo 机构注册
				err = institutionRegister(resp.Item[0])
			case "add_mtch":
				// todo 商户注册
				err = merchantRegister(resp.Item[0])
			case "del_ins":
				// todo 删除机构
			case "del_mtch":
				// todo 删除商户

			}
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			completeResp, err := client.ExternalTask.Complete(ctx, &pb.CompleteExternalTaskReq{
				Id: resp.Item[0].Id,
				Body: &pb.CompleteExternalTaskReqBody{
					WorkerId: id,
				},
			})
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			if camunda.CheckError(completeResp) {
				logrus.Errorln(camunda.TransError(completeResp))
				continue
			}
			logrus.Infoln("任务完成: ", resp.Item[0].Id, ", topic:", topic)
		}
	}
}

func institutionRegister(in *pb.FetchAndLockExternalTaskRespItem) error {
	db := common.DB
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceById(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// todo 查询机构信息
	info, err := institution.FindInstitutionInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("institution %s not found", instance.DataId)
	}

	// todo 修改状态
	// todo 入正式表
	return nil
}

func merchantRegister(in *pb.FetchAndLockExternalTaskRespItem) error {
	return nil
}

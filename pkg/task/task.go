package task

import (
	"fmt"
	"strconv"
	"userService/pkg/camunda"
	"userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/institution"

	"github.com/jinzhu/gorm"
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
		for _, t := range topics {
			func(topic string) {
				db := common.DB.Begin()
				defer db.Rollback()
				logrus.Debugln("拉取外部任务")
				resp, err := client.ExternalTask.FetchAndLock(ctx, &pb.FetchAndLockExternalTaskReq{
					WorkerId: id,
					MaxTasks: 1,
					Topics: []*pb.Topic{
						{
							TopicName:    topic,
							LockDuration: 100000,
						},
					},
				})
				if err != nil {
					logrus.Errorln(err)
					return
				}
				if camunda.CheckError(resp) {
					logrus.Errorln(camunda.TransError(resp))
					return
				}
				if len(resp.Item) == 0 {
					return
				}
				switch topic {
				case "add_ins":
					// 机构注册
					err = institutionRegister(db, resp.Item[0])
				case "add_mtch":
					// 商户注册
					err = merchantRegister(db, resp.Item[0])
				case "del_ins":
					err = deleteInstitution(db, resp.Item[0])
				case "del_mtch":
					// todo 删除商户

				}
				if err != nil {
					logrus.Errorln(err)
					return
				}
				completeResp, err := client.ExternalTask.Complete(ctx, &pb.CompleteExternalTaskReq{
					Id: resp.Item[0].Id,
					Body: &pb.CompleteExternalTaskReqBody{
						WorkerId: id,
					},
				})
				if err != nil {
					logrus.Errorln(err)
					return
				}
				if camunda.CheckError(completeResp) {
					logrus.Errorln(camunda.TransError(completeResp))
					return
				}
				db.Commit()
				logrus.Infoln("任务完成: ", resp.Item[0].Id, ", topic:", topic)

			}(t)
		}
	}
}

func institutionRegister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceById(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询机构信息
	info, err := institution.FindInstitutionInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("institution %s not found", instance.DataId)
	}
	// 查询fee，cash，control
	fees, err := institution.FindInstitutionFee(db, &institution.Fee{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}
	cashes, err := institution.FindInstitutionCash(db, &institution.Cash{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}
	controls, err := institution.FindInstitutionControl(db, &institution.Control{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}

	// 修改机构状态
	err = institution.UpdateInstitution(db, &institution.InstitutionInfo{
		InsIdCd: info.InsIdCd,
	}, &institution.InstitutionInfo{
		InsSta: "1",
	})
	if err != nil {
		return err
	}

	// 入正式表 institution fee cash control
	info.InsSta = "1"

	err = institution.SaveInstitutionMain(db, &institution.InstitutionInfoMain{
		InstitutionInfo: *info,
	})
	if err != nil {
		return err
	}
	for _, fee := range fees {
		err = institution.SaveInstitutionFeeMain(db, &institution.FeeMain{
			Fee: *fee,
		})
		if err != nil {
			return err
		}
	}
	for _, cash := range cashes {
		err = institution.SaveInstitutionCashMain(db, &institution.CashMain{
			Cash: *cash,
		})
		if err != nil {
			return err
		}
	}
	for _, control := range controls {
		err = institution.SaveInstitutionControlMain(db, &institution.ControlMain{
			Control: *control,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func merchantRegister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	return nil
}

func deleteInstitution(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceById(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}
	// todo 删除 institution fee cash control
	err = institution.DeleteInstitution(db, &institution.InstitutionInfo{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionFee(db, &institution.Fee{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionCash(db, &institution.Cash{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionControl(db, &institution.Control{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	return nil
}

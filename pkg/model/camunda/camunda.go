package camunda

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProcessDefinition struct {
	Id        string    `gorm:"column:process_def_id;primary_key"`
	Name      string    `gorm:"column:name;index"`
	Workflow  string    `gorm:"column:workflow"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (p ProcessDefinition) TableName() string {
	return "TBL_CAMUNDA_PROCESS_DEFINITION"
}

func QueryProcessDefinition(db *gorm.DB, query *ProcessDefinition) ([]*ProcessDefinition, error) {
	out := make([]*ProcessDefinition, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}

type ProcessInstance struct {
	InstanceId        int64  `gorm:"column:instance_id;primary_key"`
	CamundaInstanceId string `gorm:"column:camunda_instance_id;unique"`
	Title             string `gorm:"column:title"`
	DataId            string `gorm:"column:data_id"`
	UserId            int64
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}

func (p ProcessInstance) TableName() string {
	return "TBL_CAMUNDA_PROCESS_INSTANCE"
}

func FindProcessInstanceById(db *gorm.DB, id int64) (*ProcessInstance, error) {
	out := new(ProcessInstance)
	err := db.Take(out, &ProcessInstance{InstanceId: id}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}

func FindProcessInstanceByCamundaInstanceId(db *gorm.DB, id string) (*ProcessInstance, error) {
	out := new(ProcessInstance)
	err := db.Take(out, &ProcessInstance{CamundaInstanceId: id}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}

func SaveProcessInstance(db *gorm.DB, instance *ProcessInstance) error {
	return db.Create(instance).Error
}

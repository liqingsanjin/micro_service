package camunda

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	Id            int64     `gorm:"column:id;primary_key"`
	Title         string    `gorm:"column:title"`
	UserId        string    `gorm:"column:user_id"`
	CurrentNode   string    `gorm:"column:current_node"`
	CamundaTaskId string    `gorm:"column:camunda_task_id"`
	InstanceId    string    `gorm:"column:instance_id"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (t Task) TableName() string {
	return "TBL_CAMUNDA_TASK"
}

func QueryTask(db *gorm.DB, query *Task) ([]*Task, error) {
	out := make([]*Task, 0)
	err := db.Where(query).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}
	return out, err
}

func SaveTask(db *gorm.DB, task *Task) error {
	return db.Save(task).Error
}

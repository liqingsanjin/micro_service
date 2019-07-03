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

func QueryTask(db *gorm.DB, query *Task, page int32, size int32) ([]*Task, int32, error) {
	out := make([]*Task, 0)
	var count int32
	db.Model(&Task{}).Where(query).Count(&count)
	err := db.Where(query).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, count, nil
	}
	return out, count, err
}

func SaveTask(db *gorm.DB, task *Task) error {
	return db.Save(task).Error
}

func UpdateTask(db *gorm.DB, query *Task, task *Task) error {
	return db.Model(task).Where(query).Update(task).Error
}

func FindTaskById(db *gorm.DB, id int64) (*Task, error) {
	out := new(Task)
	err := db.Where("id = ?", id).Take(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}

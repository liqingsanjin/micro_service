package camunda

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProcessDefinition struct {
	Id        string    `gorm:"column:process_def_id;primary_key"`
	Name      string    `gorm:"column:name;index"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (p ProcessDefinition) TableName() string {
	return "TBL_CAMUNDA_PROCESS_DEFINITION"
}

func QueryProcessDefinition(db *gorm.DB, query *ProcessDefinition) ([]*ProcessDefinition, error) {
	out := make([]*ProcessDefinition, 0)
	err := db.Where(query).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}
	return out, err
}

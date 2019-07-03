package camunda

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Remark struct {
	RemarkId  int64     `gorm:"column:remark_id;primary_key"`
	Comment   string    `gorm:"column:comment"`
	TaskId    int64     `gorm:"column:task_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Remark) TableName() string {
	return "TBL_CAMUNDA_REMARK"
}

func SaveRemark(db *gorm.DB, remark *Remark) error {
	return db.Create(remark).Error
}

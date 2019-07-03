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

func QueryRemark(db *gorm.DB, query *Remark, page int32, size int32) ([]*Remark, int32, error) {
	out := make([]*Remark, 0)
	var count int32
	db.Model(query).Where(query).Count(&count)
	err := db.Where(query).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, count, nil
	}
	return out, count, err
}

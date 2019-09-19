package institution

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Group struct {
	GroupId   int64     `gorm:"column:group_id;primary_key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created"`
	UpdatedAt time.Time `gorm:"column:updated"`
}

func (Group) TableName() string {
	return "TBL_INS_GROUP"
}

func ListGroups(db *gorm.DB, page int32, size int32) ([]*Group, int32, error) {
	gs := make([]*Group, 0)
	var count int32
	db.Model(&Group{}).Count(&count)
	err := db.Offset((page - 1) * size).Limit(size).Find(&gs).Error
	if err == gorm.ErrRecordNotFound {
		return gs, count, nil
	}
	return gs, count, err
}

func SaveGroup(db *gorm.DB, data *Group) error {
	return db.Save(data).Error
}

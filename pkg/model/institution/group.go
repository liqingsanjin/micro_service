package institution

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Group struct {
	GroupId   int64     `gorm:"column:group_id;primary_key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type InsGroupBind struct {
	GroupId   int64     `gorm:"column:group_id;primary_key"`
	InsIdCd   string    `gorm:"column:ins_id_cd;primary_key"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Group) TableName() string {
	return "TBL_INS_GROUP"
}

func (InsGroupBind) TableName() string {
	return "TBL_INS_GROUP_BIND"
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

func ListInsGroupBind(db *gorm.DB, groupId int64) ([]*InsGroupBind, error) {
	out := make([]*InsGroupBind, 0)
	return out, db.Where(&InsGroupBind{GroupId: groupId}).Find(&out).Error
}

func SaveInsGroupBind(db *gorm.DB, data *InsGroupBind) error {
	return db.Save(data).Error
}

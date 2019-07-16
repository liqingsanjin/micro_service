package merchant

import "github.com/jinzhu/gorm"

type Picture struct {
}

type PictureMain struct {
	Picture
}

func (Picture) TableName() string {
	return "TBL_EDIT_MCHT_PICTURE"
}

func (PictureMain) TableName() string {
	return "TBL_MCHT_PICTURE"
}

func SavePicture(db *gorm.DB, data *Picture) error {
	return db.Create(data).Error
}

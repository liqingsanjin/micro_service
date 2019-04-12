package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type User struct {
	UserID    int64  `gorm:"column:USER_ID;primary_key"`
	LeaguerNO string `gorm:"column:LEAGUER_NO"`
	UserName  string `gorm:"column:USER_NAME;unique"`
}

func FindUserByUserName(db *gorm.DB, userName string) (*User, error) {
	user := &User{
		UserName: userName,
	}
	err := db.Where(&User{UserName: userName}).First(user).Error
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (u User) TableName() string {
	return "TBL_USER"
}

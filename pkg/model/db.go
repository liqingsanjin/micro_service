package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Options struct {
	User     string
	Password string
	DB       string
	Addr     string
}

func NewDB(opt *Options) (*gorm.DB, error) {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			opt.User,
			opt.Password,
			opt.Addr,
			opt.DB,
		),
	)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}

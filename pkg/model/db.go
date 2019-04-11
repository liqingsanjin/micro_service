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
	return gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			opt.User,
			opt.Password,
			opt.Addr,
			opt.DB,
		),
	)
}

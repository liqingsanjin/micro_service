package rabc

import (
	"fmt"

	"github.com/casbin/casbin"

	"userService/pkg/gormadapter"
	"userService/pkg/model"
)

func NewCasbin(fileName string, options *model.Options) *casbin.Enforcer {
	adapter := gormadapter.NewAdapter(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/", options.User, options.Password, options.Addr),
	)
	e := casbin.NewEnforcer(fileName, adapter)
	return e
}

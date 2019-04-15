package rabc

import "github.com/casbin/casbin"

// todo
func Casbin() *casbin.Enforcer {
	e := casbin.NewEnforcer()
	e.AddPolicy()
	return e
}

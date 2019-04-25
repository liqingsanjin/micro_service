package gateway

import "github.com/gin-gonic/gin"

var (
	engin = gin.New()
)

type service struct {
	handlers map[string]gin.HandlerFunc
	name     string
}

func newService(name string) *service {
	return &service{
		handlers: make(map[string]gin.HandlerFunc),
		name:     name,
	}
}

func (s *service) register(method string, handler gin.HandlerFunc) {

}

func (s *service) getHandler(method string) gin.HandlerFunc {
	return s.handlers[method]
}

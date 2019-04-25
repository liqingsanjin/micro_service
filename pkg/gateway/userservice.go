package gateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var userservice = newService("userservice")

func init() {
	engin.POST("/userservice/:method", NewUserServiceHandler())
}

func NewUserServiceHandler() gin.HandlerFunc {
	mu := http.NewServeMux()
	mu.Handle()
	return gin.WrapH(mu)
}

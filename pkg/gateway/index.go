package gateway

import (
	"net/http"
	"userService/pkg/userservice"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(endpoints *userservice.UserEndpoints) http.Handler {
	engine := gin.New()
	RegisterUserHandler(engine, endpoints)
	return engine
}

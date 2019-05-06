package gateway

import (
	"context"
	"net/http"
	"userService/pkg/pb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHttpHandler(endpoints *UserEndpoints, staticEndpoints *StaticEndpoints, institutionEndpoints *InstitutionEndpoints) http.Handler {
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	RegisterUserHandler(engine, endpoints)
	RegisterStaticHandler(engine, staticEndpoints)
	RegisterInstitutionHandler(engine, institutionEndpoints)
	return engine
}

func encodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

type StatusError interface {
	GetErr() *pb.Error
}

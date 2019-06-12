package gateway

import (
	"context"
	"net/http"
	"userService/pkg/pb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewHttpHandler(endpoints *UserEndpoints, staticEndpoints *StaticEndpoints, institutionEndpoints *InstitutionEndpoints, merchantEndpoints *MerchantEndpoint) http.Handler {
	engine := gin.New()
	engine.Use(logRequest)
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	RegisterUserHandler(engine, endpoints)
	RegisterStaticHandler(engine, staticEndpoints)
	RegisterInstitutionHandler(engine, institutionEndpoints)
	RegisterMerchantHandler(engine, merchantEndpoints)
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

func logRequest(c *gin.Context) {
	logrus.Debugln(c.Request.URL.String())
	c.Next()
}

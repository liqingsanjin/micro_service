package gateway

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(endpoints *UserEndpoints) http.Handler {
	engine := gin.New()
	RegisterUserHandler(engine, endpoints)
	return engine
}

func encodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

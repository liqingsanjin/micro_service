package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterProductHandler(engine *gin.Engine, endpoints *ProductEndpoints) {
	group := engine.Group("/product")

	group.POST("/listTransMap", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTransMapEndpoint,
		decodeHttpRequest(&pb.ListTransMapRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

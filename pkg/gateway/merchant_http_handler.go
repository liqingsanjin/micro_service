package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterMerchantHandler(engine *gin.Engine, endpoints *MerchantEndpoints) {
	group := engine.Group("/merchant")

	group.POST("/listMerchant", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListMerchantEndpoint,
		decodeHttpRequest(&pb.ListMerchantRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listGroup", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListGroupMerchantEndpoint,
		decodeHttpRequest(&pb.ListGroupMerchantRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

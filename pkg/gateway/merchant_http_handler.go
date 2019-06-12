package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterMerchantHandler(engine *gin.Engine, endpoints *MerchantEndpoint) {
	group := engine.Group("/merchant")

	group.POST("/listMerchant", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListMerchantEndpoint,
		decodeHttpRequest(&pb.ListMerchantRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

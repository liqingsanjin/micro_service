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

	group.POST("/saveMerchant", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantEndpoint,
		decodeHttpRequest(&pb.SaveMerchantRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveBankAccount", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantBankAccountEndpoint,
		decodeHttpRequest(&pb.SaveMerchantBankAccountRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveMerchantBizDeal", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantBizDealEndpoint,
		decodeHttpRequest(&pb.SaveMerchantBizDealRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

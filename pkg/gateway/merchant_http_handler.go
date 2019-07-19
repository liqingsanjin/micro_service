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

	group.POST("/saveMerchantBizFee", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantBizFeeEndpoint,
		decodeHttpRequest(&pb.SaveMerchantBizFeeRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveMerchantBusiness", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantBusinessEndpoint,
		decodeHttpRequest(&pb.SaveMerchantBusinessRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveMerchantPicture", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveMerchantPictureEndpoint,
		decodeHttpRequest(&pb.SaveMerchantPictureRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/getMerchantBankAccount", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetMerchantBankAccountEndpoint,
		decodeHttpRequest(&pb.GetMerchantBankAccountRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/getMerchantBizDeal", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetMerchantBizDealEndpoint,
		decodeHttpRequest(&pb.GetMerchantBizDealRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/getMerchantBizFee", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetMerchantBizFeeEndpoint,
		decodeHttpRequest(&pb.GetMerchantBizFeeRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/getMerchantBusiness", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetMerchantBusinessEndpoint,
		decodeHttpRequest(&pb.GetMerchantBusinessRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/getMerchantPicture", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetMerchantPictureEndpoint,
		decodeHttpRequest(&pb.GetMerchantPictureRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

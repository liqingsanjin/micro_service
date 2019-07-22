package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterStaticHandler(engine *gin.Engine, endpoints *StaticEndpoints) {
	engine.POST("/static/syncData", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SyncDataEndpoint,
		decodeHttpRequest(&pb.StaticSyncDataReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getDictionaryItem", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetDictionaryItemEndpoint,
		decodeHttpRequest(&pb.StaticGetDictionaryItemReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getDicByProdAndBiz", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetDicByProdAndBizEndpoint,
		decodeHttpRequest(&pb.StaticGetDicByProdAndBizReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getDicByInsCmpCd", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetDicByInsCmpCdEndpoint,
		decodeHttpRequest(&pb.StaticGetDicByInsCmpCdReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/checkValues", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.CheckValuesEndpoint,
		decodeHttpRequest(&pb.StaticCheckValuesReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getDictionaryLayerItem", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetDictionaryLayerItemEndpoint,
		decodeHttpRequest(&pb.GetDictionaryLayerItemReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getDictionaryItemByPk", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetDictionaryItemByPkEndpoint,
		decodeHttpRequest(&pb.GetDictionaryItemByPkReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/getUnionPayBankListByCode", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetUnionPayBankListByCodeEndpoint,
		decodeHttpRequest(&pb.GetUnionPayBankListByCodeRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/static/findUnionPayMccList", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.FindUnionPayMccListEndpoint,
		decodeHttpRequest(&pb.FindUnionPayMccListRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

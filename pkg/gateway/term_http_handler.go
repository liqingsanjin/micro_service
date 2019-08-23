package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterTermHandler(engine *gin.Engine, endpoints *TermEndpoints) {
	group := engine.Group("/term")

	group.POST("/listTermInfo", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTermInfoEndpoint,
		decodeHttpRequest(&pb.ListTermInfoRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveTerm", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveTermEndpoint,
		decodeHttpRequest(&pb.SaveTermRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveTermRisk", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveTermRiskEndpoint,
		decodeHttpRequest(&pb.SaveTermRiskRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listTermRisk", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTermRiskEndpoint,
		decodeHttpRequest(&pb.ListTermRiskRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveTermActivationState", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveTermActivationStateEndpoint,
		decodeHttpRequest(&pb.SaveTermActivationStateRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterWorkflowHandler(engine *gin.Engine, endpoints *WorkflowEndpoints) {
	group := engine.Group("/workflow")

	group.POST("/start", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.StartEndpoint,
		decodeHttpRequest(&pb.StartWorkflowRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listTask", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTaskEndpoint,
		decodeHttpRequest(&pb.ListTaskRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/handleTask", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.HandleTaskEndpoint,
		decodeHttpRequest(&pb.HandleTaskRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listRemark", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListRemarkEndpoint,
		decodeHttpRequest(&pb.ListRemarkRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

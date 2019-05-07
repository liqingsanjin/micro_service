package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterInstitutionHandler(engine *gin.Engine, endpoints *InstitutionEndpoints) {
	engine.POST("/institution/tnxHisDownload", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.TnxHisDownloadEndpoint,
		decodeHttpRequest(&pb.InstitutionTnxHisDownloadReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/institution/getTfrTrnLogs", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetTfrTrnLogsEndpoint,
		decodeHttpRequest(&pb.GetTfrTrnLogsReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/institution/getTfrTrnLog", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.GetTfrTrnLogEndpoint,
		decodeHttpRequest(&pb.GetTfrTrnLogReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/institution/downloadTfrTrnLogs", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.DownloadTfrTrnLogsEndpoint,
		decodeHttpRequest(&pb.DownloadTfrTrnLogsReq{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}

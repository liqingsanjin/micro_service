package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"userService/pkg/pb"
	"userService/pkg/userservice"

	"github.com/dgrijalva/jwt-go"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterUserHandler(engine *gin.Engine, endpoints *UserEndpoints) {
	engine.POST("/user/login", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.LoginEndpoint,
		decodeHttpLoginRequest,
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/register", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.RegisterEndpoint,
		decodeHttpRegisterRequest,
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/getPermissions",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.GetPermissionsEndpoint,
			decodeHttpGetPermissionsRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/checkPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CheckPermissionEndpoint,
			decodeHttpCheckPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRoutesEndpoint,
			decodeHttpAddRoutesRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/listRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.ListRoutesEndpoint,
			decodeHttpListRoutesRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/createPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CreatePermissionEndpoint,
			decodeHttpCreatePermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/updatePermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.UpdatePermissionEndpoint,
			decodeHttpUpdatePermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRouteForPermissionEndpoint,
			decodeHttpAddRouteForPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/removeRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.RemoveRouteForPermissionEndpoint,
			decodeHttpRemoveRouteForPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

}

func convertHttpHandlerToGinHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func decodeHttpRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.RegisterRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.LoginRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpGetPermissionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.GetPermissionsRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	userid, _ := strconv.Atoi(r.Form.Get("userid"))
	request.User = &pb.UserInfo{
		Username: r.Form.Get("username"),
		Userid:   int64(userid),
	}
	return &request, err
}

func decodeHttpCheckPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.CheckPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	userid, _ := strconv.Atoi(r.Form.Get("userid"))
	request.User = &pb.UserInfo{
		Username: r.Form.Get("username"),
		Userid:   int64(userid),
	}
	return &request, err
}

func decodeHttpAddRoutesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.AddRoutesRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpListRoutesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.ListRoutesRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpCreatePermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.CreatePermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpUpdatePermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.UpdatePermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpAddRouteForPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.AddRouteForPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpRemoveRouteForPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.RemoveRouteForPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func encodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	code, msg := err2codeAndMessage(err)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(gin.H{
		"error": msg,
	})
}

func err2codeAndMessage(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}
	switch e := err.(type) {
	case lb.RetryError:
		err = e.Final
	}
	s, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, err.Error()
	}
	code := s.Code()
	msg := s.Message()
	switch code {
	case codes.PermissionDenied:
		return http.StatusUnauthorized, msg
	case codes.Internal:
		return http.StatusInternalServerError, msg
	case codes.NotFound:
		return http.StatusBadRequest, msg
	case codes.AlreadyExists:
		return http.StatusBadRequest, msg
	case codes.InvalidArgument:
		return http.StatusBadRequest, msg
	}
	return http.StatusInternalServerError, msg
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(userservice.SignedKey), nil
}

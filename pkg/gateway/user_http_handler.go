package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"userService/pkg/pb"
	"userService/pkg/userservice"

	"github.com/dgrijalva/jwt-go"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc/metadata"
)

func RegisterUserHandler(engine *gin.Engine, endpoints *UserEndpoints) {
	engine.POST("/user/login", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.LoginEndpoint,
		decodeHttpRequest(&pb.LoginRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/getPermissions",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.GetPermissionsEndpoint,
			decodeHttpRequest(&pb.GetPermissionsRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
			httptransport.ServerBefore(setUserInfoContext),
		)))

	engine.POST("/user/checkPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CheckPermissionEndpoint,
			decodeHttpRequest(&pb.CheckPermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
			httptransport.ServerBefore(setUserInfoContext),
		)))

	engine.POST("/user/register", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.RegisterEndpoint,
		decodeHttpRequest(&pb.RegisterRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/addPermissionForRole",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddPermissionForRoleEndpoint,
			decodeHttpRequest(&pb.AddPermissionForRoleRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/createRole",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CreateRoleEndpoint,
			decodeHttpRequest(&pb.CreateRoleRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRoleForUser",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRoleForUserEndpoint,
			decodeHttpRequest(&pb.AddRoleForUserRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRoutesEndpoint,
			decodeHttpRequest(&pb.AddRoutesRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/listRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.ListRoutesEndpoint,
			decodeHttpRequest(&pb.ListRoutesRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/createPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CreatePermissionEndpoint,
			decodeHttpRequest(&pb.CreatePermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/updatePermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.UpdatePermissionEndpoint,
			decodeHttpRequest(&pb.UpdatePermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRouteForPermissionEndpoint,
			decodeHttpRequest(&pb.AddRouteForPermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/removeRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.RemoveRouteForPermissionEndpoint,
			decodeHttpRequest(&pb.RemoveRouteForPermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/removePermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.RemovePermissionEndpoint,
			decodeHttpRequest(&pb.RemovePermissionRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/listPermissions",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.ListPermissionsEndpoint,
			decodeHttpRequest(&pb.ListPermissionsRequest{}),
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

}

func convertHttpHandlerToGinHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func decodeHttpRequest(ins interface{}) httptransport.DecodeRequestFunc {
	tp := reflect.TypeOf(ins)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		request := reflect.New(tp).Interface()
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&request)
		return request, err
	}
}

func encodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s, ok := response.(StatusError)
	if ok {
		err := s.GetErr()
		if err != nil {
			w.WriteHeader(int(err.GetCode()))
			return json.NewEncoder(w).Encode(gin.H{
				"err":  err.GetMessage(),
				"desc": err.GetDescription(),
			})
		}
	}
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	code, msg := err2codeAndMessage(err)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(gin.H{
		"err":  msg,
		"desc": "未知错误",
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
	return http.StatusInternalServerError, err.Error()
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(userservice.SignedKey), nil
}

func setUserInfoContext(ctx context.Context, r *http.Request) context.Context {
	username := r.Form.Get("username")
	id := r.Form.Get("id")
	md := metadata.New(map[string]string{
		"username": username,
		"id":       id,
	})
	return context.WithValue(ctx, "userInfo", md)
}

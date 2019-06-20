package institutionservice

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

//Passport .
func Passport() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Println("into passport")
			ctx = context.WithValue(ctx, "result", "true")
			return next(ctx, req)
		}
	}
}

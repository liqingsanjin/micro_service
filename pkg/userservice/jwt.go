package userservice

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/metadata"
)

var (
	signedKey = "huiepay"
	jwtToken  = "jwtToken"
)

type ClaimsFactory func() jwt.Claims

type UserClaims struct {
	User *UserInfo
	jwt.StandardClaims
}

func genToken(user *UserInfo, exTime time.Time) (string, error) {
	claims := UserClaims{
		User: user,
	}
	claims.ExpiresAt = exTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signedKey))
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(signedKey), nil
}

func UserClaimFactory() jwt.Claims {
	return &UserClaims{}
}

func jwtParser(keyFunc jwt.Keyfunc, method jwt.SigningMethod, newClaims ClaimsFactory) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, ErrTokenContextMissing
			}

			tk := md.Get(jwtToken)
			if len(tk) == 0 || tk[0] == "" {
				return nil, ErrTokenContextMissing
			}

			tokenString := tk[0]

			token, err := jwt.ParseWithClaims(tokenString, newClaims(), func(token *jwt.Token) (interface{}, error) {
				if token.Method != method {
					return nil, ErrUnexpectedSigningMethod
				}

				return keyFunc(token)
			})
			if err != nil {
				if e, ok := err.(*jwt.ValidationError); ok {
					switch {
					case e.Errors&jwt.ValidationErrorMalformed != 0:
						return nil, ErrTokenMalformed
					case e.Errors&jwt.ValidationErrorExpired != 0:
						return nil, ErrTokenExpired
					case e.Errors&jwt.ValidationErrorNotValidYet != 0:
						return nil, ErrTokenNotActive
					case e.Inner != nil:
						return nil, e.Inner
					}
				}
				return nil, err
			}

			if !token.Valid {
				return nil, ErrTokenInvalid
			}

			claims, ok := token.Claims.(*UserClaims)
			if !ok || claims.User == nil {
				return nil, ErrTokenInvalid
			}

			ctx = context.WithValue(ctx, "userInfo", claims.User)

			return next(ctx, request)
		}
	}
}

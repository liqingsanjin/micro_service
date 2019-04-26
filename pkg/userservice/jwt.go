package userservice

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	SignedKey = "huiepay"
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
	return token.SignedString([]byte(SignedKey))
}

func UserClaimFactory() jwt.Claims {
	return &UserClaims{}
}

func JwtMiddleware(keyFunc jwt.Keyfunc, method jwt.SigningMethod, newClaims ClaimsFactory) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		token, err := jwt.ParseWithClaims(tokenString, newClaims(), func(token *jwt.Token) (interface{}, error) {
			if token.Method != method {
				return nil, ErrUnexpectedSigningMethod
			}

			return keyFunc(token)
		})
		if err != nil {
			logrus.Errorln(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			logrus.Errorln(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*UserClaims)
		if !ok || claims.User == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Request.Form = url.Values{}
		c.Request.Form.Set("username", claims.User.UserName)
		c.Request.Form.Set("userid", fmt.Sprintf("%d", claims.User.ID))
		c.Next()
	}
}

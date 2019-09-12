package middleware

import (
	"Microservice/backend/model"
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ErrMissingHeader 缺少Authorization
var errMissingHeader = errors.New("The length of the `Authorization` header is zero.")

// AuthMiddleware gin的验证登录拦截器
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := parseRequest(c); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "未登陆",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// parseRequest 解析 Header 中的Authorization
func parseRequest(c *gin.Context) (*model.User, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return nil, errMissingHeader
	}

	var t string
	return parse(t, "myproject")
}

func parse(tokenString string, secret string) (*model.User, error) {
	user := new(model.User)

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return user, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Name = claims["userName"].(string)
		user.Account = claims["account"].(string)
		user.Exp = claims["exp"].(int64)

		return user, nil
	} else {
		return user, err
	}
}

// jwt 框架使用的验证方法
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

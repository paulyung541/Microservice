package middleware

import (
	"Microservice/backend/constants"
	"Microservice/backend/model"
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	errMissingHeader = errors.New("The length of the `Authorization` header is zero.") // 缺少 token
	errTokenExpired  = errors.New("token had been expired.")                           // token 过期
)

// AuthMiddleware gin的验证登录拦截器
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := parseRequest(c); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "验证无效: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// parseRequest 解析 Header 中的Authorization
func parseRequest(c *gin.Context) (*model.User, error) {
	authorization := c.Request.Header.Get("Authorization")

	if len(authorization) == 0 {
		return nil, errMissingHeader
	}

	return parse(authorization, constants.JWTSecretString)
}

func parse(tokenString string, secret string) (*model.User, error) {
	user := new(model.User)

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return user, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Name = claims["userName"].(string)
		user.Account = claims["account"].(string)
		user.Exp = int64(claims["exp"].(float64))

		// 当前时间大于了过期时间，则token已过期，需要重新登录
		if time.Now().Unix() > user.Exp {
			return user, errTokenExpired
		}

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

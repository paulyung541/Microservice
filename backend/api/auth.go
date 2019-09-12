package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	authpb "Microservice/idls/outfile/auth"
)

const authAddress = "localhost:40001"

var authClient AuthClient

// AuthClient xxx
type AuthClient struct {
	authpb.AuthServiceClient
}

func init() {
	conn, err := grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc connect fail")
	}
	defer conn.Close()

	authClient = AuthClient{
		authpb.NewAuthServiceClient(conn),
	}
}

/**
curl -XPOST http://localhost:8080/auth/sign_up -d 'name=ysy&account=as@sina.com&password=123'
*/
func signUp(c *gin.Context) {
	// 请求auth微服务完成注册
	name := c.PostForm("name")
	account := c.PostForm("account")
	password := c.PostForm("password")

	returnJSON := make(map[string]interface{})

	log.Println(name, account, password)
	resp, err := authClient.SignUp(context.Background(),
		&authpb.SignUpRequest{Name: name, Account: account, Password: password})
	if resp != nil {
		returnJSON["success"] = resp.Success
		returnJSON["msg"] = resp.Msg
	}
	if err != nil {
		log.Println("call auth SignUp fail", err.Error())
		returnJSON["success"] = "false"
		returnJSON["msg"] = err.Error()
	}

	c.JSON(200, returnJSON)
}

// curl -XPOST http://localhost:8080/auth/login -d'account=as@sina.com&password=123'
func login(c *gin.Context) {
	//请求auth微服务完成登录

	account := c.PostForm("account")
	password := c.PostForm("password")

	returnJSON := make(map[string]interface{})
	resp, err := authClient.Login(context.Background(), &authpb.LoginRequest{Account: account, Password: password})
	if resp != nil {
		returnJSON["success"] = resp.Success
		returnJSON["msg"] = "登录成功"
		returnJSON["token"] = resp.Token
	}

	if err != nil {
		log.Println("call auth login fail", err.Error())
		returnJSON["success"] = "false"
		returnJSON["msg"] = err.Error()
	}

	c.JSON(200, returnJSON)
}

func classList(c *gin.Context) {

	c.JSON(200, gin.H{
		"success": true,
	})
}

func addClass(c *gin.Context) {

}

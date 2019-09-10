package main

/*
	curl -XPOST http://localhost:8080/auth/sign_up

*/

import (
	"context"
	"log"
	"net/http"
	"time"

	authpb "Microservice/idls/outfile/auth"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AuthClient struct {
	authpb.AuthServiceClient
}

const address = "localhost:40001"

var authClient AuthClient

func main() {
	router := gin.Default()

	router.HEAD("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", signUp)
		auth.POST("/login", login)
	}

	school := router.Group("/school")
	{
		school.GET("/class", classList)
		school.POST("/class", addClass)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc connect fail")
	}
	defer conn.Close()

	authClient = AuthClient{
		authpb.NewAuthServiceClient(conn),
	}

	go func() {
		pingServer()
	}()

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// 服务器健康检查
func pingServer() {
	i := 0
	for {
		_, err := http.Head("http://localhost:8080/ping")
		if err != nil {
			i++
			if i > 3 {
				log.Fatal("ping fail, stop server!!!")
			}
		} else {
			i = 0
		}
		time.Sleep(3 * time.Second)
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

}

func addClass(c *gin.Context) {

}

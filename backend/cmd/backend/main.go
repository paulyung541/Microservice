package main

import (
	"Microservice/backend/api"

	"log"
	"net/http"
	"time"
)

const port = ":8080"

func main() {
	router := api.InitAPI()

	// go func() {
	// 	pingServer()
	// }()

	router.Run(port) // listen and serve on 0.0.0.0:8080
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

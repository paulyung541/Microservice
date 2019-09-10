package main

import (
	"Microservice/auth/api"
	"Microservice/auth/config"
	pb "Microservice/idls/outfile/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":40001"

func main() {
	config.InitDb()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("err listen")
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, new(api.Serv))

	s.Serve(listen)
}

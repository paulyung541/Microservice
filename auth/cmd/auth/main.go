package main

import (
	"github.com/sirupsen/logrus"
	"Microservice/auth/api"
	pb "Microservice/idls/outfile/auth"
	"fmt"
	"log"
	"net"

	"github.com/paulyung541/jotnar"
	"google.golang.org/grpc"

	grpc_mw "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

const port = ":40001"

func main() {
	jotnar.New().
		InitConfigViperToml().
		InitMysql().
		InitLogger().
		Init(func() {
			fmt.Println("auth 启动成功")
		})

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("err listen")
	}
	defer listen.Close()

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_mw.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(jotnar.GetLogger())),
	)))

	pb.RegisterAuthServiceServer(s, new(api.Serv))

	s.Serve(listen)
}

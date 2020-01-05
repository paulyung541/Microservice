module Microservice/auth

go 1.12

require (
	Microservice/idls v0.0.0
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/paulyung541/jotnar v0.0.0-20191224085714-bb6dcbaf7048
	github.com/sirupsen/logrus v1.4.2

	golang.org/x/net v0.0.0
	google.golang.org/grpc v1.26.0
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297
	google.golang.org/grpc => ../../grpc-go
)

replace Microservice/idls => ../idls

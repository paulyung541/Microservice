module Microservice/auth

go 1.12

require (
	Microservice/idls v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/pkg/errors v0.8.0

	golang.org/x/net v0.0.0
	google.golang.org/grpc v1.23.0
)

replace golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297

replace Microservice/idls => ../idls

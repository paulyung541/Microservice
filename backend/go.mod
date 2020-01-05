module Microservice/backend

go 1.12

require (
	Microservice/idls v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	golang.org/x/net v0.0.0 // indirect
	golang.org/x/sys v0.0.0 // indirect
	google.golang.org/grpc v1.26.0
)

replace (
	Microservice/idls => ../idls
	golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190904005037-43c01164e931
	google.golang.org/grpc => ../../grpc-go
)

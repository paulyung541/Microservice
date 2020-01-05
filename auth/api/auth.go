package api

import (
	"github.com/paulyung541/jotnar"
	"time"

	"Microservice/auth/constants"
	"Microservice/auth/model"
	"Microservice/auth/services"
	pb "Microservice/idls/outfile/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

// Serv xxx
type Serv struct {
	services.User
}

// SignUp xxx
func (s *Serv) SignUp(c context.Context, in *pb.SignUpRequest) (*pb.SignUpReply, error) {
	jotnar.GetLogger().Infof("name = %s, account = %s\n", in.Name, in.Account)
	err := s.AddUser(&model.User{Name: in.Name, Account: in.Account, Password: in.Password})
	if err != nil {
		jotnar.GetLogger().Error("注册失败", err.Error())
		return &pb.SignUpReply{Success: "false", Msg: "注册失败"}, err
	}
	return &pb.SignUpReply{Success: "true", Msg: "注册成功"}, nil
}

// Login xxx
func (s *Serv) Login(c context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.GetUser(in.Account, in.Password)
	if err != nil {
		return &pb.LoginReply{Success: "false"}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": user.Name,
		"account":  user.Account,
		"exp":      time.Now().Add(1 * time.Hour).Unix(), // 1小时过期
	})

	tokenString, err := token.SignedString([]byte(constants.JWTSecretString))
	if err != nil {
		return &pb.LoginReply{Success: "false"}, err
	}

	jotnar.GetLogger().Info("登录成功")
	return &pb.LoginReply{Success: "true", Token: tokenString}, nil
}

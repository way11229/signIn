package grpcHandler

import (
	"context"
	"errors"
	"signIn/gateway/domain"
	pb "signIn/gateway/gen/gateway"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedGatewayServer

	SignInService domain.SignInService
}

func New(s *grpcLib.Server, serviceList domain.ServiceList) {
	handler := &GrpcHandler{
		SignInService: serviceList.SignInService,
	}

	pb.RegisterGatewayServer(s, handler)
}

func (g *GrpcHandler) SignIn(c context.Context, inputData *pb.SignInData) (*pb.Response, error) {
	var err error
	var signInData domain.SignInData

	rtn := pb.Response{}
	if (inputData.Method == "") || (inputData.VerifyCode == "") {
		return &rtn, errors.New("missing paramters")
	}

	accessData := domain.AccessData{
		Token: inputData.VerifyCode,
		Extra: inputData.Extra,
	}

	switch inputData.Method {
	case domain.LINE_METHOD:
		signInData, err = g.SignInService.SignInWithLine(c, accessData)
	case domain.FB_METHOD:
		signInData, err = g.SignInService.SignInWithFb(c, accessData)
	case domain.GOOGLE_METHOD:
		signInData, err = g.SignInService.SignInWithGoogle(c, accessData)
	}

	if err != nil {
		return &rtn, errors.New("Sign In error: method:" + inputData.Method + ",error:" + err.Error())
	}

	rtn = pb.Response{
		Id:    signInData.ID,
		Name:  signInData.Name,
		Email: signInData.Email,
		Phone: signInData.Phone,
		Extra: signInData.Extra,
	}

	return &rtn, nil
}

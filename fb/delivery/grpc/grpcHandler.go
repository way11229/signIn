package grpcHandler

import (
	"context"
	"errors"
	"signIn/fb/domain"
	pb "signIn/fb/gen/fb"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedFbServer

	SignInService domain.SignInService
}

func New(s *grpcLib.Server, serviceList domain.ServiceList) {
	handler := &GrpcHandler{
		SignInService: serviceList.SignInService,
	}

	pb.RegisterFbServer(s, handler)
}

func (g *GrpcHandler) SignIn(cxt context.Context, signInData *pb.SignInData) (*pb.FbResponse, error) {
	rtn := pb.FbResponse{}
	if signInData.VerifyCode == "" {
		rtn.Error = "Parameters is empty"
		return &rtn, errors.New("parameters is empty")
	}

	signInResponse, err := g.SignInService.SignIn(signInData.VerifyCode)
	if err != nil {
		rtn.Error = err.Error()
		return &rtn, err
	}

	rtn = pb.FbResponse{
		AccessToken:         signInResponse.AccessToken,
		AccessTokenExpireIn: signInResponse.AccessTokenExpireIn,
		UserId:              signInResponse.UserId,
		Name:                signInResponse.Name,
		Picture:             signInResponse.Picture,
		Email:               signInResponse.Email,
		Birthday:            signInResponse.Birthday,
		Error:               "",
	}

	return &rtn, nil
}

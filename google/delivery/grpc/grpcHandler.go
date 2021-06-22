package grpcHandler

import (
	"context"
	"errors"
	"signIn/google/domain"
	pb "signIn/google/gen/google"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedGoogleServer

	SignInService domain.SignInService
}

func New(s *grpcLib.Server, serviceList domain.ServiceList) {
	handler := &GrpcHandler{
		SignInService: serviceList.SignInService,
	}

	pb.RegisterGoogleServer(s, handler)
}

func (g *GrpcHandler) SignIn(cxt context.Context, signInData *pb.SignInData) (*pb.GoogleResponse, error) {
	rtn := pb.GoogleResponse{}
	if signInData.VerifyCode == "" {
		rtn.Error = "Verify code is empty"
		return &rtn, errors.New("verify code is empty")
	}

	signInResponse, err := g.SignInService.SignIn(signInData.VerifyCode)
	if err != nil {
		rtn.Error = err.Error()
		return &rtn, err
	}

	rtn = pb.GoogleResponse{
		AccessToken:         signInResponse.AccessToken,
		AccessTokenExpireIn: signInResponse.AccessTokenExpireIn,
		RefreshToken:        signInResponse.RefreshToken,
		UserId:              signInResponse.UserId,
		Name:                signInResponse.Name,
		Picture:             signInResponse.Picture,
		Email:               signInResponse.Email,
		Error:               "",
	}

	return &rtn, nil
}

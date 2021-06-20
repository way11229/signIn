package grpcHandler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "signIn/fb/gen/fb"

	grpcHandler "signIn/fb/delivery/grpc"
	"signIn/fb/domain"
	"signIn/fb/domain/mocks"
)

func TestSignIn(t *testing.T) {
	ctx := context.Background()
	fbSignInData := &pb.SignInData{
		VerifyCode: "123456789",
	}

	t.Run("success", func(t *testing.T) {
		expectSignInResponse := domain.SignInResponse{
			AccessToken:         "123456789",
			AccessTokenExpireIn: 2592000,
			UserId:              "123456789",
			Name:                "test",
			Picture:             "",
			Email:               "test@test.com",
			Birthday:            "2021-01-01",
		}

		expectGrpcResponse := pb.FbResponse{
			AccessToken:         expectSignInResponse.AccessToken,
			AccessTokenExpireIn: expectSignInResponse.AccessTokenExpireIn,
			UserId:              expectSignInResponse.UserId,
			Name:                expectSignInResponse.Name,
			Picture:             expectSignInResponse.Picture,
			Email:               expectSignInResponse.Email,
			Birthday:            expectSignInResponse.Birthday,
			Error:               "",
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignIn",
			fbSignInData.VerifyCode,
		).Return(expectSignInResponse, nil).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, fbSignInData)

		assert.NoError(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})

	t.Run("fail", func(t *testing.T) {
		expectGrpcResponse := pb.FbResponse{
			AccessToken:         "",
			AccessTokenExpireIn: 0,
			UserId:              "",
			Name:                "",
			Picture:             "",
			Email:               "",
			Birthday:            "",
			Error:               "invalid access token",
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignIn",
			fbSignInData.VerifyCode,
		).Return(domain.SignInResponse{}, errors.New("invalid access token")).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, fbSignInData)

		assert.Error(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})
}

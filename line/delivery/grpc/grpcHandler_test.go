package grpcHandler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "signIn/line/gen/line"

	grpcHandler "signIn/line/delivery/grpc"
	"signIn/line/domain"
	"signIn/line/domain/mocks"
)

func TestSignIn(t *testing.T) {
	ctx := context.Background()
	lineSignInData := &pb.SignInData{
		VerifyCode: "7G7ovtjlalaDzWtUVO2",
	}

	t.Run("success", func(t *testing.T) {
		expectSignInResponse := domain.SignInResponse{
			AccessToken:         "123456789",
			AccessTokenExpireIn: 2592000,
			RefreshToken:        "123456789",
			UserId:              "123456789",
			Name:                "test",
			Picture:             "",
			Email:               "test@test.com",
			StatusMessage:       "Just to eat",
		}

		expectGrpcResponse := pb.LineResponse{
			AccessToken:         expectSignInResponse.AccessToken,
			AccessTokenExpireIn: expectSignInResponse.AccessTokenExpireIn,
			RefreshToken:        expectSignInResponse.RefreshToken,
			UserId:              expectSignInResponse.UserId,
			Name:                expectSignInResponse.Name,
			Picture:             expectSignInResponse.Picture,
			Email:               expectSignInResponse.Email,
			StatusMessage:       expectSignInResponse.StatusMessage,
			Error:               "",
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignIn",
			"7G7ovtjlalaDzWtUVO2",
		).Return(expectSignInResponse, nil).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, lineSignInData)

		assert.NoError(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})

	t.Run("fail", func(t *testing.T) {
		expectGrpcResponse := pb.LineResponse{
			AccessToken:         "",
			AccessTokenExpireIn: 0,
			RefreshToken:        "",
			UserId:              "",
			Name:                "",
			Picture:             "",
			Email:               "",
			StatusMessage:       "",
			Error:               "invalid verify code",
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignIn",
			"7G7ovtjlalaDzWtUVO2",
		).Return(domain.SignInResponse{}, errors.New("invalid verify code")).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, lineSignInData)

		assert.Error(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})
}

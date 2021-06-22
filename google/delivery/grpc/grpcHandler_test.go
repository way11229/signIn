package grpcHandler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "signIn/google/gen/google"

	grpcHandler "signIn/google/delivery/grpc"
	"signIn/google/domain"
	"signIn/google/domain/mocks"
)

func TestSignIn(t *testing.T) {
	ctx := context.Background()
	googleSignInData := &pb.SignInData{
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
		}

		expectGrpcResponse := pb.GoogleResponse{
			AccessToken:         expectSignInResponse.AccessToken,
			AccessTokenExpireIn: expectSignInResponse.AccessTokenExpireIn,
			RefreshToken:        expectSignInResponse.RefreshToken,
			UserId:              expectSignInResponse.UserId,
			Name:                expectSignInResponse.Name,
			Picture:             expectSignInResponse.Picture,
			Email:               expectSignInResponse.Email,
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

		grpcRtn, err := handler.SignIn(ctx, googleSignInData)

		assert.NoError(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})

	t.Run("fail", func(t *testing.T) {
		expectGrpcResponse := pb.GoogleResponse{
			AccessToken:         "",
			AccessTokenExpireIn: 0,
			RefreshToken:        "",
			UserId:              "",
			Name:                "",
			Picture:             "",
			Email:               "",
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

		grpcRtn, err := handler.SignIn(ctx, googleSignInData)

		assert.Error(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})
}

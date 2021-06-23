package grpcHandler_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "signIn/gateway/gen/gateway"

	grpcHandler "signIn/gateway/delivery/grpc"
	"signIn/gateway/domain"
	"signIn/gateway/domain/mocks"
)

func TestSignIn(t *testing.T) {
	ctx := context.Background()
	inputData := &pb.SignInData{
		Method:     "line",
		VerifyCode: "7G7ovtjlalaDzWtUVO2",
		Extra:      "",
	}

	accessData := domain.AccessData{
		Token: inputData.VerifyCode,
		Extra: "",
	}

	t.Run("success", func(t *testing.T) {
		jsonExtra, _ := json.Marshal(domain.LineResponse{})
		expectSignInResponse := domain.SignInData{
			ID:    "123456789",
			Name:  "Way",
			Email: "test@test.com",
			Phone: "",
			Extra: string(jsonExtra),
		}

		expectGrpcResponse := pb.Response{
			Id:    expectSignInResponse.ID,
			Name:  expectSignInResponse.Name,
			Email: expectSignInResponse.Email,
			Phone: expectSignInResponse.Phone,
			Extra: expectSignInResponse.Extra,
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignInWithLine",
			ctx,
			accessData,
		).Return(expectSignInResponse, nil).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, inputData)

		assert.NoError(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})

	t.Run("fail", func(t *testing.T) {
		expectGrpcResponse := pb.Response{
			Id:    "",
			Name:  "",
			Email: "",
			Phone: "",
			Extra: "",
		}

		mockSignInService := new(mocks.SignInService)
		mockSignInService.On(
			"SignInWithLine",
			ctx,
			accessData,
		).Return(domain.SignInData{}, errors.New("invalid verify code")).Once()

		handler := grpcHandler.GrpcHandler{
			SignInService: mockSignInService,
		}

		grpcRtn, err := handler.SignIn(ctx, inputData)

		assert.Error(t, err)
		assert.Equal(t, &expectGrpcResponse, grpcRtn)
	})
}

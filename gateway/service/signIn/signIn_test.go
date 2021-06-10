package signInService_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"signIn/gateway/domain"
	"signIn/gateway/domain/mocks"
	signInService "signIn/gateway/service/signIn"
)

func TestSignInWithLine(t *testing.T) {
	ctx := context.Background()
	accessData := domain.AccessData{
		Token: "7G7ovtjlalaCDzWtUVO2",
		Extra: "",
	}

	expectLineResponse := domain.LineResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		RefreshToken:        "123456789",
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
	}
	extraJson, _ := json.Marshal(expectLineResponse)

	expectRtn := domain.SignInData{
		ID:    expectLineResponse.UserId,
		Name:  expectLineResponse.Name,
		Email: expectLineResponse.Email,
		Phone: "",
		Extra: string(extraJson),
	}

	t.Run("success", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)

		mockLineRepo.On(
			"SendSignInRequest",
			mock.MatchedBy(
				func(ctx context.Context) bool {
					return true
				},
			),
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(expectLineResponse, nil).Once()

		s := signInService.New(mockLineRepo)
		signInData, err := s.SignInWithLine(ctx, accessData)

		assert.NoError(t, err)
		assert.Equal(t, expectRtn, signInData)

		mockLineRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)

		mockLineRepo.On(
			"SendSignInRequest",
			mock.MatchedBy(
				func(ctx context.Context) bool {
					return true
				},
			),
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(domain.LineResponse{}, errors.New("Line sign in Error")).Once()

		s := signInService.New(mockLineRepo)
		signInData, err := s.SignInWithLine(ctx, accessData)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInData{}, signInData)

		mockLineRepo.AssertExpectations(t)
	})
}

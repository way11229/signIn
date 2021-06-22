package signinService_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"signIn/google/domain"
	"signIn/google/domain/mocks"
	signInService "signIn/google/service/signIn"
)

func TestSignIn(t *testing.T) {
	verifyCode := "7G7ovtjlalaCDzWtUVO2"

	expectSignInResponse := domain.SignInResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		RefreshToken:        "123456789",
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
	}

	getAccessTokenResponse := domain.GetAccessTokenResponse{
		AccessToken:      "123456789",
		ExpiresIn:        2592000,
		IdToken:          "1234567890",
		RefreshToken:     "123456789",
		Scope:            "email&openid",
		TokenType:        "",
		Error:            "",
		ErrorDescription: "",
	}

	getUserInfoResponse := domain.GetUserInfoResponse{
		Sub:              "123456789",
		Name:             "test",
		Picture:          "",
		Email:            "test@test.com",
		Error:            "",
		ErrorDescription: "",
	}

	mockGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)
	mockGetUserInfoRepo := new(mocks.GetUserInfoRepository)

	mockGetAccessTokenRepo.On(
		"GetAccessToken",
		"7G7ovtjlalaCDzWtUVO2",
	).Return(getAccessTokenResponse, nil).Times(3)

	mockGetUserInfoRepo.On(
		"GetUserInfo",
		"123456789",
	).Return(getUserInfoResponse, nil).Once()

	t.Run("success", func(t *testing.T) {
		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			GetUserInfoRepository:    mockGetUserInfoRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.NoError(t, err)
		assert.Equal(t, expectSignInResponse, signInResponse)
	})

	t.Run("getAccessTokenfail", func(t *testing.T) {
		mockFailGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)

		getAccessTokenResponseErr := domain.GetAccessTokenResponse{
			AccessToken:      "",
			ExpiresIn:        0,
			IdToken:          "",
			RefreshToken:     "",
			Scope:            "",
			TokenType:        "",
			Error:            "invalid code",
			ErrorDescription: "invalid verify code",
		}

		mockFailGetAccessTokenRepo.On(
			"GetAccessToken",
			"7G7ovtjlalaCDzWtUVO2",
		).Return(getAccessTokenResponseErr, errors.New("invalid verify code")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockFailGetAccessTokenRepo,
			GetUserInfoRepository:    mockGetUserInfoRepo,
		}
		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetAccessTokenRepo.AssertExpectations(t)
	})

	t.Run("getUserInfofail", func(t *testing.T) {
		mockFailGetUserInfoRepo := new(mocks.GetUserInfoRepository)

		getUserInfoResponseErr := domain.GetUserInfoResponse{
			Error:            "invalid access token",
			ErrorDescription: "invalid access token",
		}

		mockFailGetUserInfoRepo.On(
			"GetUserInfo",
			"123456789",
		).Return(getUserInfoResponseErr, errors.New("invalid access token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			GetUserInfoRepository:    mockFailGetUserInfoRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetUserInfoRepo.AssertExpectations(t)
	})
}

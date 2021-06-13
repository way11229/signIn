package signinService_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"signIn/line/domain"
	"signIn/line/domain/mocks"
	signInService "signIn/line/service/signIn"
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
		StatusMessage:       "Just to eat",
	}

	getAccessTokenResponse := domain.AccessTokenResponse{
		AccessToken:      "123456789",
		ExpireIn:         2592000,
		IdToken:          "1234567890",
		RefreshToken:     "123456789",
		Scope:            "email&openid",
		TokenType:        "",
		Error:            "",
		ErrorDescription: "",
	}

	verifyIdTokenResponse := domain.VerifyIdTokenResponse{
		Iss:              "https://access.line.me",
		Sub:              "123456789",
		Aud:              "1234567890",
		Exp:              1504169092,
		Iat:              1504263657,
		Nonce:            "",
		Amr:              []string{"pwd"},
		Name:             "test",
		Picture:          "",
		Email:            "test@test.com",
		Error:            "",
		ErrorDescription: "",
	}

	getUserProfileResponse := domain.UserProfileResponse{
		UserId:           "123456789",
		DisplayName:      "test",
		PictureUrl:       "",
		StatusMessage:    "Just to eat",
		Error:            "",
		ErrorDescription: "",
	}

	mockGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)
	mockVerifyIdTokenRepo := new(mocks.VerifyIdTokenRepository)
	mockGetUserProfileRepo := new(mocks.GetUserProfileRepository)

	mockGetAccessTokenRepo.On(
		"GetAccessToken",
		"7G7ovtjlalaCDzWtUVO2",
	).Return(getAccessTokenResponse, nil).Times(3)

	mockVerifyIdTokenRepo.On(
		"VerifyIdToken",
		"1234567890",
	).Return(verifyIdTokenResponse, nil).Times(2)

	mockGetUserProfileRepo.On(
		"GetUserProfile",
		"123456789",
	).Return(getUserProfileResponse, nil).Once()

	t.Run("success", func(t *testing.T) {
		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyIdTokenRepository:  mockVerifyIdTokenRepo,
			GetUserProfileRepository: mockGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.NoError(t, err)
		assert.Equal(t, expectSignInResponse, signInResponse)
	})

	t.Run("getAccessTokenfail", func(t *testing.T) {
		mockFailGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)

		getAccessTokenResponseErr := domain.AccessTokenResponse{
			AccessToken:      "",
			ExpireIn:         0,
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
			VerifyIdTokenRepository:  mockVerifyIdTokenRepo,
			GetUserProfileRepository: mockGetUserProfileRepo,
		}
		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetAccessTokenRepo.AssertExpectations(t)
	})

	t.Run("verifyIdTokenfail", func(t *testing.T) {
		mockFailVerifyIdTokenRepo := new(mocks.VerifyIdTokenRepository)

		verifyIdTokenResponseErr := domain.VerifyIdTokenResponse{
			Iss:              "",
			Sub:              "",
			Aud:              "",
			Exp:              0,
			Iat:              0,
			Nonce:            "",
			Amr:              []string{},
			Name:             "",
			Picture:          "",
			Email:            "",
			Error:            "invalid id token",
			ErrorDescription: "invalid id token",
		}

		mockFailVerifyIdTokenRepo.On(
			"VerifyIdToken",
			"1234567890",
		).Return(verifyIdTokenResponseErr, errors.New("invalid id token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyIdTokenRepository:  mockFailVerifyIdTokenRepo,
			GetUserProfileRepository: mockGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailVerifyIdTokenRepo.AssertExpectations(t)
	})

	t.Run("getUserProfilefail", func(t *testing.T) {
		mockFailGetUserProfileRepo := new(mocks.GetUserProfileRepository)

		getUserProfileResponseErr := domain.UserProfileResponse{
			UserId:           "",
			DisplayName:      "",
			PictureUrl:       "",
			StatusMessage:    "",
			Error:            "invalid access token",
			ErrorDescription: "invalid access token",
		}

		mockFailGetUserProfileRepo.On(
			"GetUserProfile",
			"123456789",
		).Return(getUserProfileResponseErr, errors.New("invalid access token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyIdTokenRepository:  mockVerifyIdTokenRepo,
			GetUserProfileRepository: mockFailGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetUserProfileRepo.AssertExpectations(t)
	})
}

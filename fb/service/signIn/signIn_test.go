package signinService_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"signIn/fb/domain"
	"signIn/fb/domain/mocks"
	signInService "signIn/fb/service/signIn"
)

func TestSignIn(t *testing.T) {
	verifyCode := "7G7ovtjlalaCDzWtUVO2"

	expectSignInResponse := domain.SignInResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
		Birthday:            "2021-06-30",
	}

	getAccessTokenResponse := domain.GetAccessTokenResponse{
		AccessToken: "123456789",
		ExpiresIn:   2592000,
		TokenType:   "",
	}

	verifyTokenResponse := domain.VerifyTokenResponse{
		Data: struct {
			AppId       uint64 "json:\"app_id\""
			Type        string "json:\"type\""
			Application string "json:\"application\""
			ExpiresAt   uint32 "json:\"expires_at\""
			IsValid     bool   "json:\"isValid\""
			IssuedAt    uint32 "json:\"issue_at\""
			Metadata    struct {
				Sso string "json:\"sso\""
			} "json:\"metadata\""
			Scopes []string "json:\"scopes\""
			UserId string   "json:\"user_id\""
		}{
			AppId:       138483919580948,
			Type:        "USER",
			Application: "test",
			ExpiresAt:   1352419328,
			IsValid:     true,
			IssuedAt:    1347235328,
			Metadata: struct {
				Sso string "json:\"sso\""
			}{Sso: "iphone-safari"},
			Scopes: []string{"email", "publish_actions"},
			UserId: "123456789",
		},
	}

	getUserProfileResponse := domain.GetUserProfileResponse{
		UserId: "123456789",
		Name:   "test",
		Picture: domain.PictureContent{
			Data: struct {
				Height        uint32 "json:\"height\""
				Is_silhouette bool   "json:\"is_silhouette\""
				Url           string "json:\"url\""
				Width         uint32 "json:\"width\""
			}{
				Height:        50,
				Is_silhouette: true,
				Url:           "",
				Width:         50,
			},
		},
		Email:    "test@test.com",
		Birthday: "2021-06-30",
	}

	mockGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)
	mockVerifyTokenRepo := new(mocks.VerifyTokenRepository)
	mockGetUserProfileRepo := new(mocks.GetUserProfileRepository)

	mockGetAccessTokenRepo.On(
		"GetAccessToken",
		verifyCode,
	).Return(getAccessTokenResponse, nil).Times(3)

	mockVerifyTokenRepo.On(
		"VerifyToken",
		verifyCode,
		getAccessTokenResponse.AccessToken,
	).Return(verifyTokenResponse, nil).Times(2)

	mockGetUserProfileRepo.On(
		"GetUserProfile",
		verifyTokenResponse.Data.UserId,
		getAccessTokenResponse.AccessToken,
	).Return(getUserProfileResponse, nil).Once()

	t.Run("success", func(t *testing.T) {
		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyTokenRepository:    mockVerifyTokenRepo,
			GetUserProfileRepository: mockGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.NoError(t, err)
		assert.Equal(t, expectSignInResponse, signInResponse)
	})

	t.Run("getAccessTokenfail", func(t *testing.T) {
		mockFailGetAccessTokenRepo := new(mocks.GetAccessTokenRepository)

		getAccessTokenResponseErr := domain.GetAccessTokenResponse{
			AccessToken: "",
			ExpiresIn:   0,
			TokenType:   "",
		}

		mockFailGetAccessTokenRepo.On(
			"GetAccessToken",
			verifyCode,
		).Return(getAccessTokenResponseErr, errors.New("invalid verify code")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockFailGetAccessTokenRepo,
			VerifyTokenRepository:    mockVerifyTokenRepo,
			GetUserProfileRepository: mockGetUserProfileRepo,
		}
		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetAccessTokenRepo.AssertExpectations(t)
	})

	t.Run("verifyTokenfail", func(t *testing.T) {
		mockFailVerifyIdTokenRepo := new(mocks.VerifyTokenRepository)

		verifyIdTokenResponseErr := domain.VerifyTokenResponse{
			Data: struct {
				AppId       uint64 "json:\"app_id\""
				Type        string "json:\"type\""
				Application string "json:\"application\""
				ExpiresAt   uint32 "json:\"expires_at\""
				IsValid     bool   "json:\"isValid\""
				IssuedAt    uint32 "json:\"issue_at\""
				Metadata    struct {
					Sso string "json:\"sso\""
				} "json:\"metadata\""
				Scopes []string "json:\"scopes\""
				UserId string   "json:\"user_id\""
			}{},
		}

		mockFailVerifyIdTokenRepo.On(
			"VerifyToken",
			verifyCode,
			getAccessTokenResponse.AccessToken,
		).Return(verifyIdTokenResponseErr, errors.New("invalid id token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyTokenRepository:    mockFailVerifyIdTokenRepo,
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

		getUserProfileResponseErr := domain.GetUserProfileResponse{
			UserId:   "",
			Name:     "",
			Picture:  domain.PictureContent{},
			Email:    "",
			Birthday: "",
		}

		mockFailGetUserProfileRepo.On(
			"GetUserProfile",
			verifyTokenResponse.Data.UserId,
			getAccessTokenResponse.AccessToken,
		).Return(getUserProfileResponseErr, errors.New("invalid access token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetAccessTokenRepository: mockGetAccessTokenRepo,
			VerifyTokenRepository:    mockVerifyTokenRepo,
			GetUserProfileRepository: mockFailGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(verifyCode)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetUserProfileRepo.AssertExpectations(t)
	})
}

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
	accessToken := "123456789"
	extra := `{"userId": "123456789" , "expireIn": 2592000}`

	expectSignInResponse := domain.SignInResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
		Birthday:            "2021-01-01",
	}

	getUserProfileResponse := domain.UserProfileResponse{
		UserId: "123456789",
		Name:   "test",
		Picture: domain.PictureContent{
			Data: domain.PictureContentData{
				Height:        50,
				Is_silhouette: true,
				Url:           "",
				Width:         50,
			},
		},
		Email:    "test@test.com",
		Birthday: "2021-01-01",
	}

	mockGetUserProfileRepo := new(mocks.GetUserProfileRepository)

	mockGetUserProfileRepo.On(
		"GetUserProfile",
		"123456789",
		"123456789",
	).Return(getUserProfileResponse, nil).Once()

	t.Run("success", func(t *testing.T) {
		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetUserProfileRepository: mockGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(accessToken, extra)

		assert.NoError(t, err)
		assert.Equal(t, expectSignInResponse, signInResponse)
	})

	t.Run("getUserProfilefail", func(t *testing.T) {
		mockFailGetUserProfileRepo := new(mocks.GetUserProfileRepository)

		getUserProfileResponseErr := domain.UserProfileResponse{}
		mockFailGetUserProfileRepo.On(
			"GetUserProfile",
			"123456789",
			"123456789",
		).Return(getUserProfileResponseErr, errors.New("invalid access token")).Once()

		signInServiceRepositoryList := domain.SignInServiceRepositoryList{
			GetUserProfileRepository: mockFailGetUserProfileRepo,
		}

		s := signInService.New(signInServiceRepositoryList)
		signInResponse, err := s.SignIn(accessToken, extra)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInResponse{}, signInResponse)

		mockFailGetUserProfileRepo.AssertExpectations(t)
	})
}

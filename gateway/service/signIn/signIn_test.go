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
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockLineRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(expectLineResponse, nil).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithLine(ctx, accessData)

		assert.NoError(t, err)
		assert.Equal(t, expectRtn, signInData)

		mockLineRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockLineRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(domain.LineResponse{}, errors.New("Line sign in Error")).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithLine(ctx, accessData)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInData{}, signInData)

		mockLineRepo.AssertExpectations(t)
	})
}

func TestSignInWithFb(t *testing.T) {
	ctx := context.Background()
	accessData := domain.AccessData{
		Token: "7G7ovtjlalaCDzWtUVO2",
		Extra: "",
	}

	expectFbResponse := domain.FbResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
		Birthday:            "2021-01-01",
	}

	extraJson, _ := json.Marshal(expectFbResponse)

	expectRtn := domain.SignInData{
		ID:    expectFbResponse.UserId,
		Name:  expectFbResponse.Name,
		Email: expectFbResponse.Email,
		Phone: "",
		Extra: string(extraJson),
	}

	t.Run("success", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockFbRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(expectFbResponse, nil).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithFb(ctx, accessData)

		assert.NoError(t, err)
		assert.Equal(t, expectRtn, signInData)

		mockLineRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockFbRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(domain.FbResponse{}, errors.New("Fb sign in Error")).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithFb(ctx, accessData)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInData{}, signInData)

		mockLineRepo.AssertExpectations(t)
	})
}

func TestSignInWithGoogle(t *testing.T) {
	ctx := context.Background()
	accessData := domain.AccessData{
		Token: "7G7ovtjlalaCDzWtUVO2",
		Extra: "",
	}

	expectGoogleResponse := domain.GoogleResponse{
		AccessToken:         "123456789",
		AccessTokenExpireIn: 2592000,
		RefreshToken:        "123456789",
		UserId:              "123456789",
		Name:                "test",
		Picture:             "",
		Email:               "test@test.com",
	}

	extraJson, _ := json.Marshal(expectGoogleResponse)

	expectRtn := domain.SignInData{
		ID:    expectGoogleResponse.UserId,
		Name:  expectGoogleResponse.Name,
		Email: expectGoogleResponse.Email,
		Phone: "",
		Extra: string(extraJson),
	}

	t.Run("success", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockGoogleRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(expectGoogleResponse, nil).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithGoogle(ctx, accessData)

		assert.NoError(t, err)
		assert.Equal(t, expectRtn, signInData)

		mockLineRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockLineRepo := new(mocks.LineRepository)
		mockFbRepo := new(mocks.FbRepository)
		mockGoogleRepo := new(mocks.GoogleRepository)

		mockGoogleRepo.On(
			"SendSignInRequest",
			mock.Anything,
			mock.MatchedBy(
				func(accessData domain.AccessData) bool {
					return (accessData.Token == "7G7ovtjlalaCDzWtUVO2") && (accessData.Extra == "")
				},
			),
		).Return(domain.GoogleResponse{}, errors.New("Google sign in Error")).Once()

		repositoryList := domain.RepositoryList{
			LineRepository:   mockLineRepo,
			FbRepository:     mockFbRepo,
			GoogleRepository: mockGoogleRepo,
		}

		s := signInService.New(repositoryList)
		signInData, err := s.SignInWithGoogle(ctx, accessData)

		assert.Error(t, err)
		assert.Equal(t, domain.SignInData{}, signInData)

		mockLineRepo.AssertExpectations(t)
	})
}

package signinService

import (
	"signIn/line/domain"
)

type signInService struct {
	GetAccessTokenRepository domain.GetAccessTokenRepository
	VerifyIdTokenRepository  domain.VerifyIdTokenRepository
	GetUserProfileRepository domain.GetUserProfileRepository
}

func New(repositoryList domain.SignInServiceRepositoryList) domain.SignInService {
	return &signInService{
		GetAccessTokenRepository: repositoryList.GetAccessTokenRepository,
		VerifyIdTokenRepository:  repositoryList.VerifyIdTokenRepository,
		GetUserProfileRepository: repositoryList.GetUserProfileRepository,
	}
}

func (ss *signInService) SignIn(verifyCode string) (domain.SignInResponse, error) {
	rtn := domain.SignInResponse{}

	getAccessTokenResponse, getAccessTokenErr := ss.GetAccessTokenRepository.GetAccessToken(verifyCode)
	if getAccessTokenErr != nil {
		return rtn, getAccessTokenErr
	}

	verifyIdTokenResponse, verifyIdTokenErr := ss.VerifyIdTokenRepository.VerifyIdToken(getAccessTokenResponse.IdToken)
	if verifyIdTokenErr != nil {
		return rtn, verifyIdTokenErr
	}

	getUserProfileResponse, getUserProfileErr := ss.GetUserProfileRepository.GetUserProfile(getAccessTokenResponse.AccessToken)
	if getUserProfileErr != nil {
		return rtn, getUserProfileErr
	}

	rtn = domain.SignInResponse{
		AccessToken:         getAccessTokenResponse.AccessToken,
		AccessTokenExpireIn: getAccessTokenResponse.ExpireIn,
		RefreshToken:        getAccessTokenResponse.RefreshToken,
		UserId:              getUserProfileResponse.UserId,
		Name:                getUserProfileResponse.DisplayName,
		Picture:             getUserProfileResponse.PictureUrl,
		Email:               verifyIdTokenResponse.Email,
		StatusMessage:       getUserProfileResponse.StatusMessage,
	}

	return rtn, nil
}

package signinService

import (
	"signIn/fb/domain"
)

type signInService struct {
	GetAccessTokenRepository domain.GetAccessTokenRepository
	VerifyTokenRepository    domain.VerifyTokenRepository
	GetUserProfileRepository domain.GetUserProfileRepository
}

func New(repositoryList domain.SignInServiceRepositoryList) domain.SignInService {
	return &signInService{
		GetAccessTokenRepository: repositoryList.GetAccessTokenRepository,
		VerifyTokenRepository:    repositoryList.VerifyTokenRepository,
		GetUserProfileRepository: repositoryList.GetUserProfileRepository,
	}
}

func (ss *signInService) SignIn(verifyCode string) (domain.SignInResponse, error) {
	rtn := domain.SignInResponse{}

	getAccessTokenResponse, getAccessTokenErr := ss.GetAccessTokenRepository.GetAccessToken(verifyCode)
	if getAccessTokenErr != nil {
		return rtn, getAccessTokenErr
	}

	verifyTokenResponse, verifyTokenErr := ss.VerifyTokenRepository.VerifyToken(getAccessTokenResponse.AccessToken)
	if verifyTokenErr != nil {
		return rtn, verifyTokenErr
	}

	getUserProfileResponse, getUserProfileErr := ss.GetUserProfileRepository.GetUserProfile(verifyTokenResponse.Data.UserId, getAccessTokenResponse.AccessToken)
	if getUserProfileErr != nil {
		return rtn, getUserProfileErr
	}

	rtn = domain.SignInResponse{
		AccessToken:         getAccessTokenResponse.AccessToken,
		AccessTokenExpireIn: getAccessTokenResponse.ExpiresIn,
		UserId:              getUserProfileResponse.UserId,
		Name:                getUserProfileResponse.Name,
		Picture:             getUserProfileResponse.Picture.Data.Url,
		Email:               getUserProfileResponse.Email,
		Birthday:            getUserProfileResponse.Birthday,
	}

	return rtn, nil
}

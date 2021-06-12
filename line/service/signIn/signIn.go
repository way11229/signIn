package signinService

import (
	"signIn/line/domain"
)

type signInService struct {
	GetAccessTokenRepository domain.GetAccessTokenRepository
	GetUserDataRepository    domain.GetUserDataRepository
}

func New(repositoryList domain.SignInServiceRepositoryList) domain.SignInService {
	return &signInService{
		GetAccessTokenRepository: repositoryList.GetAccessTokenRepository,
		GetUserDataRepository:    repositoryList.GetUserDataRepository,
	}
}

func (ss *signInService) SignIn(verifyCode string) (domain.SignInResponse, error) {
	rtn := domain.SignInResponse{}

	getAccessTokenResponse, getAccessTokenErr := ss.GetAccessTokenRepository.GetAccessToken(verifyCode)
	if getAccessTokenErr != nil {
		return rtn, getAccessTokenErr
	}

	userData, getUserDataErr := ss.GetUserDataRepository.GetUserData(getAccessTokenResponse.IdToken)
	if getUserDataErr != nil {
		return rtn, getUserDataErr
	}

	rtn = domain.SignInResponse{
		AccessToken:         getAccessTokenResponse.AccessToken,
		AccessTokenExpireIn: getAccessTokenResponse.ExpireIn,
		RefreshToken:        getAccessTokenResponse.RefreshToken,
		UserId:              userData.Sub,
		Name:                userData.Name,
		Picture:             userData.Picture,
		Email:               userData.Email,
	}

	return rtn, nil
}

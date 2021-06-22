package signinService

import (
	"signIn/google/domain"
)

type signInService struct {
	GetAccessTokenRepository domain.GetAccessTokenRepository
	GetUserInfoRepository    domain.GetUserInfoRepository
}

func New(repositoryList domain.SignInServiceRepositoryList) domain.SignInService {
	return &signInService{
		GetAccessTokenRepository: repositoryList.GetAccessTokenRepository,
		GetUserInfoRepository:    repositoryList.GetUserInfoRepository,
	}
}

func (ss *signInService) SignIn(verifyCode string) (domain.SignInResponse, error) {
	rtn := domain.SignInResponse{}

	getAccessTokenResponse, getAccessTokenErr := ss.GetAccessTokenRepository.GetAccessToken(verifyCode)
	if getAccessTokenErr != nil {
		return rtn, getAccessTokenErr
	}

	getUserInfoResponse, getUserInfoErr := ss.GetUserInfoRepository.GetUserInfo(getAccessTokenResponse.AccessToken)
	if getUserInfoErr != nil {
		return rtn, getUserInfoErr
	}

	rtn = domain.SignInResponse{
		AccessToken:         getAccessTokenResponse.AccessToken,
		AccessTokenExpireIn: getAccessTokenResponse.ExpiresIn,
		RefreshToken:        getAccessTokenResponse.RefreshToken,
		UserId:              getUserInfoResponse.Sub,
		Name:                getUserInfoResponse.Name,
		Picture:             getUserInfoResponse.Picture,
		Email:               getUserInfoResponse.Email,
	}

	return rtn, nil
}

package signinService

import (
	"encoding/json"
	"signIn/fb/domain"
)

type signInService struct {
	GetUserProfileRepository domain.GetUserProfileRepository
}

func New(repositoryList domain.SignInServiceRepositoryList) domain.SignInService {
	return &signInService{
		GetUserProfileRepository: repositoryList.GetUserProfileRepository,
	}
}

func (ss *signInService) SignIn(accessToken, extra string) (domain.SignInResponse, error) {
	var extraDecode domain.GrpcExtraContent
	rtn := domain.SignInResponse{}

	extraDecodeErr := json.Unmarshal([]byte(extra), &extraDecode)
	if extraDecodeErr != nil {
		return rtn, extraDecodeErr
	}

	getUserProfileResponse, getUserProfileErr := ss.GetUserProfileRepository.GetUserProfile(extraDecode.UserId, accessToken)
	if getUserProfileErr != nil {
		return rtn, getUserProfileErr
	}

	rtn = domain.SignInResponse{
		AccessToken:         accessToken,
		AccessTokenExpireIn: extraDecode.ExpireIn,
		UserId:              getUserProfileResponse.UserId,
		Name:                getUserProfileResponse.Name,
		Picture:             getUserProfileResponse.Picture.Data.Url,
		Email:               getUserProfileResponse.Email,
		Birthday:            getUserProfileResponse.Birthday,
	}

	return rtn, nil
}

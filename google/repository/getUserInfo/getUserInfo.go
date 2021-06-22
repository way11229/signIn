package getUserInfoRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"signIn/google/domain"
)

type getUserInfoRepository struct {
	GoogleConfig domain.GoogleConfig
}

func New(gc domain.GoogleConfig) domain.GetUserInfoRepository {
	return &getUserInfoRepository{GoogleConfig: gc}
}

func (g *getUserInfoRepository) GetUserInfo(accessToken string) (domain.GetUserInfoResponse, error) {
	rtn := domain.GetUserInfoResponse{}

	request, requstErr := http.NewRequest("GET", g.GoogleConfig.UserInfoApi, nil)
	if requstErr != nil {
		return rtn, requstErr
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return rtn, responseErr
	}

	defer response.Body.Close()

	var responseDecode domain.GetUserInfoResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseDecode)
	if decodeErr != nil {
		return rtn, decodeErr
	}

	if responseDecode.Error != "" {
		return rtn, errors.New(responseDecode.ErrorDescription)
	}

	rtn = responseDecode

	return rtn, nil
}

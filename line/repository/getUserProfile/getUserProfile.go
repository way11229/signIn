package getUserProfileRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"signIn/line/domain"
)

type getUserProfileRepository struct {
	LineConfig domain.LineConfig
}

func New(lc domain.LineConfig) domain.GetUserProfileRepository {
	return &getUserProfileRepository{LineConfig: lc}
}

func (g *getUserProfileRepository) GetUserProfile(accessToken string) (domain.UserProfileResponse, error) {
	rtn := domain.UserProfileResponse{}

	request, requstErr := http.NewRequest("GET", g.LineConfig.ProfileApi, nil)
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

	var responseDecode domain.UserProfileResponse
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

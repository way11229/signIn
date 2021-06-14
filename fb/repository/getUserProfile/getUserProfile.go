package getUserProfileRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"signIn/fb/domain"
)

type getUserProfileRepository struct {
	FbConfig domain.FbConfig
}

func New(fc domain.FbConfig) domain.GetUserProfileRepository {
	return &getUserProfileRepository{FbConfig: fc}
}

func (g *getUserProfileRepository) GetUserProfile(userId, accessToken string) (domain.UserProfileResponse, error) {
	rtn := domain.UserProfileResponse{}

	requestUrl := g.FbConfig.GraphApi + userId + "?fields=id,email,name,picture,birthday&" + accessToken
	request, requstErr := http.NewRequest("GET", requestUrl, nil)
	if requstErr != nil {
		return rtn, requstErr
	}

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return rtn, responseErr
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var errorDecode domain.ErrorResponse
		decodeErr := json.NewDecoder(response.Body).Decode(&errorDecode)
		if decodeErr != nil {
			return rtn, decodeErr
		}

		return rtn, errors.New(errorDecode.Error.Message)
	}

	var responseDecode domain.UserProfileResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseDecode)
	if decodeErr != nil {
		return rtn, decodeErr
	}

	rtn = responseDecode

	return rtn, nil
}

package getUserDataRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"signIn/line/domain"
)

type getUserDataRepository struct {
	LineConfig domain.LineConfig
}

func New(lc domain.LineConfig) domain.GetUserDataRepository {
	return &getUserDataRepository{LineConfig: lc}
}

func (g *getUserDataRepository) GetUserData(idToken string) (domain.UserDataResponse, error) {
	rtn := domain.UserDataResponse{}

	requestBody := url.Values{
		"id_token":  {idToken},
		"client_id": {g.LineConfig.ClientId},
	}

	response, err := http.PostForm(g.LineConfig.VerifyApi, requestBody)
	if err != nil {
		return rtn, err
	}

	var responseDecode domain.UserDataResponse
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

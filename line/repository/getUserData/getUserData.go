package getUserDataRepository

import (
	"encoding/json"
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

	json.NewDecoder(response.Body).Decode(&rtn)

	return rtn, nil
}

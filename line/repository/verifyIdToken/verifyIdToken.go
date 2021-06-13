package verifyIdTokenRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"signIn/line/domain"
)

type verifyIdTokenRepository struct {
	LineConfig domain.LineConfig
}

func New(lc domain.LineConfig) domain.VerifyIdTokenRepository {
	return &verifyIdTokenRepository{LineConfig: lc}
}

func (v *verifyIdTokenRepository) VerifyIdToken(idToken string) (domain.VerifyIdTokenResponse, error) {
	rtn := domain.VerifyIdTokenResponse{}

	requestBody := url.Values{
		"id_token":  {idToken},
		"client_id": {v.LineConfig.ClientId},
	}

	response, err := http.PostForm(v.LineConfig.VerifyApi, requestBody)
	if err != nil {
		return rtn, err
	}

	defer response.Body.Close()

	var responseDecode domain.VerifyIdTokenResponse
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

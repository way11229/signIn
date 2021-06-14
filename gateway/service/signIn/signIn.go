package signInService

import (
	"context"
	"encoding/json"
	"signIn/gateway/domain"
)

type signInService struct {
	LineRepository domain.LineRepository
}

func New(lr domain.LineRepository) domain.SignInService {
	return &signInService{LineRepository: lr}
}

func (ss *signInService) SignInWithLine(c context.Context, accessData domain.AccessData) (domain.SignInData, error) {
	lineResponse, error := ss.LineRepository.SendSignInRequest(c, accessData)
	if error != nil {
		return domain.SignInData{}, error
	}

	extraJson, jsonErr := json.Marshal(lineResponse)
	if jsonErr != nil {
		return domain.SignInData{}, jsonErr
	}

	rtn := domain.SignInData{
		ID:    lineResponse.UserId,
		Name:  lineResponse.Name,
		Email: lineResponse.Email,
		Phone: "",
		Extra: string(extraJson),
	}
	return rtn, nil
}

func (ss *signInService) SignInWithFb(c context.Context, accessData domain.AccessData) (domain.SignInData, error) {
	rtn := domain.SignInData{
		ID:    "11223456",
		Name:  "test",
		Email: "test",
		Phone: "",
		Extra: accessData.Extra,
	}
	return rtn, nil
}

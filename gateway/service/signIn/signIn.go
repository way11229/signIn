package signInService

import (
	"context"
	"encoding/json"
	"signIn/gateway/domain"
)

type signInService struct {
	LineRepository domain.LineRepository
	FbRepository   domain.FbRepository
}

func New(repositoryList domain.RepositoryList) domain.SignInService {
	return &signInService{
		LineRepository: repositoryList.LineRepository,
		FbRepository:   repositoryList.FbRepository,
	}
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
	fbResponse, error := ss.FbRepository.SendSignInRequest(c, accessData)
	if error != nil {
		return domain.SignInData{}, error
	}

	extraJson, jsonErr := json.Marshal(fbResponse)
	if jsonErr != nil {
		return domain.SignInData{}, jsonErr
	}

	rtn := domain.SignInData{
		ID:    fbResponse.UserId,
		Name:  fbResponse.Name,
		Email: fbResponse.Email,
		Phone: "",
		Extra: string(extraJson),
	}

	return rtn, nil
}

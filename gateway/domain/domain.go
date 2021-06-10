package domain

import (
	"context"
)

type AccessData struct {
	Token string `json:"token"`
	Extra string `json:"extra"`
}

type SignInData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Extra string `json:"extra"`
}

type ServiceList struct {
	SignInService SignInService
}

type LineResponse struct {
	AccessToken         string `json:"accessToken"`
	AccessTokenExpireIn uint32 `json:"accessTokenExpireIn"`
	RefreshToken        string `json:"refreshToken"`
	UserId              string `json:"userId"`
	Name                string `json:"name"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
}

type SignInService interface {
	SignInWithLine(context.Context, AccessData) (SignInData, error)
}

type LineRepository interface {
	SendSignInRequest(context.Context, AccessData) (LineResponse, error)
}

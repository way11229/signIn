package domain

import "context"

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

type SignInServiceList struct {
	SignInWithLineService SignInWithLineService
}

type SignInWithLineService interface {
	SignInWithLine(context.Context, AccessData) (SignInData, error)
}

type SignInWithLineRepository interface {
	SignInWithLine(context.Context, AccessData) (SignInData, error)
}

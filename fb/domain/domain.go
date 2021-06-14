package domain

type FbConfig struct {
	GraphApi string `json:"graphApi"`
}

type GrpcExtraContent struct {
	UserId   string `json:"userId"`
	ExpireIn uint32 `json:"expireIn"`
}

type ErrorContent struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      uint32 `json:"code"`
	FbTraceId string `json:"fbtrace_id"`
}

type ErrorResponse struct {
	Error ErrorContent `json:"error"`
}

type UserProfileResponse struct {
	UserId   string `json:"id"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
}

type SignInResponse struct {
	AccessToken         string `json:"accessToken"`
	AccessTokenExpireIn uint32 `json:"accessTokenExpireIn"`
	UserId              string `json:"userId"`
	Name                string `json:"name"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
	Birthday            string `json:"birthday"`
}

type ServiceList struct {
	SignInService SignInService
}

type SignInServiceRepositoryList struct {
	GetUserProfileRepository GetUserProfileRepository
}

type SignInService interface {
	SignIn(string, string) (SignInResponse, error)
}

type GetUserProfileRepository interface {
	GetUserProfile(string, string) (UserProfileResponse, error)
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "signIn/google/domain"
import mock "github.com/stretchr/testify/mock"

// GetUserInfoRepository is an autogenerated mock type for the GetUserInfoRepository type
type GetUserInfoRepository struct {
	mock.Mock
}

// GetUserInfo provides a mock function with given fields: _a0
func (_m *GetUserInfoRepository) GetUserInfo(_a0 string) (domain.GetUserInfoResponse, error) {
	ret := _m.Called(_a0)

	var r0 domain.GetUserInfoResponse
	if rf, ok := ret.Get(0).(func(string) domain.GetUserInfoResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.GetUserInfoResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

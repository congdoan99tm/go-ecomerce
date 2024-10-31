package service

import (
	"context"
	"github.com/dinos/go-ecommerce-be-api/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context) error
		UpdatePasswordRegister(ctx context.Context) error
	}
	IUserInfo interface {
		GetInfoById(ctx context.Context) error
		GetAll(ctx context.Context) error
	}
	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement localUserAdmin not found for interface")
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found for interface")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

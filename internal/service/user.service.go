package service

import (
	"github.com/dinos/go-ecommerce-be-api/internal/repo"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{userRepo: repo.NewUserRepo()}
//}
//
//func (us *UserService) GetInfoUser() string {
//	return us.userRepo.GetInfoUser()
//}

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

func (us *userService) Register(email string, purpose string) int {
	//1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrorCodeSuccess
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{userRepo: userRepo}
}

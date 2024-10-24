package repo

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	"time"
)

type IUserAuthRepo interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepo struct{}

func (u *userAuthRepo) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // usr:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}

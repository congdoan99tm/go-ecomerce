package service

import (
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/internal/repo"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/crypto"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/random"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/sendto"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepo
}

func NewUserService(userRepo repo.IUserRepo, userAuthRepo repo.IUserAuthRepo) IUserService {
	return &userService{userRepo: userRepo, userAuthRepo: userAuthRepo}
}

func (us *userService) Register(email string, purpose string) int {
	//0. hasEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)
	//5. check OTP is available
	//6. user spam email

	//1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	//2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is :::%d", otp)
	//3. save OTP in Redis with expiration time
	err := us.userAuthRepo.AddOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	//4. send Email OTP
	err = sendto.SendTextEmailOTP([]string{email}, "congdoan99tm@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}

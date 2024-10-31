package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	consts "github.com/dinos/go-ecommerce-be-api/internal/const"
	"github.com/dinos/go-ecommerce-be-api/internal/database"
	"github.com/dinos/go-ecommerce-be-api/internal/model"
	"github.com/dinos/go-ecommerce-be-api/internal/utils"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/crypto"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/random"
	"github.com/dinos/go-ecommerce-be-api/internal/utils/sendto"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
	"log"
	"strconv"
	"strings"
	"time"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}
func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (code int, err error) {
	// 1. hash Email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	// 2. check user exists in user base
	userFound, err := s.r.CheckUserHasExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// 3. create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, code, err := utils.GetOtpInRedis(ctx, userKey)
	if err != nil {
		return code, err
	}
	fmt.Println(otpFound)
	// 4. gen new otp
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Println("otpNew:", otpNew)

	// 5. save OTP in Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISRER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	// 6. send OTP
	switch in.VerifyType {
	case consts.Email:
		err = sendto.SendTextEmailOTP([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// 7. save OTP to MySQL
		result, err := s.r.InsertOTPVerify(ctx,
			database.InsertOTPVerifyParams{
				VerifyOtp:     strconv.Itoa(otpNew),
				VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
				VerifyKey:     in.VerifyKey,
				VerifyKeyHash: hashKey,
			},
		)
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// 8. GetLastId
		lastIDVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		log.Print("lastIdVerifyUser:", lastIDVerifyUser)
		return response.ErrCodeSuccess, nil
	case consts.Mobile:
		return response.ErrCodeSuccess, nil
	}
	return response.ErrCodeSuccess, nil
}
func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}

package utils

import (
	"context"
	"fmt"
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/dinos/go-ecommerce-be-api/pkg/response"
	"github.com/redis/go-redis/v9"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GetOtpInRedis(ctx context.Context, userKey string) (otpFound string, code int, err error) {
	otpFound, err = global.Rdb.Get(ctx, userKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist:")
		// Handle the case where the key doesn't exist
	case err != nil:
		fmt.Println("get failed::", err)
		return "", response.ErrInvalidOTP, err
	case otpFound != "":
		return "", response.ErrCodeOtpNotExists, fmt.Errorf("otp already registered")
	}
	return otpFound, 0, nil
}

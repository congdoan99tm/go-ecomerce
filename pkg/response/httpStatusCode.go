package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParmaInvalid  = 20003 // Success
	ErrInvalidToken      = 30001 // Success
	ErrInvalidOTP        = 30002
	ErrSendEmailOtp      = 30003
	ErrCodeUserHasExists = 50001 // user has already registered
	ErrCodeOtpNotExists  = 60009
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParmaInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrSendEmailOtp:      "Failed to send email OTP",
	ErrCodeUserHasExists: "user has already registered",
	ErrCodeOtpNotExists:  "otp exists",
}

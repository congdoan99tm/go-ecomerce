package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParmaInvalid = 20003 // Success
	ErrorInvalidToken     = 30001 // Success

)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParmaInvalid: "Email is invalid",
	ErrorInvalidToken:     "Token is invalid",
}

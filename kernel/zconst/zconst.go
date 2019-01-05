package zconst

import "github.com/satori/go.uuid"

var InitTrackId string

func init() {
	InitTrackId = uuid.NewV4().String()
}

const (
	LOGINSSO  = "/member/login_sso"
	LOGIN_SUC = "/member/succ"
)

const (
	RES_SUCCESS     = 0
	RES_SUCCESS_MSG = "success"

	RES_SERVER_EXCEPTION     = 88
	RES_SERVER_EXCEPTION_MSG = "server exception"

	RES_COMMON     = 99
	RES_COMMON_MSG = "common exception"

	RES_AUTH     = 10000
	RES_AUTH_MSG = "auth exception"

	RES_ACOUNT     = 10001
	RES_ACOUNT_MSG = "aount is error"

	RES_PASSWORD     = 10002
	RES_PASSWORD_MSG = "password is error"
)

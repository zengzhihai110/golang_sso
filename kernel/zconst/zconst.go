package zconst

import "github.com/satori/go.uuid"

var InitTrackId string

func init() {
	InitTrackId = uuid.NewV4().String()
}

const (
	LOGINSSO = "/member/login_sso"
)

const (
	RES_SUCCESS     = 0
	RES_SUCCESS_MSG = "success"

	RES_SERVER_EXCEPTION     = 88
	RES_SERVER_EXCEPTION_MSG = "server exception"

	RES_COMMON     = 99
	RES_COMMON_MSG = "common exception"
)

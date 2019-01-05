package util

import "net/http"

type ReqParam struct {
	Res     http.ResponseWriter
	Req     *http.Request
	TrackId string
}

package zengzhihai_pipeline

import (
	"errors"
	"github.com/json-iterator/go"
	"fmt"
	"runtime/debug"
	"zengzhihai.com/golang_sso/kernel/comm_log"
	"zengzhihai.com/golang_sso/kernel/zconst"
	"zengzhihai.com/golang_sso/kernel/util"
)

type RequestFilter struct {
}

func (this *RequestFilter) Process(data interface{}) (interface{}, error) {
	res := new(util.ReturnResult)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reqParam, ok := data.(*util.ReqParam)
	if !ok {
		res.Msg = zconst.RES_SERVER_EXCEPTION_MSG
		res.Code = zconst.RES_SERVER_EXCEPTION
		return util.DataToStr(res), errors.New("RequestFilter input type should be error")
	}

	defer func() {
		if r := recover(); r != nil {
			tLog := make(map[string]interface{})
			tLog["error"] = fmt.Sprintf("request error：%+v", r)
			tLog["debug_stack"] = fmt.Sprintf("stack error：%+v", string(debug.Stack()))
			jsonLog, _ := json.Marshal(tLog)
			comm_log.Error(reqParam.TrackId, string(jsonLog))
			debug.PrintStack()
		}
	}()

	//这里打印请求info
	tLog := make(map[string]interface{})
	tLog["method"] = reqParam.Req.Method
	tLog["remote_addr"] = reqParam.Req.RemoteAddr
	tLog["ua"] = reqParam.Req.Header.Get("User-Agent")
	jsonLog, _ := json.Marshal(tLog)
	comm_log.Info(reqParam.TrackId, util.ClientIP(reqParam.Req), reqParam.Req.RequestURI, string(jsonLog))

	return reqParam, nil
}

package zengzhihai_pipeline

import (
	"zengzhihai.com/golang_sso/kernel/util"
	"github.com/gorilla/sessions"
	"fmt"
	"github.com/json-iterator/go"
	"errors"
	"runtime/debug"
	"zengzhihai.com/golang_sso/kernel/comm_log"
	"zengzhihai.com/golang_sso/kernel/zconst"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

type LoginSsoFilter struct {
}

type LoginSsoParam struct {
	UserId   int64
	UserName string
	Password string
}

func (this *LoginSsoFilter) Process(data interface{}) (interface{}, error) {
	res := new(util.ReturnResult)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reqParam, ok := data.(*util.ReqParam)
	if !ok {
		res.Msg = zconst.RES_SERVER_EXCEPTION_MSG
		res.Code = zconst.RES_SERVER_EXCEPTION
		return util.DataToCommJsonStr(res), errors.New("LoginSsoFilter input type should be http.Request")
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
	req := new(LoginSsoParam)
	req.UserId = 1
	req.UserName = reqParam.Req.PostFormValue("username")
	req.Password = reqParam.Req.PostFormValue("password")
	session, err := store.Get(reqParam.Req, "SESSION_KEY")
	if err != nil {
		tLog := make(map[string]interface{})
		tLog["error"] = fmt.Sprintf("request error：%+v", err)
		jsonLog, _ := json.Marshal(tLog)
		comm_log.Error(reqParam.TrackId, string(jsonLog))
		res.Msg = zconst.RES_COMMON_MSG
		res.Code = zconst.RES_COMMON
		return util.DataToCommJsonStr(res), nil
	}
	session.Options = &sessions.Options{
		Path:     "/",
		Domain:   "test.zengzhihai.com",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Values["username"] = req.UserName
	session.Values["userid"] = req.UserId
	err = session.Save(reqParam.Req, reqParam.Res)
	if err == nil {
		http.Redirect(reqParam.Res, reqParam.Req, zconst.LOGIN_SUC, 302)
		res.Msg = zconst.RES_SUCCESS_MSG
		res.Code = zconst.RES_SUCCESS
		return util.DataToCommJsonStr(res), nil
	} else {
		res.Msg = zconst.RES_AUTH_MSG
		res.Code = zconst.RES_AUTH
		return util.DataToCommJsonStr(res), nil
	}

	res.Msg = zconst.RES_COMMON_MSG
	res.Code = zconst.RES_COMMON
	return util.DataToCommJsonStr(res), nil
}

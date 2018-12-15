package zengzhihai_pipeline

import (
	"io"
	"net/http"

	"github.com/satori/go.uuid"
	"zengzhihai.com/golang_sso/kernel/comm_log"
	"github.com/easierway/pipefiter_framework/pipefilter"
	"zengzhihai.com/golang_sso/kernel/util"
	"zengzhihai.com/golang_sso/kernel/zconst"
)

type BusinessHandler struct {
	Pipeline pipefilter.Filter
}

func CreateHandler(pipelineName string) (BusinessHandler, error) {
	pl, err := CreatePipeline(pipelineName)
	if err != nil {
		return BusinessHandler{}, err
	}

	handler := BusinessHandler{
		Pipeline: pl,
	}
	return handler, nil
}

func (self *BusinessHandler) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	//根据http协议生成trackId
	trackId := ""
	//catch panic
	defer func() {
		if r := recover(); r != nil {
			comm_log.Error(trackId, util.ToCommJsonError("http error ", r))
			c.WriteHeader(500)
			res := new(util.ReturnResult)
			res.Msg = zconst.RES_SERVER_EXCEPTION_MSG
			res.Code = zconst.RES_SERVER_EXCEPTION
			io.WriteString(c, *sutil.InterfaceToStr(res))
		}
	}()

	//追加trackId
	if 0 == len(req.URL.Query().Get("trackid")) {
		trackId = uuid.NewV4().String()
	} else {
		trackId = req.URL.Query().Get("trackid")
	}

	//重新传输参数
	reqParam := new(sutil.ReqParam)
	reqParam.TrackId = trackId
	reqParam.Req = req

	//business process
	ret, err := self.Pipeline.Process(reqParam)
	if err != nil {
		//add log
		comm_log.Error(trackId, sutil.ToCommJsonError("error with: "+err.Error()))
		c.WriteHeader(500)
	}
	c.Header().Set("Content-Type", "application/json")
	io.WriteString(c, *ret.(*string))
}
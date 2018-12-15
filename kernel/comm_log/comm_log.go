package comm_log

import (
	"log"

	"zengzhihai.com/golang_sso/kernel/comm_base"

	"github.com/cihub/seelog"
	"zengzhihai.com/golang_sso/kernel/comm_cfg"
)

var logg seelog.LoggerInterface
var appPath string
var workPath string
var debug string
var err error


func init() {
	appPath = comm_base.Get_main_path()
	workPath = comm_base.Get_work_path()
	debug = comm_cfg.GetValue("log_conf", "level")

	runtimeLog := comm_cfg.GetValue("log_conf", "runtime_log")

	logg, err = seelog.LoggerFromConfigAsFile(comm_base.Join_path(appPath, runtimeLog))

	if err != nil {
		logg, err = seelog.LoggerFromConfigAsFile(comm_base.Join_path(workPath, runtimeLog))
		if err != nil {
			log.Fatalf("config log err, err = %s", err)
		}
	}

}

func Flush() {
	logg.Flush()
}

func Debug(trackId string, v ...interface{}) {
	if debug == "debug" {
		logg.Debug("["+trackId+"] ", v)
	}
}

func Info(trackId string, v ...interface{}) {
	logg.Info("["+trackId+"] ", v)
}

func Warn(trackId string, v ...interface{}) {
	logg.Warn("["+trackId+"] ", v)
}

func Error(trackId string, v ...interface{}) {
	logg.Error("["+trackId+"] ", v)
}

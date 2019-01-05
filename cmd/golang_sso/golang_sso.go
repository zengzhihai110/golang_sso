package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"github.com/json-iterator/go"
	"time"

	"zengzhihai.com/golang_sso/kernel/comm_builder"
	"zengzhihai.com/golang_sso/kernel/comm_cfg"
	"zengzhihai.com/golang_sso/kernel/comm_log"
	"zengzhihai.com/golang_sso/kernel/zconst"
	"syscall"
	"os/signal"
	"net/http"
	"log"
	//"golang.org/x/sys/windows"
	"zengzhihai.com/golang_sso/kernel/zengzhihai_pipeline"
)

var (
	PROGRAM_VERSION  string = "1.0.0"
	COMPILER_VERSION string = "1.0.0"
	BUILD_TIME       string = time.Now().Format("2006-01-02 15:04:05")
	AUTHOR           string = "zengzhihai110@gmail.com"
)

func main() {
	//捕获fata，或者程序奔溃
	logFile, err := os.OpenFile("/tmp/fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Fatal("服务启动出错", "打开异常日志文件失败", err)
	}
	if runtime.GOOS == "windows" {
		// 将进程标准出错重定向至文件，进程崩溃时运行时将向该文件记录协程调用栈信息
		//windows.SetStdHandle(windows.STD_ERROR_HANDLE, windows.Handle(logFile.Fd()))
	} else {
		syscall.Dup2(int(logFile.Fd()), 2)
	}

	//加载图标
	comm_builder.Banner_show(PROGRAM_VERSION, COMPILER_VERSION, BUILD_TIME, AUTHOR)

	//加载配置文件信息
	address := comm_cfg.GetValue("system", "server_address")
	cpuCore := comm_cfg.Int("system", "cpu_num")

	//  设置并发度
	CORE_NUM := runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(CORE_NUM * cpuCore)

	//处理信号量
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	//启动服务
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t, " ready to start http server ", address)
	tLog := make(map[string]interface{})
	tLog["now_time"] = t
	tLog["data"] = "ready to start http server " + address
	jsonLog, _ := json.Marshal(tLog)
	comm_log.Info(zconst.InitTrackId, string(jsonLog))
	go func() {
		defer func() {
			if err := recover(); err != nil {
				tLog = make(map[string]interface{})
				tLog["data"] = fmt.Sprintf("srv ListenAdnServe defer error: %+v ", err)
				jsonLog, _ = json.Marshal(tLog)
				comm_log.Error(zconst.InitTrackId, string(jsonLog))
				debug.PrintStack()
			}
		}()

		//loginsso
		LoginSso, err := zengzhihai_pipeline.CreateHandler(zengzhihai_pipeline.LoginSso)
		if err != nil {
			tLog = make(map[string]interface{})
			tLog["data"] = "CreateHanler error: " + err.Error()
			jsonLog, _ = json.Marshal(tLog)
			comm_log.Error(zconst.InitTrackId, string(jsonLog))
			exit <- syscall.SIGTERM
		}
		http.Handle(zconst.LOGINSSO, &LoginSso)

		if err := http.ListenAndServe(address, nil); err != nil {
			tLog = make(map[string]interface{})
			tLog["data"] = "start http server failed:" + err.Error()
			jsonLog, _ = json.Marshal(tLog)
			comm_log.Error(zconst.InitTrackId, string(jsonLog))
			exit <- syscall.SIGTERM
		}
	}()

	<-exit
	tLog = make(map[string]interface{})
	tLog["data"] = "http server will be down"
	jsonLog, _ = json.Marshal(tLog)
	comm_log.Error(zconst.InitTrackId, string(jsonLog))
	comm_log.Flush()

}

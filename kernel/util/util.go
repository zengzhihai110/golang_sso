package util

import (
	"github.com/json-iterator/go"
	"net/http"
	"strings"
	"net"
	"fmt"
	"encoding/json"
)

func DataToStr(res interface{}) *string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	jRes, _ := json.Marshal(res)
	tRes := string(jRes)
	return &tRes
}

func ClientIP(req *http.Request) string {
	forwardedByClientIP := true
	if forwardedByClientIP {
		clientIP := strings.TrimSpace(GetRequestHeader(req, "X-Real-Ip"))
		if len(clientIP) > 0 {
			return clientIP
		}
		clientIP = GetRequestHeader(req, "X-Forwarded-For")
		if index := strings.IndexByte(clientIP, ','); index >= 0 {
			clientIP = clientIP[0:index]
		}
		clientIP = strings.TrimSpace(clientIP)
		if len(clientIP) > 0 {
			return clientIP
		}
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func GetRequestHeader(req *http.Request, key string) string {
	if values, _ := req.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func DataToCommJsonStr(str ...interface{}) string {
	tmpLog := make(map[string]interface{})
	lastStr := ""
	for _, v := range str {
		if tmpV, ok := v.(string); ok {
			lastStr = lastStr + tmpV
		} else {
			lastStr = lastStr + fmt.Sprintf(" %+v", str)
		}
	}
	tmpLog["error"] = lastStr
	jsonLog, _ := json.Marshal(tmpLog)
	return string(jsonLog)
}

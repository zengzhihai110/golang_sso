package comm_base

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Get_main_name() string {
	return strings.Split(filepath.Base(os.Args[0]), ".")[0]
}

func Get_main_path() string {
	fileAbsPath, _ := filepath.Abs(os.Args[0])
	program_path := filepath.Dir(fileAbsPath)

	return program_path
}

func Get_work_path() string {
	work_path, _ := os.Getwd()
	work_path, _ = filepath.Abs(work_path)
	return work_path
}

func Join_path(arr_str ...string) string {
	return strings.Join(arr_str, "/")
}

func Get_tcp_address(ip string, port interface{}) string {
	var str_port string
	switch value := port.(type) {
	case int:
		str_port = strconv.Itoa(value)
	case string:
		str_port = value
	}
	return ip + ":" + str_port
}

// Convert uint to string
func Inet_ntoa(ipnr int64) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

// Convert string to int64
func Inet_aton(ipnr string) int64 {
	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

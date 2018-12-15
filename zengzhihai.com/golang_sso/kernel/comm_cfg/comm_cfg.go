package comm_cfg

import (
	"log"
	"path/filepath"

	"zengzhihai.com/golang_sso/kernel/comm_base"

	"github.com/Unknwon/goconfig"
)

var default_cfg *goconfig.ConfigFile

func init() {
	app_path := comm_base.Get_main_path()
	// 支持go run 和 go test
	work_path := comm_base.Get_work_path()
	config_path := "conf/etc.ini"

	// 现在应用对应目录查找
	cfg, err := goconfig.LoadConfigFile(comm_base.Join_path(app_path, config_path))

	// 查找失败，在当前目录查找
	if err != nil {
		cfg, err = goconfig.LoadConfigFile(comm_base.Join_path(work_path, config_path))
	}

	// 当前目录查找失败，在去当前目录上一级目录查找
	if err != nil {
		cfg, err = goconfig.LoadConfigFile(comm_base.Join_path(filepath.Dir(work_path), config_path))
	}

	// 都没有找到，返回失败
	if err != nil {
		program_name := comm_base.Get_main_name()
		log.Fatalf("program %s, load config conf.ini failed, error = %s!", program_name, err)
	}

	default_cfg = cfg
}

func GetValue(section, key string) string {
	value, err := default_cfg.GetValue(section, key)
	if err != nil {
		log.Fatalf("section = %s, get item key = %s failed, err = %s!", section, key, err)
	}
	return value
}

func GetSections() []string {
	return default_cfg.GetSectionList()
}

func GetDefaultCofig() *goconfig.ConfigFile {
	return default_cfg
}

func Int(section, key string) int {
	value, err := default_cfg.Int(section, key)
	if err != nil {
		log.Fatalf("section = %s, get item key = %s failed, Int val err = %s!", section, key, err)
	}
	return value
}
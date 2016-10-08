package ckenv

// 프로그램에서 간단하게 사용하기 위해서 viper 를 감싼 것.
// 내부에서 로그 공용 object 까지 생성 한다.

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/dhkim94/cklog"
	"path/filepath"
	"strings"
)

var (
	clog *cklog.Cklogger
	conf *viper.Viper
)

func Init(confFileName string) bool {
	confDir := filepath.Dir(confFileName)
	baseName := strings.Split(filepath.Base(confFileName), ".")
	baseNameLen := len(baseName)

	if baseNameLen < 2 {
		fmt.Printf("[FAIL] Invalid properties file: %s\n", confFileName)
		return false
	}

	if baseName[baseNameLen - 1] != "properties" {
		fmt.Printf("[FAIL] config file is not properties file: %s\n", confFileName)
		return false
	}

	confFile := strings.Join(baseName[:baseNameLen - 1], ".")

	conf = viper.New()
	conf.SetConfigName(confFile)
	conf.AddConfigPath(confDir)

	err := conf.ReadInConfig()
	if err != nil {
		fmt.Printf("[FAIL] Not found config file: %s\n", confFileName)
		return false
	}
	//fmt.Printf("read config file [%s/%s.properties]\n", confDir, confFile);

	logLevel := fmt.Sprintf("%s", conf.Get("log.level"))
	logOut := fmt.Sprintf("%s", conf.Get("log.output"))

	if logOut == "file" && conf.Get("log.file") == nil {
		logOut = "stdout"
	}

	clog = cklog.NewLogger(logLevel, logOut,
		fmt.Sprintf("%s", conf.Get("log.file")))

	return true
}

func SetStdOutLogger(level string)  {
	clog = cklog.NewLogger(level, "stdout", "")
}

func GetLogger() *cklog.Cklogger {
	return clog
}

func GetConf() *viper.Viper {
	return conf
}

// 값이 없으면 empty string 을 리턴 한다.
func GetValue(key string) string {
	if conf.Get(key) == nil {
		return ""
	}

	return fmt.Sprintf("%s", conf.Get(key))
}


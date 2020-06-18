package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgJwt *viper.Viper
var cfgLog *viper.Viper
var cfgSsl *viper.Viper

func CfgSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Printf("Read config file fail: %s", err.Error()))
	}

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		panic("config not found settings.jwt")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgLog = viper.Sub("settings.log")
	if cfgLog == nil {
		panic("config not found settings.log")
	}
	LogConfig = InitLog(cfgLog)

	cfgSsl = viper.Sub("settings.ssl")
	if cfgSsl == nil {
		panic("config not found settings.ssl")
	}
	SslConfig = InitSsl(cfgSsl)
}

func SetConfig(configPath string, key string, value string) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}

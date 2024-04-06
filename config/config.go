package config

import (
	"fmt"
	"github.com/Tontonnow/ddddhmlist/utils"
	"github.com/spf13/viper"
	"os"
)

var (
	Conf   = NewConfig()
	IsLoad = false
)

type Config struct {
	Mac       string               `json:"Mac"`
	AndroidId string               `json:"AndroidId"`
	Proxy     string               `json:"Proxy"`
	WebConfig map[string]WebConfig `json:"WebConfig"`
}
type WebConfig struct {
	Authorization string            `json:"Authorization"`
	ExpiresIn     int               `json:"ExpiresIn"`
	Proxy         string            `json:"Proxy"`
	Headers       map[string]string `json:"Headers"`
	Cookie        string            `json:"Cookie"`
}

func NewConfig() *Config {
	return &Config{
		Mac:       "00:00:00:00:00:00",
		AndroidId: "0000000000000000",
		Proxy:     "",
	}
}
func updateConfig(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Printf("写入配置文件失败：%s\n", err)
	}
}
func InitConfig() {
	if IsLoad {
		return
	}
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取目录失败：%s\n", err))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件失败：%s\n", err)
		err = viper.SafeWriteConfigAs("config.yml")
		if err != nil {
			fmt.Printf("写入配置文件失败：%s\n", err)
		}
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Printf("再次读取配置文件失败：%s\n", err)
			os.Exit(1)
		}
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("解析配置文件失败：%s\n", err)
		os.Exit(1)
	}
	if Conf.AndroidId == "" {
		Conf.AndroidId = utils.GenerateRandomAndroidId()
	}
	if Conf.Mac == "" {
		Conf.Mac = utils.GenerateRandomMac()
	}
	IsLoad = true
}

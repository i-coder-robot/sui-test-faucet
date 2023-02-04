package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"sui-test-faucet/model"
)

var (
	DiscordConf *model.Discord
)

func getDiscordConf() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("无法解析config.yaml配置文件, %v", err)
		panic(err)
	}
	DiscordConf = &model.Discord{}
	err = viper.Unmarshal(DiscordConf)
	if err != nil {
		fmt.Printf("无法解析config.yaml配置文件, %v", err)
		panic(err)
	}
	fmt.Println(DiscordConf.Authorization)
}

func GetHeader() http.Header {
	header := http.Header{
		"Authorization": {DiscordConf.Authorization},
		"Content-Type":  {"application/json; charset=UTF-8"},
		"User-Agent":    {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36"},
	}
	return header
}

func init() {
	fmt.Println("conf init...")
	getDiscordConf()
}

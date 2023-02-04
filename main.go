package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sui-test-faucet/conf"
	"sui-test-faucet/model"
	"time"
)

func GetScriptCommand() []string {
	if conf.DiscordConf.ScriptModeFlag {
		return conf.DiscordConf.ScriptTextList
	} else {
		//去网络上获取
		//GetContentFromChannel()
		return []string{""}
	}
}

func Visit(url, method string, data *model.DiscordContent) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var req *http.Request
	var err error
	if method == "POST" {
		b, err := json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("post data error:%s", err.Error())
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(b))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return "", fmt.Errorf("url:%s,error:%s", url, err.Error())
	}
	req.Header = conf.GetHeader()
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("url:%s,error:%s", url, err.Error())
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	result := string(body)
	return result, nil
}

func GetContentFromChannel() string {
	channelId := conf.DiscordConf.TextChannelId
	url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?limit=100", channelId)
	result, err := Visit(url, "GET", nil)
	if err != nil {
		return ""
	}
	fmt.Println(result)
	//json.Unmarshal(body,model)
	return result
}

func chat() {
	if len(conf.DiscordConf.AuthorizationList) > 0 {
		for _, channelId := range conf.DiscordConf.ChannelIdList {
			n := fmt.Sprintf("82329451214%d33232234", rand.Intn(1000))
			data := &model.DiscordContent{}
			data.Content = GetScriptCommand()[0]
			data.Nonce = n
			data.TTS = false

			url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", channelId)
			result, _ := Visit(url, "POST", data)
			fmt.Println(result)
		}
	}
}

func main() {
	fmt.Println("main...")
	for {
		chat()
		time.Sleep(time.Duration(conf.DiscordConf.ChatIntervalTime+1) * time.Minute)
	}

}

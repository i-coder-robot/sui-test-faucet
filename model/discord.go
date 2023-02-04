package model

type Discord struct {
	Authorization     string   `yaml:"authorization"`
	AuthorizationList []string `yaml:"AuthorizationList"`
	TextChannelId     string   `yaml:"TextChannelId"`
	ChannelIdList     []string `yaml:"ChannelIdList"`
	ScriptTextList    []string `yaml:"ScriptTextList"`
	ChatIntervalTime  int      `yaml:"ChatIntervalTime"`
	ScriptModeFlag    bool     `yaml:"ScriptModeFlag"`
}

type DiscordContent struct {
	Content string `json:"content"`
	Nonce   string `json:"nonce"`
	TTS     bool   `json:"tts"`
}

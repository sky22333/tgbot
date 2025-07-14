package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	BotToken        string  `toml:"bot_token"`
	AllowedUserIDs  []int64 `toml:"allowed_user_ids"`
	AllowedGroupIDs []int64 `toml:"allowed_group_ids"`
}

var Cfg Config

func LoadConfig() {
	configBytes, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("读取配置失败: %v", err)
	}

	if err := toml.Unmarshal(configBytes, &Cfg); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}

	if Cfg.BotToken == "" {
		log.Fatal("bot_token不能为空")
	}
}

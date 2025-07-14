package main

import (
	"mybot/bot"
	"mybot/config"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化机器人
	bot.InitBot()

	// 启动机器人
	bot.Start()
}

package bot

import (
	"log"
	"time"

	tb "gopkg.in/telebot.v3"
	"mybot/config"
)

var Bot *tb.Bot

// 全局菜单和按钮
var (
	replyMenu  *tb.ReplyMarkup
	btnStatus  tb.Btn
	btnSetting tb.Btn

	inlineMenu *tb.ReplyMarkup
	btnHelp    tb.InlineButton
)

// 初始化菜单
func InitMenus() {
	replyMenu = &tb.ReplyMarkup{ResizeKeyboard: true}
	
	// 正确的按钮创建方式
	btnStatus = tb.Btn{Text: "📊 状态"}
	btnSetting = tb.Btn{Text: "⚙️ 设置"}
	
	// 设置回复键盘布局
	replyMenu.Reply(
		replyMenu.Row(btnStatus, btnSetting),
	)

	// 内联按钮
	btnHelp = tb.InlineButton{
		Unique: "help_btn",
		Text:   "帮助",
	}
	inlineMenu = &tb.ReplyMarkup{}
	inlineMenu.InlineKeyboard = [][]tb.InlineButton{{btnHelp}}
}

// 初始化机器人
func InitBot() {
	pref := tb.Settings{
		Token:  config.Cfg.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	Bot, err = tb.NewBot(pref)
	if err != nil {
		log.Fatalf("机器人启动失败: %v", err)
	}

	InitMenus()

	// 设置左侧快捷命令
	Bot.SetCommands([]tb.Command{
		{Text: "start", Description: "开始使用"},
		{Text: "help", Description: "帮助信息"},
	})

	// /start 命令
	Bot.Handle("/start", func(c tb.Context) error {
		if !CheckPermission(c) {
			return c.Send("抱歉，你无权限使用此机器人。")
		}
		text := "*欢迎使用通用机器人框架！*\n\n请点击下面的按钮获取帮助，或者使用回复键盘快捷操作。"
		
		// 先发送带回复键盘的消息
		err := c.Send(text, replyMenu, tb.ModeMarkdown)
		if err != nil {
			return err
		}
		
		// 再发送带内联键盘的消息
		return c.Send("点击下方按钮获取更多帮助：", inlineMenu)
	})

	// /help 命令
	Bot.Handle("/help", func(c tb.Context) error {
		if !CheckPermission(c) {
			return c.Send("抱歉，你无权限使用此机器人。")
		}
		help := "*帮助信息*\n- /start - 启动机器人\n- /help - 查看帮助"
		return c.Send(help, tb.ModeMarkdown)
	})

	// 行内按钮处理
	Bot.Handle(&btnHelp, func(c tb.Context) error {
		if !CheckPermission(c) {
			return c.Respond(&tb.CallbackResponse{
				Text:      "无权限",
				ShowAlert: true,
			})
		}
		return c.Edit("你点击了帮助按钮。")
	})

	// 回复按钮处理 - 使用按钮变量
	Bot.Handle(&btnStatus, func(c tb.Context) error {
		if !CheckPermission(c) {
			return c.Send("无权限")
		}
		return c.Send("状态：运行中")
	})

	Bot.Handle(&btnSetting, func(c tb.Context) error {
		if !CheckPermission(c) {
			return c.Send("无权限")
		}
		return c.Send("进入设置菜单")
	})
}

// 启动机器人
func Start() {
	log.Println("机器人启动成功")
	Bot.Start()
}
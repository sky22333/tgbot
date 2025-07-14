## 通用TG机器人脚手架

这是一个通用的 Go 语言机器人项目框架，结构清晰，易于扩展。

包含信息内联按钮，菜单按钮，回复键盘按钮。

## 文件结构说明

```
.
├── bot/
│   ├── bot.go          # 机器人核心逻辑
│   └── middleware.go   # 机器人权限中间件
├── cmd/
│   └── main.go         # 项目主入口文件，负责启动机器人
├── config/
│   └── config.go       # 加载和解析 config.toml 配置文件
├── config.toml         # 项目配置文件 (示例)
├── go.mod              # Go 模块依赖管理文件
├── go.sum              # Go 模块依赖校验文件
└── README.md           # 项目说明文档
```

## 如何运行

1.  根据 `config.toml` 的格式创建并填写您自己的配置文件。
2.  在项目根目录下执行 `go run ./cmd` 启动机器人。
3.  构建：`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o tgbot ./cmd`

## 鸣谢
[telebot v3](https://github.com/tucnak/telebot)

package util

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"os"
)

func ReadEnv() {
	// 初始化 .env 的配置，将 .env 中的配置加载到 Go 的 env 环境中
	if err := godotenv.Load(".env"); err != nil {
		golog.Println(err)
		os.Exit(1)
	}
}

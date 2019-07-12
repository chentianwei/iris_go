package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/joho/godotenv"
	"./routers"
	"./models"
	"log"
	"os"
)

func main() {

	// 初始化 .env 的配置，将 .env 中的配置加载到 Go 的 env 环境中
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	envPort := os.Getenv("PORT")
	issucc := models.GetInstance().InitDataPool()
	if !issucc {
		log.Println("init database pool failure...")
		os.Exit(1)
	} else {
		log.Println("init database pool success")
	}

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	//mvc.New(app).Handle(new(controllers.BaseController))
	routers.FindRoute(app)

	app.Run(iris.Addr(envPort))
}




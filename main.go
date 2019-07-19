package main

import (
	"os"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"./models"
	"./routers"
	_ "./util"
	"github.com/kataras/iris/middleware/logger"
)

func main() {

	envPort := os.Getenv("PORT")
	issucc := models.GetInstance().InitDataPool()
	if !issucc {
		golog.Println("init database pool failure...")
		os.Exit(1)
	} else {
		golog.Println("init database pool success")
	}

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	//mvc.New(app).Handle(new(controllers.BaseController))
	routers.FindRoute(app)

	app.Run(iris.Addr(envPort))
}





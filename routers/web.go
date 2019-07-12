package routers

import (
	"github.com/kataras/iris"
	"../controllers"
	)

func FindRoute(app *iris.Application)  {
	app.Get("/", controllers.Home)
	app.Get("/hello", controllers.Hello)
	app.Post("/edit/user", controllers.EditUser)
	app.Get("user/get", controllers.Userget)
	app.Get("user/list", controllers.UserList)
	app.Post("/user/create", controllers.CreateUser)
}
package controllers

import (
	"github.com/kataras/iris"
	"os"
	"../models"
)


func Home(ctx iris.Context) {
	Response("welcome", 200, "ok", ctx)
	envPort := os.Getenv("PORT")
	ctx.Writef(envPort)
}

func Hello(ctx iris.Context) {
	Response("Hello Iris!", 200, "ok", ctx)
}

func Userget(ctx iris.Context) {
	userId := ctx.FormValue("user_id")
	user := new(models.User)
	userinfo := user.GetUserInfo(userId)
	Response(userinfo, 200, "ok", ctx)
}

func UserList(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	limit := ctx.URLParamIntDefault("limit", 10)
	offset := (page -1) * limit
	user := new(models.User)
	users := user.GetUserList(limit, offset)
	Response(users, 200, "ok", ctx)
}

func CreateUser(ctx iris.Context) {
	userModel := new(models.User)
	n, _ := ctx.PostValueBool("name")
	if !n {
		userModel.Name = ctx.PostValue("name")
	}
	m, _ := ctx.PostValueBool("mobile")
	if !m {
		userModel.Mobile = ctx.PostValue("mobile")
	}
	userModel.CreateUser(userModel)
	Response("", 200, "ok", ctx)
}

func EditUser(ctx iris.Context) {

	//get
	ctx.Writef(ctx.FormValue("name"))
	ctx.Writef(ctx.FormValue("age"))
	ctx.Writef(ctx.FormValue("sex"))
	ctx.Writef(ctx.FormValue("grade"))
	ctx.Writef("11111")

	//form-data
	ctx.Writef(ctx.PostValue("name"))
	ctx.Writef(ctx.PostValue("age"))
	ctx.Writef(ctx.PostValue("sex"))
	ctx.Writef(ctx.PostValue("grade"))
	ctx.Writef("2222")

	Response("Hello Iris!", 200, "ok", ctx)
}
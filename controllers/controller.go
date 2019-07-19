package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/golog"
)

type BaseController struct{}

type BaseResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func Response(data interface{}, code int, msg string, ctx iris.Context) {
	var res BaseResponse
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	res.Message = msg
	res.Code = code
	res.Data = data

	jsonRes, err := json.Marshal(res)
	if err != nil {
		golog.Println("json err: ", err)
	}
	ctx.Writef(string(jsonRes))
}

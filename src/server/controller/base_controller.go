package controller

import "github.com/gin-gonic/gin"

type IController interface {
	Route(router *gin.Engine)
}

type BaseController struct {
	uid int
	ip  string
}

type ControllerResult struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

//func ApiResult(int code, string msg, mag data) string {
//	return "";
//}
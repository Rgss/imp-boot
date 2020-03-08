package v1

import (
	"github.com/Rgss/imp-boot/src/common/utils/cresponse"
	"github.com/Rgss/imp-boot/src/server/entity"
	"github.com/Rgss/imp-boot/src/server/entity/form"
	"github.com/Rgss/imp-boot/src/server/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

type AccountController struct {
}

func (this *AccountController) Route(router *gin.Engine) {

	group := router.Group("/v1")
	{
		group.GET("/account/register/", this.actionRegister)
		group.POST("/account/register/", this.actionRegister)
		group.GET("/account/login/", this.actionLogin)
		group.POST("/account/login/", this.actionLogin)
	}

	// 校验中间件
	router.Use(middleware.Authorization_middleware())
	{
		group := router.Group("/v1")
		{
			log.Printf("group: %v", group)
		}
	}
}

// @desc 注册
// @api /v1/account/register/
func (this *AccountController) actionRegister(ctx *gin.Context) {
	var account = &entity.Account{}
	if err := ctx.Bind(account); err != nil {
		ctx.JSON(200, cresponse.Fail("请填写正确账号信息！"))
		return
	}


	data := map[string]interface{} {
		"accountId": 0,
	}
	ctx.JSON(200, cresponse.Success(data))
}

// @desc 登陆
// @api /v1/account/login/
func (this *AccountController) actionLogin(ctx *gin.Context) {
	var loginForm = &form.LoginForm{}
	if err := ctx.Bind(loginForm); err != nil {
		ctx.JSON(200, cresponse.Fail("账号密码不能为空！"))
		return
	}

	rData := make(map[string]interface{})
	rData["token"] = ""
	ctx.JSON(200, cresponse.Success(rData))
}
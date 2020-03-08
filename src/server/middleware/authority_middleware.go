package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Rgss/imp-boot/src/common/utils/cresponse"
	"github.com/Rgss/imp-boot/src/server/service"
)

// 检验权限 authorization
func Authorization_middleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		return

		token := ctx.Request.FormValue("token")
		if len(token) <= 0 {
			ctx.JSON(200, cresponse.Fail("访问未授权！", 10000))
			ctx.Abort()
			return
		}

		res := service.TokenServiceFactory.CheckValid(token)
		if !res {
			ctx.JSON(200, cresponse.Fail("访问未授权！", 10000))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
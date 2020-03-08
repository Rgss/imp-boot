package webserver

import (
	"fmt"
	"github.com/Rgss/imp-boot/src/common/app"
	"github.com/Rgss/imp-boot/src/server/route"
	"github.com/gin-gonic/gin"
	"log"
)

type GinServer struct {
	gin *gin.Engine
}

func New() *GinServer {
	return &GinServer{}
}

// 启动
func (g *GinServer) Run() error {
	httpConfig := app.Core().GetHttpConfig()
	host := fmt.Sprintf("%s:%d", httpConfig.Host, httpConfig.Port)
	log.Printf("[info] HttpServer start on: %v", host)

	return g.gin.Run(host)
}

// 创建
func (g *GinServer) Create() *gin.Engine {
	debug, _ := app.Configure().Get("debug")
	if debug == nil {
		debug = false
	}

	switch debug.(bool) {
	case false:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.Use(gin.Recovery())

	// 中间件
	//router.Use()

	//router.LoadHTMLGlob()
	//router.Static()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g.gin = router

	return router
}

// 注册路由
func (g *GinServer) RegisterRouters() {
	croutes := route.RegisteRoutes()
	for _, route := range croutes {
		route.Route(g.gin)
	}
}
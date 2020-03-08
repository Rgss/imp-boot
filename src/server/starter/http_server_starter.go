package starter

import (
	"log"
	"os"
	"github.com/Rgss/imp-boot/src/common/app"
	"github.com/Rgss/imp-boot/src/server/starter/webserver"
)


// start http server
func StartHttpServer() {


	// 创建服务器
	webServer := new(webserver.GinServer)
	webServer.Create()

	// 注册路由
	webServer.RegisterRouters()

	// 启动
	err := webServer.Run()
	if err != nil {
		log.Printf("[error] httpServer ListenAndServe error, %v", err.Error())
		app.Core().Destroy()
		os.Exit(1)
	}
}



// 注册路由
func RegisterRouters() {

}
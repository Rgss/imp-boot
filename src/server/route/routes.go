package route

import (
	"github.com/Rgss/imp-boot/src/server/controller"
	"github.com/Rgss/imp-boot/src/server/controller/v1"
)

// 初始化路由，系统所有api模块都必须由此注册
func RegisteRoutes() ([]controller.IController) {

	var routes []controller.IController

	/*************************** v1 ***************************/
	routes = append(routes, new(v1.AccountController))

	/*************************** v1 ***************************/


	/*************************** v2 ***************************/

	/*************************** v2 ***************************/

	return routes
}
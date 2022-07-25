package routers

import (
	"gin_server/controllers"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	router := gin.Default()    // 创建一个默认的gin服务器
	gin.SetMode(gin.DebugMode) // 设置打包和运行模式，生产模式: gin.SetMode(gin.DebugMode)

	// -----------------------------------------

	router.GET("/", controllers.TestController) // 在全局中间件被注册之前的路由请求不会触发全局中间件

	userAuth := router.Group("/userAuth") // 路由组
	{
		userAuth.POST("/login", controllers.LoginController)
		userAuth.POST("/register", controllers.RegisterController)
	}

	// --------------------------------------------------

	return router // 返回配置完成的路由
}

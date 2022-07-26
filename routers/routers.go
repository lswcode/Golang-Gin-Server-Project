package routers

import (
	"gin_server/controllers"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	router := gin.Default()      // 创建一个默认的gin服务器
	gin.SetMode(gin.ReleaseMode) // 设置打包和运行模式，生产模式: gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Default())   // 开启跨域 要写在路由定义函数之前

	gin.DisableConsoleColor() // 将日志写入文件中
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// --------------------------------------------------
	store := cookie.NewStore([]byte("lsw"))               // 初始化一个cookie存储对象，里面的参数是自定义的密钥，默认使用内存存储，可以使用redis数据库存储
	router.Use(sessions.Sessions("GIN_SESSIONID", store)) // 启动全局session中间件，第一个参数是浏览器保存的cookie的键名，第二个参数是存储引擎

	router.GET("/userInfo", controllers.GetInfoController)

	userAuth := router.Group("/userAuth") // 路由组
	{
		userAuth.POST("/login", controllers.LoginController)
		userAuth.POST("/register", controllers.RegisterController)
	}

	// --------------------------------------------------

	return router // 返回配置完成的路由
}

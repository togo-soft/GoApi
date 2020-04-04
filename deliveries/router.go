package deliveries

import (
	"GoApi/deliveries/GoVersion"
	"GoApi/middleware"
	"GoApi/utils"
	"github.com/gin-gonic/gin"
)

// Run 程序启动的入口
func Run() {
	var conf = utils.GetConfig()
	var router = InitRouter()
	//测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//关闭debug
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//服务运行
	_ = router.Run(conf.Server)
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	//开启允许跨域
	router.Use(middleware.CORS())

	//GoVersion路由
	var version = router.Group("/api/go/version")
	{
		//获取
		version.GET("/get", GoVersion.Get)
		//更新
		version.GET("/update", GoVersion.Update)
	}
	return router
}

package main

import (
	"GinDemo/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/hello", func(context *gin.Context) {
	//	log.Println("start gin hello")
	//	context.JSON(200,gin.H{
	//		"code":200,
	//		"success": true,
	//	})
	//})
	//// 指定地址和端口号
	//router.Run("localhost:9090")

	// 使用非匿名函数处理
	router.GET("/hello", hello)
	//加载本地的html文件
	router.LoadHTMLFiles("tmeplate/*")
	// 路由组
	user := router.Group("/user")
	{
		// 请求参数放到路径上
		user.GET("/get/:id/:username", controller.QueryById)
		user.GET("/query",controller.QueryParam)
		user.POST("/insert",controller.InsertNewUser)
		// 跳转到html
		user.GET("/from", controller.RenderForm)
		user.POST("/form/post", controller.PostFrom)
		// 可以自己添加其他的，一个请求的路径对应一个函数
	}
	file := router.Group("/file")
	{
		// 跳转上传文件页面
		file.GET("/view",controller.ReaderViewHtml)
		// 根据表单上传
		file.POST("/insert",controller.FroomUpload)
		//file.POST("/multiUpload",controller.MultiUpload)
		//// base64上传
		//file.POST("/upload",controller.Base64Upload)
	}
	router.Run(":9090")
}

// 使用非匿名函数处理
func hello(context *gin.Context)  {
	println("hello start")

}

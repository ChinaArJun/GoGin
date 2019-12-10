package controller

import "github.com/gin-gonic/gin"

func QueryById(context *gin.Context)  {
	// 获取请求参数
	id := context.GetString("id")
	name := context.GetString("name")

	println("id:", id, "name:", name)
}

func QueryParam(context *gin.Context)  {

}

// 新建用户
func InsertNewUser(context *gin.Context)  {

}

// 跳转到html
func RenderForm(context *gin.Context) {
	context.HTML(200,"insertUser.html", gin.H{})
}

func PostFrom(contxt *gin.Context) {

}
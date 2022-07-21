package controllers

import (
	"github.com/gin-gonic/gin"
)

func LoginController(ctx *gin.Context) { // gin把request和response都封装到了gin.Context
	ctx.String(200, "请求成功")
}

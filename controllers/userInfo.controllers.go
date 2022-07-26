package controllers

import (
	"fmt"
	"gin_server/db"
	"gin_server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetInfoController(c *gin.Context) { // gin把request和response都封装到了gin.Context中
	session := sessions.Default(c)
	account := session.Get("sessionid")
	fmt.Println("sessionid:-------", account)
	if account == nil {
		// c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取用户数据成功"})  // 如果前面写了这一条，下面的c.JSON函数会被忽略
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "sessionID无效"}) // c.JSON表示以JSON格式返回数据给前端，在一个处理函数中只有第一次返回数据有效，后面再返回数据会被忽略
		return
	}
	var userInfo models.User
	count := 0
	db.Db.Where("account = ?", account).First(&userInfo).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userInfo中
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "sessionID无效"})
		fmt.Println("sessionID无效")
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取用户数据成功", "userName": userInfo.Username})
}

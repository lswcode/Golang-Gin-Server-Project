package controllers

import (
	"fmt"
	"gin_server/db"
	"gin_server/models"
	"net/http"

	"gin_server/utils"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) { // gin把request和response都封装到了gin.Context中
	c.String(200, "请求成功")
	xiaobai := models.User{Username: "小白", Password: "123456"}
	db.Db.Create(&xiaobai)
}

func LoginController(c *gin.Context) {
	var loginJson models.User
	err := c.ShouldBindJSON(&loginJson) // 将request的body中的数据，自动按照json格式解析到结构体
	fmt.Println("loginJson---------------------------------------------------", loginJson)
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userDb models.User
	count := 0
	db.Db.Where("account = ?", loginJson.Account).Find(&userDb).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userDb中
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{"code": "201", "msg": "该账号不存在，请重新输入"})
		return
	}

	if afterMD5 := utils.Md5(loginJson.Password); afterMD5 != userDb.Password {
		c.JSON(http.StatusOK, gin.H{"code": "202", "msg": "密码不正确，请重新输入"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "登录成功"})
	}
}

func RegisterController(c *gin.Context) {
	var registerJson models.User
	err := c.ShouldBindJSON(&registerJson) // 将request的body中的数据，自动按照json格式解析到结构体
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userDb models.User
	count := 0
	db.Db.Where("username = ?", registerJson.Username).Find(&userDb).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userDb中
	if count > 0 {
		fmt.Println("用户名已被注册")
		c.JSON(http.StatusOK, gin.H{"code": "201", "msg": "用户名已被注册，请重新输入"})
		return
	}
	count = 0
	db.Db.Where("account = ?", registerJson.Account).Find(&userDb).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userDb中
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"code": "202", "msg": "账号已被注册，请重新输入"})
		return
	}

	fmt.Println("注册原密码", registerJson.Password)
	afterMD5 := utils.Md5(registerJson.Password)
	fmt.Println("加密后的密码", afterMD5)
	registerJson.Password = afterMD5

	err = db.Db.Create(&registerJson).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

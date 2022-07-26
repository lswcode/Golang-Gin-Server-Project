package controllers

import (
	"fmt"
	"gin_server/db"
	"gin_server/models"
	"net/http"

	"gin_server/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginJson models.User
	err := c.ShouldBindJSON(&loginJson) // 将request的body中的数据，自动按照json格式解析到结构体
	fmt.Println("loginJson---------------------------------------------------", loginJson)
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userInfo models.User
	count := 0
	db.Db.Where("account = ?", loginJson.Account).Find(&userInfo).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userInfo中
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{"code": "201", "msg": "该账号不存在，请重新输入"})
		return
	}

	if afterMD5 := utils.Md5(loginJson.Password); afterMD5 != userInfo.Password {
		c.JSON(http.StatusOK, gin.H{"code": "202", "msg": "密码不正确，请重新输入"})
	} else {
		// cookie, err := c.Cookie("key_cookie") // 先判断当前请求是否已经携带cookie
		// if err != nil {
		// 	cookie = "NotSet"
		// 	c.SetCookie("gin_cookie", "test1", 3600, "/", // 参数依次为 cookie的键，值，有效期，cookie所在目录，cookie作用域(本地调试时设置为localhost，生成环境设置为网站域名)，是否只支持https(true表示只能是https)，是否允许被JS操作
		// 		"localhost", false, true)
		// 	// c.SetCookie设置cookie后，当浏览器获取到这个请求响应后，会自动添加这里设置的cookie

		// }
		// fmt.Printf("cookie的值是： %s\n", cookie)

		// ---一般都是要配合session使用，使用第三方库session来设置cookie和session，就不需要上面单独设置cookie的代码了----------------------------------------------------------------------
		session := sessions.Default(c)             // 默认格式，创建session时必须写
		session.Set("sessionid", userInfo.Account) // 将account作为sessionID保存到浏览器的cookie中，之后就可以根据cookie解析出用户账号了(也可以专门创建一个表，使用uuid等独一无二的ID作为sessionID，然后再将sessionID和登陆成功的用户数据一起保存到数据库中，之后就可以根据这个sessionID获取用户信息了)
		err := session.Save()                      // 保存session
		if err != nil {
			fmt.Println("设置session报错", err)
		}

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
	var userInfo models.User
	count := 0
	db.Db.Where("username = ?", registerJson.Username).Find(&userInfo).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userInfo中
	if count > 0 {
		fmt.Println("用户名已被注册")
		c.JSON(http.StatusOK, gin.H{"code": "201", "msg": "用户名已被注册，请重新输入"})
		return
	}
	count = 0
	db.Db.Where("account = ?", registerJson.Account).Find(&userInfo).Count(&count) // 查询指定的字段，会将查询到的结果保存在&userInfo中
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

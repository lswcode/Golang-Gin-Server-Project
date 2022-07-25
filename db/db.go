package db

import (
	"fmt"

	"gin_server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	// 第一个参数是连接的数据库类型，第二个参数依次是用户名，密码，数据库名，后面的使用默认的即可
	// 返回两个参数，第二个参数就是连接失败时返回的错误
	database, err := gorm.Open("mysql", "root:Lswmysql123.@/gin_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm数据库连接失败")
		panic(err) // 发送错误时报错并终止程序
	}
	Db = database // 将数据库对象赋值给全局变量Db，这样其它文件就可以导入这个变量并使用了

	Db.LogMode(true)            // 开启日志记录，会打印所有gorm对应的sql语句
	Db.DB().SetMaxOpenConns(10) // 设置数据库的最大连接数
	Db.DB().SetMaxIdleConns(50) // 设置数据库的最大空闲数

	Db.AutoMigrate(&models.User{}) // 可以用逗号隔开，创建多个表格
	// AutoMigrate表示自动迁移格式创建表格，可以保护原数据
}

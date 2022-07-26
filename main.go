package main

import (
	"fmt"
	"gin_server/routers"
)

func main() {

	router := routers.RouterInit()

	err := router.Run(":8099")
	if err != nil {
		fmt.Printf("服务器开启失败, err:%v\n", err)
	}
}

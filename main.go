package main

import (
	"gin_server/routers"
)

func main() {
	router := routers.RouterInit()

	router.Run(":8060")
}

package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main(){

	r := routers.InitRouter()

	r.Run(fmt.Sprint(":",setting.HTTPPort))
}

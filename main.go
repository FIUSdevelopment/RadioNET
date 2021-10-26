package main

import (
	_ "radioNet/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/database","database")
	beego.Run()
}


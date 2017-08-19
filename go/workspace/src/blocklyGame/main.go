package main

import (
	_ "blocklyGame/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/en", "en")
	beego.Run()
}


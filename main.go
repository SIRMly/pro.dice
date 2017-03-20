package main

import (
	"github.com/astaxie/beego"
	_ "github.com/ruooooooli/RollTheDice/routers"
)

func main() {
	beego.SetStaticPath("/static/audio", "/static/audio")
	beego.Run()
}

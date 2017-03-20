package routers

import (
	"github.com/astaxie/beego"
	"github.com/ruooooooli/RollTheDice/controllers"
)

func init() {
	beego.Router("/code/:name([\\w-]+)", &controllers.IndexController{}, "get:Image")

	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/open/:name([\\w-]+)", &controllers.IndexController{}, "get:Open")

	beego.Router("/game/:name([\\w-]+)", &controllers.IndexController{}, "get:Game")
	beego.Router("/join/:name([\\w-]+)", &controllers.IndexController{}, "get:Join")
}

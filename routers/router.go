package routers

import (
	"github.com/saburo90/Wechat-Management-Go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

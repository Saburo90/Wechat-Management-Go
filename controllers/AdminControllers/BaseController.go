package AdminControllers

import (
	
)

type BaseController struct {
	beego.Controller
	TplTheme 			string
	TplStatic 			string
	AdminId 			int
}

func (this *BaseController) Prepare() {
	var ok bool
	AdminId = this.GetSession("AdminId")
	// affirm the value of "AdminId"
	this.AdminId, ok = AdminId.(int)
}
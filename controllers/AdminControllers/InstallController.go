package AdminControllers

import (
	
)

type InstallController struct {
	beego.Controller
}

type installForm struct {
	Host 		string 		`form:"host"`
	Port 		int 		`form:"port"`
	Database 	string		`form:"database"`
	Prefix		string		`form:"prefix"`
	Username	string		`form:"username"`
	Password	string		`form:"password"`
}

func (this *InstallController) Install() {
	// http status 302 indicate url redirect
	if helper.IsInstalled {
		this.Redirect("/admin/login", 302)
	}
}
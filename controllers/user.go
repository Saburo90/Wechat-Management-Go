package controllers

import (
	
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	if c.Ctx.Request.Method == "POST" {

	}
	c.TplName = c.controllerName + "/login.html"
}
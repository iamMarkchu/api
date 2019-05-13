package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) JsonReturn(message string, result interface{}, code int) {
	c.Data["json"] = map[string]interface{}{
		"message": message,
		"result": result,
		"code": code,
	}
	c.ServeJSON(true)
}

package controllers

import (
	"net/http"
)

type UserController struct {
	ApiController
}

func (c *UserController) URLMapping()  {
	c.Mapping("Index", c.Index)
}

// @router / [get]
func (c *UserController) Index()  {
	c.JsonReturn("我是用户中心", Result{"userId": c.UserId,}, http.StatusOK)
}

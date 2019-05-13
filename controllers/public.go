package controllers

import (
	"api/controllers/requests"
	"api/services"
	"net/http"
	"strconv"
)

type PublicController struct {
	ApiController
}

func (c *PublicController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// Title login
// Description 登录接口
// Param username formData string true "登录用户名"
// Param password formData string true "登录密码“
// @router /login [post]
func (c *PublicController) Login() {
	r := requests.LoginRequest{}
	if err := c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误: " + err.Error(), "", http.StatusNotFound)
		return
	}
	userService := services.NewUserService()
	token, err := userService.Login(r)
	if err != nil {
		c.JsonReturn("登录错误:" + err.Error(), "", http.StatusForbidden)
		return
	}
	result := map[string]string{
		"token": token,
	}
	c.JsonReturn("登录成功!", result, http.StatusOK)
}

// Title register
// Description 注册接口
// Param username formData string true "注册用户名"
// Param password formData string true "密码"
// Param repassword formData string true "确认"
// Param email formData string true "邮箱地址"
// @router /register [post]
func (c *PublicController) Register() {
	r := requests.RegisterRequest{}
	if err := c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误: " + err.Error(), "", http.StatusNotFound)
		return
	}
	userService := services.NewUserService()
	id, err := userService.Register(r)
	if err != nil {
		c.JsonReturn("注册错误:" + err.Error(), "", http.StatusForbidden)
		return
	}
	result := map[string]string{
		"id": strconv.Itoa(id),
	}
	c.JsonReturn("注册成功", result, http.StatusOK)
}

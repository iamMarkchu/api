package controllers

import (
	"api/controllers/requests"
	. "api/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/http"
)

var (
	result Result // 返回值
)

type ApiController struct {
	beego.Controller
}

type HttpCode int

type JsonResponse struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type Result map[string]interface{}

func (c *ApiController) JsonReturn(message string, result interface{}, code int) {
	c.Data["json"] = JsonResponse{
		Message: message,
		Result:  result,
		Code:    HttpCode(code),
	}
	c.ServeJSON(true)
}

func (c *ApiController) ValidateRequest(r requests.Request) {
	var (
		valid = validation.Validation{}
		err   error
		isValid bool
	)

	if category, ok := r.(requests.CategoryStoreRequest); ok {
		err = c.ParseForm(&category)
		isValid, _ = valid.Valid(&category)
	} else if article, ok := r.(requests.ArticleStoreRequest); ok {
		err = c.ParseForm(&article)
		isValid, _ = valid.Valid(&article)
	}

	if err != nil {
		c.JsonReturn("解析参数错误: "+err.Error(), "", http.StatusBadRequest)
	}
	if !isValid {
		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
	}
}

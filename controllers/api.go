package controllers

import (
	. "api/helpers"
	"api/helpers/cache"
	"github.com/astaxie/beego"
	bcache "github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"strconv"
	"strings"
)

var (
	result  Result // 返回值
	valid   = validation.Validation{}
	err     error
	isValid bool
)

type ApiController struct {
	beego.Controller
	UserId int  // token所携带的用户
	CacheInstance bcache.Cache  // 初始化缓存类
}

type HttpCode int

type JsonResponse struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type Result map[string]interface{}

func (c *ApiController) Prepare() {
	c.CacheInstance = cache.GetCacheInstance()
	token := strings.TrimPrefix(c.Ctx.Input.Header("Authorization"), "Bearer ")
	c.UserId = bcache.GetInt(c.CacheInstance.Get(MD5(token)))
	logs.Info("USERID:", c.UserId)
}

func (c *ApiController) JsonReturn(message string, result interface{}, code int) {
	c.Data["json"] = JsonResponse{
		Message: message,
		Result:  result,
		Code:    HttpCode(code),
	}
	c.ServeJSON(true)
}

//func (c *ApiController) ValidateRequest(r requests.Request) {
//	var (
//		valid = validation.Validation{}
//		err   error
//		isValid bool
//	)
//
//	if category, ok := r.(requests.CategoryStoreRequest); ok {
//		err = c.ParseForm(&category)
//		isValid, _ = valid.Valid(&category)
//	} else if article, ok := r.(requests.ArticleStoreRequest); ok {
//		err = c.ParseForm(&article)
//		isValid, _ = valid.Valid(&article)
//	}
//
//	if err != nil {
//		c.JsonReturn("解析参数错误: "+err.Error(), "", http.StatusBadRequest)
//	}
//	if !isValid {
//		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
//	}
//}

func (c *ApiController) getId() int {
	var id int
	id, err = strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		return 0
	}
	return id
}

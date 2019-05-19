package controllers

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/models"
	"api/services"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type CategoryController struct {
	ApiController
}

func (c *CategoryController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
}

// @router / [get]
func (c *CategoryController) Index() {

}

// @router / [post]
func (c *CategoryController) Store() {
	var (
		r               = requests.CategoryStoreRequest{}
		categoryService *services.CategoryService
		category        models.Category
		isSuccess       bool
	)
	// todo 能否使用一个基类方法统一验证请求参数？
	// c.ValidateRequest(r)
	if err = c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误: "+err.Error(), "", http.StatusBadRequest)
	}
	logs.Info(r)
	isValid, _ = valid.Valid(&r)
	if !isValid {
		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
	}
	categoryService = services.NewCategoryService()
	category, isSuccess = categoryService.Store(r, c.UserId)
	if isSuccess {
		c.JsonReturn("创建类别接口", category, http.StatusOK)
	}
	c.JsonReturn("创建类别失败", category, http.StatusBadRequest)
}

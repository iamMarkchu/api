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
	categoryService *services.CategoryService
	categoryModel   *models.Category
}

func (c *CategoryController) Prepare() {
	c.ApiController.Prepare()
	c.categoryService = services.NewCategoryService()
}

func (c *CategoryController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
}

// @router / [get]
// Description 类别列表接口
func (c *CategoryController) Index() {
	var categories []*models.Category
	categories, err = c.categoryService.GetList()
	if err != nil {
		c.JsonReturn("类别列表接口报错:"+err.Error(), "", http.StatusBadRequest)
		return
	}
	c.JsonReturn("类别列表接口", categories, http.StatusOK)
}

// @router / [post]
// Description 保存类别接口
func (c *CategoryController) Store() {
	var (
		r               = requests.CategoryStoreRequest{}
		categoryService *services.CategoryService
		category        *models.Category
		err             error
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
	category, err = categoryService.Store(r, c.UserId)
	if err != nil {
		c.JsonReturn("创建类别失败:" + err.Error(), category, http.StatusBadRequest)
	} else {
		c.JsonReturn("创建类别接口", category, http.StatusOK)
	}

}

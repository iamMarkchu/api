package controllers

import (
	"api/controllers/requests"
	"api/helpers"
	"api/models"
	"github.com/astaxie/beego/validation"
	"net/http"
)

type ArticleController struct {
	ApiController
}

func (c *ArticleController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
}

// @router / [get]
func (c *ArticleController) Index() {
	var articles []models.Article
	c.JsonReturn("文章列表接口", articles, http.StatusOK)
}

// @router / [post]
func (c *ArticleController) Store() {
	r := requests.ArticleStoreRequest{}
	if err := c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误: " + err.Error(), "", http.StatusNotFound)
		return
	}

	valid := validation.Validation{}
	isValid, _ := valid.Valid(&r)
	if !isValid {
		c.JsonReturn("参数不符合要求!", helpers.GetErrorMap(valid.Errors), http.StatusNotFound)
		return
	}

	c.JsonReturn("新建文章接口", "", http.StatusOK)
}

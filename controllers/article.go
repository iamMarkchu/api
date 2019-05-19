package controllers

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/helpers/page"
	"api/models"
	"api/services"
	"net/http"
	"strconv"
)

type ArticleController struct {
	ApiController
}

func (c *ArticleController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
}

// Title articles/index
// Param status formData int8 false "状态"
// Param page   formData int  false "页数"
// Param limit  formData int  false "每页数量"
// @router / [get]
func (c *ArticleController) Index() {
	var (
		articleService = services.NewArticleService()
		r              requests.ArticleIndexRequest
		err            error
		queryMap       map[string]string
		result         page.Page
	)
	if err = c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误:"+err.Error(), "", http.StatusBadRequest)
	}

	queryMap = map[string]string{
		"Status": strconv.Itoa(int(r.Status)),
	}
	result, err = articleService.GetArticles(queryMap, r.Page, r.Limit)
	if err != nil {
		c.JsonReturn("文章列表报错:"+err.Error(), "", http.StatusBadRequest)
		return
	}
	c.JsonReturn("文章列表接口", result, http.StatusOK)
}

// @router / [post]
func (c *ArticleController) Store() {
	// todo 能否使用一个基类方法统一验证请求参数？
	// c.ValidateRequest(requests.ArticleStoreRequest{})
	var r = requests.ArticleStoreRequest{}
	if err = c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	isValid, _ = valid.Valid(&r)
	if !isValid {
		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
		return
	}

	var (
		articleService = services.NewArticleService()
		article        = models.NewArticle()
		isSuccess      bool
	)
	article, isSuccess = articleService.Store(r, c.UserId)
	if isSuccess {
		c.JsonReturn("创建文章成功!", article, http.StatusOK)
		return
	}
	c.JsonReturn("创建文章失败!", article, http.StatusBadRequest)
}

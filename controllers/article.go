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
	articleService *services.ArticleService
	articleModel   *models.Article
}

func (c *ArticleController) Prepare() {
	c.ApiController.Prepare()
	c.articleService = services.NewArticleService()
	c.articleModel = models.NewArticle()
}

func (c *ArticleController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
	c.Mapping("Update", c.Update)
	c.Mapping("Show", c.Show)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Publish", c.Publish)
	c.Mapping("Reset", c.Reset)
}

// Title articles/index
// Param status formData int8 false "状态"
// Param page   formData int  false "页数"
// Param limit  formData int  false "每页数量"
// @router / [get]
func (c *ArticleController) Index() {
	var (
		r        requests.ArticleIndexRequest
		queryMap map[string]string
		result   page.Page
	)
	if err = c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误:"+err.Error(), "", http.StatusBadRequest)
	}
	queryMap = map[string]string{
		"Status": strconv.Itoa(int(r.Status)),
	}
	result, err = c.articleService.GetArticles(queryMap, r.Page, r.Limit)
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
		c.JsonReturn("解析参数错误: "+err.Error(), nil, http.StatusBadRequest)
		return
	}
	isValid, _ = valid.Valid(&r)
	if !isValid {
		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
		return
	}

	var isSuccess bool
	c.articleModel, isSuccess = c.articleService.Store(r, c.UserId)
	if isSuccess {
		c.JsonReturn("创建文章成功!", c.articleModel, http.StatusOK)
		return
	}
	c.JsonReturn("创建文章失败!", models.NewArticle(), http.StatusBadRequest)
}

// @router /:id [put]
func (c *ArticleController) Update() {
	var (
		id int
		r  requests.ArticleUpdateRequest
	)
	id = c.getId()
	if err != nil {
		c.JsonReturn("获取文章id失败", "", http.StatusBadRequest)
		return
	}

	if err = c.ParseForm(&r); err != nil {
		c.JsonReturn("解析参数错误", "", http.StatusBadRequest)
		return
	}

	isValid, _ = valid.Valid(&r)
	if !isValid {
		c.JsonReturn("参数不符合要求!", GetErrorMap(valid.Errors), http.StatusBadRequest)
		return
	}
	c.articleModel, err = c.articleService.Update(r, id)
	if err != nil {
		c.JsonReturn("文章更新失败:"+err.Error(), "", http.StatusBadRequest)
		return
	}
	c.JsonReturn("文章更新成功!", c.articleModel, http.StatusBadRequest)
}

// @router /:id [get]
func (c *ArticleController) Show() {
	var id int
	id = c.getId()
	if err != nil {
		c.JsonReturn("获取文章id失败", "", http.StatusBadRequest)
		return
	}
	c.articleModel, err = c.articleService.FetchOne(id)
	if err == nil {
		c.JsonReturn("获取文章成功", c.articleModel, http.StatusOK)
		return
	}
	c.JsonReturn("获取文章失败: "+err.Error(), "", http.StatusBadRequest)
}

// @router /:id [delete]
func (c *ArticleController) Delete() {
	var id int
	id = c.getId()
	if err != nil {
		c.JsonReturn("获取文章id失败", "", http.StatusBadRequest)
		return
	}
	_, err = c.articleService.DeleteById(id)
	if err == nil {
		c.JsonReturn("文章删除成功!", "", http.StatusOK)
		return
	}
	c.JsonReturn("文章删除失败:"+err.Error(), "", http.StatusBadRequest)
}

// @router /:id/publish [put]
func (c *ArticleController) Publish() {
	var id int
	id = c.getId()
	if err != nil {
		c.JsonReturn("获取文章id失败", "", http.StatusBadRequest)
	}
	_, err = c.articleService.PublishById(id)
	if err == nil {
		c.JsonReturn("文章发布成功!", "", http.StatusOK)
		return
	}
	c.JsonReturn("文章发布失败:"+err.Error(), "", http.StatusBadRequest)
}

// @router /:id/reset [put]
func (c *ArticleController) Reset()  {
	var id int
	id = c.getId()
	if err != nil {
		c.JsonReturn("获取文章id失败", "", http.StatusBadRequest)
	}
	_, err = c.articleService.ResetById(id)
	if err == nil {
		c.JsonReturn("文章重置成功!", "", http.StatusOK)
		return
	}
	c.JsonReturn("文章重置失败:"+err.Error(), "", http.StatusBadRequest)
}

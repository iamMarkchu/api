package controllers

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/helpers/cache"
	"api/models"
	"api/services"
	"fmt"
	bcache "github.com/astaxie/beego/cache"
	"net/http"
	"strings"
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
	c.ValidateRequest(r)
	articleService := services.NewArticleService()
	// 获取UserId
	bm := cache.GetCacheInstance()
	token := strings.TrimPrefix(c.Ctx.Input.Header("Authorization"), "Bearer ")
	userId := bcache.GetInt(bm.Get(MD5(token)))
	fmt.Println("userId:", userId)
	article, isSuccess := articleService.Store(r, userId)
	if isSuccess {
		c.JsonReturn("新建文章接口", article, http.StatusOK)
	}
	c.JsonReturn("新建文章失败", models.Article{}, http.StatusBadRequest)
}

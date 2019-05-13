package controllers

import (
	"api/models"
	"net/http"
	"strconv"
	"time"
)

type ArticleController struct {
	ApiController
}

func (c *ArticleController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// @router / [get]
func (c *ArticleController) Index() {
	var articles []models.Article
	for i := 1; i < 10; i++ {
		tmp := models.Article{
			Id:           i,
			Title:        "Ttile" + strconv.Itoa(i),
			Description:  "11111111111",
			Author:       "chukui",
			CreatedAt:    time.Now(),
			CreatedAtStr: time.Now().Format("2006-01-02 15:04:05"),
		}
		articles = append(articles, tmp)
	}
	c.JsonReturn("文章列表接口", articles, http.StatusOK)
}

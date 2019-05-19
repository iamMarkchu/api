package services

import (
	"api/controllers/requests"
	"api/helpers/page"
	"api/models"
	"github.com/astaxie/beego/logs"
)

type ArticleService struct {
}

func (s *ArticleService) Store(r requests.ArticleStoreRequest, userId int) (*models.Article, bool) {
	var (
		article   = models.NewArticle()
		isSuccess = true
	)
	article.Author = &models.User{Id: userId}
	article.Category = &models.Category{Id: r.CategoryId}
	article.Title = r.Title
	article.Description = r.Description
	article.ImageUrl = r.ImageUrl
	article.Status = models.StatusInActive
	_, err := article.Store()
	if err != nil {
		logs.Info("插入失败:", err.Error())
		isSuccess = false
	}
	return article, isSuccess
}

func (s *ArticleService) GetArticles(queryMap map[string]string, pageNo int, limit int) (page.Page, error) {
	var (
		article  = models.NewArticle()
		articles []*models.Article
		count    int64
		err      error
	)
	if pageNo == 0 {
		pageNo = 1
	}
	if limit == 0 {
		limit = 20
	}
	articles, count, err = article.GetList(queryMap, pageNo, limit)
	return page.PageUtil(int(count), pageNo, limit, articles), err
}

func NewArticleService() ArticleService {
	return ArticleService{}
}

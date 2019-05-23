package services

import (
	"api/controllers/requests"
	"api/helpers/page"
	"api/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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

func (s *ArticleService) Update(r requests.ArticleUpdateRequest, articleId int) (*models.Article, error) {
	var (
		article = models.NewArticle()
		err     error
		o       = orm.NewOrm()
	)
	article.Id = articleId
	if err = o.Read(article); err == nil {
		article.Title = r.Title
		article.Description = r.Description
		article.Category.Id = r.CategoryId
		article.ImageUrl = r.ImageUrl
		if _, err := o.Update(article); err == nil {
			article.FormatDatetime()
			logs.Info("更新文章:", article)
			return article, nil
		}
	}
	return nil, err
}

func (s *ArticleService) FetchOne(articleId int) (*models.Article, error) {
	var (
		article = models.NewArticle()
		err     error
		o       = orm.NewOrm()
	)
	article.Id = articleId
	if err  = o.Read(article); err == nil{
		article.FormatDatetime()
		_, _ = o.LoadRelated(article,"Author")
		article.Author.FormatDatetime()
		_, _ = o.LoadRelated(article,"Category")
		article.Category.FormatDatetime()
		return article, nil
	}
	return nil, err
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

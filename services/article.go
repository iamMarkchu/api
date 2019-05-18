package services

import (
	"api/controllers/requests"
	"api/models"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type ArticleService struct {
}

func (s *ArticleService) Store(request requests.ArticleStoreRequest, userId int) (models.Article, bool) {
	articleModel := models.Article{}
	articleModel.Title = request.Title
	articleModel.Description = request.Description
	articleModel.ImageUrl = request.ImageUrl
	fmt.Println("article:", articleModel)
	author := models.User{Id:userId}
	articleModel.Author = &author
	articleModel.Status = models.StatusInActive
	_, err := articleModel.Store()
	var isSuccess = true
	if err != nil {
		logs.Info("插入失败:", err.Error())
		isSuccess = false
	}
	return articleModel, isSuccess
}

func NewArticleService() ArticleService {
	return ArticleService{}
}

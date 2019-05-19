package services

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/models"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type CategoryService struct {
}

func (s *CategoryService) Store(r requests.CategoryStoreRequest, userId int) (models.Category, bool) {
	var (
		categoryModel models.Category
		err           error
		isSuccess     = true
	)
	categoryModel.Name = r.Name
	categoryModel.ParentCate = &models.Category{Id: r.ParentCate}
	categoryModel.Author = &models.User{Id: userId}
	categoryModel.Status = models.StatusNormal
	fmt.Println(categoryModel)
	_, err = categoryModel.Store()
	go CheckError(err, "创建Category报错:")

	if err != nil {
		logs.Info("插入失败:", err.Error())
		isSuccess = false
	}
	return categoryModel, isSuccess
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

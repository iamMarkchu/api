package services

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type CategoryService struct {
	category   *models.Category
	dbInstance orm.Ormer
	query      orm.QuerySeter
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

func (s *CategoryService) GetList() ([]*models.Category, error) {
	var categories []*models.Category
	_, err = s.dbInstance.QueryTable(s.category).All(&categories)
	if err == nil {
		for _, category := range categories {
			category.FormatDatetime()
		}
		return categories, nil
	}
	return nil, err
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		category:   models.NewCategory(),
		dbInstance: orm.NewOrm(),
	}
}

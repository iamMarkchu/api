package services

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/models"
	"errors"
	"github.com/astaxie/beego/orm"
)

type CategoryService struct {
	category   *models.Category
	dbInstance orm.Ormer
	query      orm.QuerySeter
}

func (s *CategoryService) Store(r requests.CategoryStoreRequest, userId int) (*models.Category, error) {
	var (
		categoryModel models.Category
		err           error
	)
	categoryModel.Name = r.Name
	categoryModel.ParentId = r.ParentCate
	categoryModel.AuthorId = userId
	categoryModel.Status = models.StatusNormal
	// 如果 r.ParentCate指定了类别，需要类别是否存在
	if r.ParentCate > 0 {
		if err = s.CheckExist(r.ParentCate); err != nil {
			return nil, errors.New("父类别不存在!")
		}
	}
	_, err = categoryModel.Store()
	go CheckError(err, "创建Category报错:")

	return &categoryModel, err
}

func (s *CategoryService) GetList() ([]*models.Category, error) {
	var categories []*models.Category
	_, err = s.dbInstance.QueryTable(s.category).All(&categories)
	if err == nil {
		return categories, nil
	}
	return nil, err
}

func (s *CategoryService) CheckExist(i int) error {
	var (
		parentCategory = &models.Category{Id:i}
		err error
	)
	err = s.dbInstance.Read(parentCategory)
	if err == orm.ErrNoRows {
		return errors.New("找不到指定数据")
	} else if err == orm.ErrMissPK {
		return errors.New("缺少主键")
	} else {
		return nil
	}
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		category:   models.NewCategory(),
		dbInstance: orm.NewOrm(),
	}
}

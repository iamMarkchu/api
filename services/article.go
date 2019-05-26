package services

import (
	"api/controllers/requests"
	"api/helpers/page"
	"api/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ArticleService struct {
	article    *models.Article
	dbInstance orm.Ormer
	query      orm.QuerySeter
}

func (s *ArticleService) Store(r requests.ArticleStoreRequest, userId int) (*models.Article, bool) {
	// 保存值
	s.article.Author = &models.User{Id: userId}
	s.article.Category = &models.Category{Id: r.CategoryId}
	s.article.Title = r.Title
	s.article.Description = r.Description
	s.article.ImageUrl = r.ImageUrl
	s.article.Status = models.StatusInActive

	// 插入数据库
	if _, err = s.dbInstance.Insert(s.article); err == nil {
		// 格式化时间并载入author,category关系
		s.article.FormatDatetime()
		_, err = s.dbInstance.LoadRelated(s.article, "Author")
		if err != nil {
			logs.Error("载入Author关系报错:", err)
		}
		s.article.Author.FormatDatetime()
		if s.article.Category.Id != 0 {
			_, err = s.dbInstance.LoadRelated(s.article, "Category")
			if err != nil {
				logs.Error("载入Category关系报错:", err)
			}
			s.article.Category.FormatDatetime()
		}
		return s.article, true
	}
	return nil, false
}

func (s *ArticleService) GetArticles(queryMap map[string]string, pageNo int, limit int) (page.Page, error) {
	var (
		articles []*models.Article
		count    int64
	)
	// 设置默认页数以及每页数量
	if pageNo == 0 {
		pageNo = 1
	}
	if limit == 0 {
		limit = 20
	}
	s.query = s.dbInstance.QueryTable(s.article)
	// 判断status是否存在
	if status, ok := queryMap["Status"]; ok && status != "0" {
		s.query = s.query.Filter("status", status)
	}
	_, err = s.query.Limit(limit, (pageNo-1)*limit).RelatedSel().All(&articles)
	count, err = s.query.Count()

	return page.PageUtil(int(count), pageNo, limit, articles), err
}

func (s *ArticleService) Update(r requests.ArticleUpdateRequest, articleId int) (*models.Article, error) {
	s.article.Id = articleId
	// 查询文章
	if err = s.GetArticle(s.article); err == nil {
		// 赋值
		s.article.Title = r.Title
		s.article.Description = r.Description
		s.article.Category.Id = r.CategoryId
		s.article.ImageUrl = r.ImageUrl
		// 更新操作
		if _, err := s.dbInstance.Update(s.article); err == nil {
			s.article.FormatDatetime()
			logs.Info("更新文章:", s.article)
			return s.article, nil
		}
	}
	return nil, err
}

func (s *ArticleService) FetchOne(articleId int) (*models.Article, error) {
	s.article.Id = articleId
	// 查询文章
	if err = s.GetArticle(s.article); err == nil {
		s.article.FormatDatetime()
		_, _ = s.dbInstance.LoadRelated(s.article, "Author")
		s.article.Author.FormatDatetime()
		_, _ = s.dbInstance.LoadRelated(s.article, "Category")
		s.article.Category.FormatDatetime()
		return s.article, nil
	}
	return nil, err
}

func (s *ArticleService) DeleteById(articleId int) (*models.Article, error) {
	s.article.Id = articleId
	if err = s.GetArticle(s.article); err == nil {
		if s.article.Status != models.StatusNormal {
			s.article.Status = models.StatusBanned
			if _, err = s.dbInstance.Update(s.article, "Status"); err == nil {
				s.article.FormatDatetime()
				return s.article, nil
			}
		} else {
			return nil, errors.New("该文章状态已经无效!")
		}
	}
	return nil, err
}

func (s *ArticleService) PublishById(id int) (*models.Article, error) {
	s.article.Id = id
	if err = s.GetArticle(s.article); err == nil {
		if s.article.Status == models.StatusInActive {
			s.article.Status = models.StatusNormal
			if _, err = s.dbInstance.Update(s.article, "Status"); err == nil {
				s.article.FormatDatetime()
				return s.article, nil
			}
		} else {
			return nil, errors.New("该文章不允许发布!")
		}
	}
	return nil, err
}

func (s *ArticleService) ResetById(id int) (*models.Article, error) {
	if err = s.GetArticle(s.article); err == nil {
		if s.article.Status == models.StatusBanned {
			s.article.Status = models.StatusInActive
			if _, err = s.dbInstance.Update(s.article, "Status"); err == nil {
				s.article.FormatDatetime()
				return s.article, nil
			}
		} else {
			return nil, errors.New("该文章状态不允许重置")
		}
	}
	return nil, err
}

func (s *ArticleService) GetArticle(article *models.Article) error {
	return s.dbInstance.Read(article)
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		article:    models.NewArticle(),
		dbInstance: orm.NewOrm(),
	}
}

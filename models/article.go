package models

import (
	. "api/helpers"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id          int       `orm:"auto" json:"id"`
	Author      *User     `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Category    *Category `orm:"rel(fk);null;on_delete(set_null)" json:"category"`
	Title       string    `orm:"size(30);description(用户名);unique" json:"title"`
	Description string    `orm:"type(text);description(文章内容)" json:"description"`
	ImageUrl    string    `orm:"description(封面图)" json:"image_url"`
	BaseModel
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) Store() (int64, error) {
	o := orm.NewOrm()
	insertId, err := o.Insert(a)
	_, _ = o.LoadRelated(a, "Author")
	go CheckError(err, "Article载入Author关系报错:")

	if a.Category.Id != 0 {
		_, _ = o.LoadRelated(a, "Category")
		a.Category.FormatDatetime()
		go CheckError(err, "Article载入Author关系报错:")
	}
	a.FormatDatetime()
	a.Author.FormatDatetime()
	return insertId, err
}

func (a *Article) GetList(queryMap map[string]string, page int, limit int) ([]*Article, int64, error) {
	var (
		articles []*Article
		o        = orm.NewOrm()
		err      error
		q        orm.QuerySeter
		count    int64
	)
	q = o.QueryTable(a)
	if status, ok := queryMap["Status"]; ok && status != "0" {
		q = q.Filter("status", status)
	}
	_, err = q.Limit(limit, (page-1)*limit).RelatedSel().All(&articles)
	go CheckError(err, "获取文章列表报错:")
	count, err = q.Count()
	go CheckError(err, "获取文章列表数量报错:")
	for _, article := range articles {
		article.FormatDatetime()
		article.Author.FormatDatetime()
		article.Category.FormatDatetime()
	}
	return articles, count, err
}

func (a *Article) GetById(i int) (*Article, error) {
	var (
		o   = orm.NewOrm()
		err error
	)
	a.Id = i
	err = o.Read(a)
	go CheckError(err, "通过id获取文章")
	return a, err
}

func NewArticle() *Article {
	return &Article{}
}

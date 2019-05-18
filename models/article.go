package models

import (
	. "api/helpers"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id              int       `orm:"auto" json:"id"`
	Author          *User     `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Category        *Category `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Title           string    `orm:"size(30);description(用户名);unique" json:"title"`
	Description     string    `orm:"type(text);description(文章内容)" json:"description"`
	ImageUrl        string    `orm:"description(封面图)" json:"image_url"`
	BaseModel
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) Store() (int64, error) {
	o := orm.NewOrm()
	insertId, err := o.Insert(a)
	_, err = o.LoadRelated(a, "Author")
	go CheckError(err, "Article载入Author关系报错:")
	a.FormatDatetime()
	a.Author.FormatDatetime()
	return insertId, err
}

func NewArticle() *Article  {
	return &Article{}
}

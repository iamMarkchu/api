package models

import "time"

type Article struct {
	Id              int       `orm:"auto"`
	Title           string    `orm:"size(30);description(用户名);unique"`
	Description     string    `orm:"type(text);description(文章内容)"`
	Author          *User     `orm:"rel(fk);null;on_delete(set_null)"`
	ImageUrl        string    `orm:"description(封面图)"`
	Status          uint8     `orm:"default(1);description(文章状态)"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)"`
	CreatedAtFormat string    `orm:"-"`
	UpdatedAtFormat string    `orm:"-"`
}

func (a *Article) TableName() string {
	return "articles"
}

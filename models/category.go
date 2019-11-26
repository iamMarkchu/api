package models

import (
	. "api/helpers"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id       int    `orm:"auto" json:"id"`
	ParentId int    `orm:"description(父类别id)" json:"parent_id"`
	AuthorId int    `orm:"description(作者id)" json:"author_id"`
	Name     string `orm:"size(30);description(类别名称)" json:"name"`
	BaseModel
}

func (c *Category) Store() (int64, error) {
	var (
		o        = orm.NewOrm()
		err      error
		insertId int64
	)
	insertId, err = o.Insert(c)
	go CheckError(err, "Category插入报错:")
	return insertId, err
}

func NewCategory() *Category {
	return &Category{}
}

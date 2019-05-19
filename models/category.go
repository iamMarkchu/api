package models

import (
	. "api/helpers"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id         int       `orm:"auto"`
	ParentCate *Category `orm:"rel(fk);null;on_delete(set_null)" json:"parent_cate"`
	Author     *User     `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Name       string    `orm:"size(30);description(类别名称)" json:"name"`
	BaseModel
}

func (c *Category) Store() (int64, error) {
	var (
		o        = orm.NewOrm()
		err      error
		insertId int64
	)
	fmt.Println("cccc:", c)
	insertId, err = o.Insert(c)
	go CheckError(err, "Category插入报错:")
	if c.ParentCate.Id != 0 {
		_, err = o.LoadRelated(c, "ParentCate")
		c.ParentCate.FormatDatetime()
	}
	_, err = o.LoadRelated(c, "Author")
	c.FormatDatetime()
	c.Author.FormatDatetime()
	go CheckError(err, "Category载入关系报错:")
	return insertId, err
}

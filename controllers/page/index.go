package page

import (
	"api/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	var (
		o          = orm.NewOrm()
		articles   []*models.Article
		categories []*models.Category
		err        error
	)
	if _, err = o.QueryTable("categories").Limit(5).All(&categories); err != nil {
		logs.Emergency("1111111")
	}
	if _, err = o.QueryTable("articles").All(&articles); err != nil {
		logs.Emergency("1111111")
	}
	fmt.Println(articles)
	c.Data["pageCode"] = "index"
	c.Data["articles"] = articles
	c.Data["categories"] = categories
	c.TplName = "index/index.tpl"
}

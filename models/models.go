package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	StatusInActive = iota
	StatusNormal
	StatusBanned
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:chukui@tcp(localhost)/api_base?charset=utf8")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

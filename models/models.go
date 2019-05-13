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
	// orm.RunSyncdb("default", false, true)
	orm.RegisterModel(new(User))
	orm.Debug = true
}

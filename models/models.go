package models

import (
	"api/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	StatusInActive = iota
	StatusNormal
	StatusBanned
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	helpers.CheckError(err, "[RegisterDriver Error]")
	err = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	helpers.CheckError(err, "[RegisterDataBase Error]")
	orm.RegisterModel(new(User), new(Article))
	err = orm.RunSyncdb("default", false, true)
	helpers.CheckError(err, "[RunSyncdb Error]")
	orm.Debug, err = beego.AppConfig.Bool("ormdebug")
	helpers.CheckError(err, "[orm Debug Error]")
}

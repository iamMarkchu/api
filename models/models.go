package models

import (
	. "api/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	StatusNormal = iota + 1
	StatusInActive
	StatusBanned

	FormatTimeString = "2006-01-02 15:04:05"
)

type BaseModel struct {
	Status          uint8     `orm:"default(1);description(状态字段)" json:"status"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)" json:"-"`
	CreatedAtFormat string    `orm:"-" json:"created_at"`
	UpdatedAtFormat string    `orm:"-" json:"updated_at"`
}

func (m *BaseModel) FormatDatetime() {
	m.CreatedAtFormat = m.CreatedAt.Format(FormatTimeString)
	m.UpdatedAtFormat = m.UpdatedAt.Format(FormatTimeString)
}

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	go CheckError(err, "[RegisterDriver Error]")
	err = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	go CheckError(err, "[RegisterDataBase Error]")
	orm.RegisterModel(new(User), new(Article), new(Category))
	err = orm.RunSyncdb("default", false, true)
	go CheckError(err, "[RunSyncdb Error]")
	orm.Debug, err = beego.AppConfig.Bool("ormdebug")
	go CheckError(err, "[orm Debug Error]")
}

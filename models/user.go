package models

import (
	"api/controllers/requests"
	"api/helpers"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int    `orm:"auto"`
	UserName  string `orm:"size(30);description(用户名 )"`
	Password  string  `orm:"default('');description(密码)"`
	Email     string  `orm:"default('');description(电子邮箱)"`
	Mobile    string    `orm:"default('');description(手机号)"`
	Age       uint8     `orm:"default(0);description(年龄)"`
	Status    uint8     `orm:"default(1);description(状态字段)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserByName(username string) (*User, error) {
	o := orm.NewOrm()
	u.UserName = username
	err := o.Read(u, "UserName")
	return u, err
}

func (u *User) Register(r requests.RegisterRequest) (*User, error) {
	o := orm.NewOrm()
	u.UserName = r.UserName
	u.Password = helpers.MD5(r.Password)
	u.Email = r.Email
	u.Status = StatusNormal
	_, err := o.Insert(u)
	if err == nil {
		return u, err
	}
	return nil, err
}

func NewUser() *User {
	return &User{}
}

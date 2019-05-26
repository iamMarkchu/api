package models

import (
	"api/controllers/requests"
	"api/helpers"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int    `orm:"auto" json:"id"`
	UserName  string `orm:"size(30);description(用户名)" json:"user_name"`
	Password  string  `orm:"default('');description(密码)" json:"-"`
	Email     string  `orm:"default('');description(电子邮箱)" json:"email"`
	Mobile    string    `orm:"default('');description(手机号)" json:"mobile"`
	Age       uint8     `orm:"default(0);description(年龄)" json:"age"`
	BaseModel
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

func (u *User) GetUserById(id int) (*User, error)  {
	o := orm.NewOrm()
	u.Id = id
	err := o.Read(u)
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

package models

import (
	"api/controllers/requests"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@localhost/api_base?charset=utf8")
	orm.RegisterModel(new(User))
}

type User struct {
	ID        int
	UserName  string
	Password  string
	Email     string
	Mobile    string
	Age       uint8
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserByName(username string) (*User, error) {
	o := orm.NewOrm()
	u.UserName = username
	err := o.Read(u)
	return u, err
}

func (u *User) Register(r requests.RegisterRequest) (*User, error) {
	o := orm.NewOrm()
	u.UserName = r.UserName
	u.Password = r.Password
	u.Email = r.Email
	_, err := o.Insert(&u)
	if err == nil {
		return u, err
	}
	return nil, err
}

func NewUser() *User {
	return &User{}
}

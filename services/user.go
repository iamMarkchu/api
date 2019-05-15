package services

import (
	"api/controllers/requests"
	. "api/helpers"
	"api/helpers/jwt"
	"api/models"
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserService struct {
}

func (c *UserService) Register(r requests.RegisterRequest) (int, error) {
	// 判断输入密码是否一致
	if r.Password != r.RePassword {
		return 0, errors.New("两次输入密码不一致!")
	}
	// 判断email是否合法
	if matched, _ := CheckEmail(r.Email); !matched {
		return 0, errors.New("邮箱不合法!")
	}
	// 判断用户名是否存在
	userModel := models.NewUser()
	if _, err := userModel.GetUserByName(r.UserName); err == nil {
		return 0, errors.New("用户名已存在!")
	}
	// 注册用户
	u, err := userModel.Register(r)
	if err != nil {
		return 0, errors.New("创建用户失败!")
	}
	return u.Id, nil
}

func (u *UserService) Login(r requests.LoginRequest) (string, string, error) {
	if r.UserName == "" || r.Password == "" {
		return "", "", errors.New("用户名，密码不能为空!")
	}
	userModel := models.NewUser()
	user, err := userModel.GetUserByName(r.UserName)
	// 验证用户名是否存在
	if err == orm.ErrNoRows {
		return "", "", errors.New("用户名不存在!")
	}
	// 验证密码是否正确
	if user.Password != MD5(r.Password) {
		return "", "", errors.New("密码错误!")
	}
	// todo生成token
	auth := jwt.GetToken(strconv.Itoa(user.Id))
	return auth.Token, auth.ExpireIn, nil
}

func NewUserService() *UserService {
	return &UserService{}
}

package services

import (
	"api/controllers/requests"
	"api/helpers"
	"api/models"
	"errors"
)

type UserService struct {

}

func (c *UserService) Register(r requests.RegisterRequest) (int, error)  {
	// 判断输入密码是否一致
	if r.Password != r.RePassword {
		return 0, errors.New("两次输入密码不一致")
	}
	// 判断email是否合法
	if matched, _ := helpers.CheckEmail(r.Email); !matched {
		return 0, errors.New("邮箱不合法!")
	}
	// 判断用户名是否存在
	userModel := models.NewUser()
	if _, err := userModel.GetUserByName(r.UserName); err == nil {
		return 0, errors.New("用户名已存在")
	}
	// 注册用户
	u, err := userModel.Register(r)
	if err != nil {
		return 0, errors.New("创建用户失败!")
	}
	return u.ID, nil
}

func NewUserService() *UserService {
	return &UserService{}
}

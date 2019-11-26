package token

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	ExpireAt = 60 * 60 * 24
	Issuer   = "goldfish"
)

var (
	key = []byte("qwertyuiop")
)

type ApiClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type Auth struct {
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	ExpireIn string `json:"expire_at"`
}

func GetToken(userId string, role string) Auth {
	claims := &ApiClaims{
		userId,
		role,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + ExpireAt),
			Issuer:    Issuer,
		},
	}
	logs.Info("获取token:", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return Auth{}
	}
	return Auth{UserId: userId, Token: ss, ExpireIn: strconv.Itoa(60 * 60 * 24)}
}

// 校验token是否有效
func CheckToken(tokenStr string) (string, string, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &ApiClaims{}, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parse with claims failed.", err)
		return "", "", false
	}
	if claims, ok := token.Claims.(*ApiClaims); ok && token.Valid {
		return claims.UserId, claims.Role, true
	} else {
		return "", "", false
	}
}

package jwt

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	key = []byte("qwertyuiop")
)

type ApiClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type Auth struct {
	Token string `json:"token"`
	ExpireIn string  `json:"expire_at"`
}

func GetToken(Userid string) Auth {
	claims := &ApiClaims{
		Userid,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + (60 * 60 * 24)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return Auth{}
	}
	return Auth{Token:ss, ExpireIn: strconv.Itoa(60 * 60 * 24)}
}

// 校验token是否有效
func CheckToken(tokenStr string) (string, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &ApiClaims{}, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return "", false
	}
	if claims, ok := token.Claims.(*ApiClaims); ok && token.Valid {
		return claims.UserId, true
	} else {
		return "", false
	}
}

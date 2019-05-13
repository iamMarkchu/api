package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

// 判断邮箱是否合法
func CheckEmail(email string) (bool, error) {
	return regexp.MatchString(`^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$`, email)
}

// md5加密
func MD5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//

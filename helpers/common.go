package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
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

// 检查错误
func CheckError(err error, msg string)  {
	if err != nil {
		logs.Info("["+msg+"]:", err.Error())
	}
}

func GetErrorMap(errs []*validation.Error) map[string]string  {
	var errorMaps = make(map[string]string)
	for _, err := range errs {
		errorMaps[err.Key] = err.Message
	}
	return errorMaps
}

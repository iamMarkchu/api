package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"
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
func CheckError(err error, msg string) {
	if err != nil {
		logs.Info("["+msg+"]:", err.Error())
	}
}

func GetErrorMap(errs []*validation.Error) map[string]string {
	var errorMaps = make(map[string]string)
	for _, err := range errs {
		errorMaps[err.Key] = err.Message
	}
	return errorMaps
}

// 检查路径是否存在，是否可写
func CheckDirectory(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		go logs.Info(err.Error())
	}
}

// 获取唯一的文件名
func GetUniqueFileName(fileName string) string {
	return time.Now().Format("2006/01/02/") + MD5(fileName+time.Now().Format("150405000000")) + path.Base(fileName)
}

func JsonReturn(message string, result interface{}, code int) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"result":  "",
		"code":    http.StatusUnauthorized,
	}
}

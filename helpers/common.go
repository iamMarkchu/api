package helpers

import "regexp"

// 判断邮箱是否合法
func CheckEmail(email string) (bool, error) {
	return regexp.MatchString(`^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$`, email)
}

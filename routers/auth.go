package routers

import (
	"api/helpers/jwt"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

var auth = func(c *context.Context) {
	token := strings.TrimPrefix(c.Input.Header("Authorization"),"Bearer ")
	_, isValid := jwt.CheckToken(token)
	if !isValid {
		c.Output.Status = http.StatusUnauthorized
		c.Output.JSON(JsonReturn("toekn验证失败", "", http.StatusUnauthorized), true, true, )
		return
	}
}

func JsonReturn(message string, result interface{}, code int) map[string]interface{}  {
	return map[string]interface{}{
		"message": message,
		"result": "",
		"code": http.StatusUnauthorized,
	}
}

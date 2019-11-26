package routers

import (
	"api/helpers"
	"api/helpers/cache"
	"api/helpers/permission"
	"api/helpers/token"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
)

var auth = func(c *context.Context) {
	var (
		tokenStr string
		userId   string
		role     string
		isValid  bool
		err      error
	)
	if tokenStr, err = request.AuthorizationHeaderExtractor.ExtractToken(c.Request); err != nil {
		logs.GetLogger("【token】").Println("token err:", err.Error())
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(helpers.JsonReturn("token解析错误:"+err.Error(), nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("auth输出json错误:" + err.Error())
		}
		return
	} else if userId, role, isValid = token.CheckToken(tokenStr); !isValid {
		logs.Info("token err:", "token 不合法, "+tokenStr)
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(helpers.JsonReturn("token不合法", nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("auth 输出json错误:" + err.Error())
		}
		return
	} else if cache.GetCacheInstance().Get(helpers.MD5(tokenStr)) == nil {
		logs.Info("token err:", "用户不存在, "+tokenStr)
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(helpers.JsonReturn("用户不存在", nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("auth 输出json错误:" + err.Error())
		}
		return
	}
	c.Input.SetParam("authUserId", userId)
	c.Input.SetParam("authRole", role)

	logs.Info("userId:", userId, "role:", role, "resource:", c.Request.URL.Path, "action:", c.Request.Method)
	if permit := permission.EnForcerInstance().Enforce(role, c.Request.URL.Path, c.Request.Method); !permit {
		logs.Info("权限验证失败:没权限, 您的角色是:", role)
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(helpers.JsonReturn("权限验证失败:没权限,您的角色是:"+ role, nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("权限验证输出json错误:" + err.Error())
		}
		return
	}
	logs.Info("tokenStr:", tokenStr, "userId:", userId, "roles:", role)
}

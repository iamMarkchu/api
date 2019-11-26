package permission

import (
	"github.com/astaxie/beego/logs"
	"github.com/casbin/casbin"
)

var e *casbin.Enforcer

func init()  {
	e = casbin.NewEnforcer("helpers/permission/model.conf", "helpers/permission/policy.csv")
	logs.GetLogger("casbin").Println("casbin初始化")
}

func EnForcerInstance() *casbin.Enforcer  {
	return e
}

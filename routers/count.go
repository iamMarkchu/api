package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

// log各个接口的调用情况
var count = func(c *context.Context) {
	logs.GetLogger("api count").Println(c.Input.URI())
}

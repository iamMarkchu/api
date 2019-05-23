package log

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init()  {
	logs.EnableFuncCallDepth(true)
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"./test.log"}`)
	if err != nil {
		fmt.Println("设置日志处理报错:", err.Error())
	}
	logs.Async()
}

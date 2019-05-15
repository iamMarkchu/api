package cache

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var (
	Instance cache.Cache
	err           error
)

func GetCacheInstance() cache.Cache {
	return Instance
}

func init() {
	logs.Info("[CACHE INIT START]")
	Instance, err = cache.NewCache("redis", `{"key": "api","conn": ":6379"}`)
	if err != nil {
		logs.Error("[CACHE INIT]:", err.Error())
	}
	logs.Info("[CACHE INIT END]")
}

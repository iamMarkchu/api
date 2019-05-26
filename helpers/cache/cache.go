package cache

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var (
	Instance cache.Cache
	err      error
)

func GetCacheInstance() cache.Cache {
	return Instance
}

func init() {
	logs.Info("[CACHE INIT START]")
	fmt.Println("[Redis Config]:", beego.AppConfig.String("redisconn"))
	Instance, err = cache.NewCache("redis", beego.AppConfig.String("redisconn"))
	if err != nil {
		logs.Error("[CACHE INIT]:", err.Error())
	}
	logs.Info("[CACHE INIT END]")
}

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

func init() {
	var auth = func(c *context.Context) {
		if c.Input.Header("Authorization") == "" {
			c.Output.Status = http.StatusForbidden
		}
	}
	// public api
	ns := beego.NewNamespace("/public",
		beego.NSInclude(
			&controllers.PublicController{},
		),
	)
	// auth api
	ns2 := beego.NewNamespace("/api",
		beego.NSBefore(
			auth,
		),
		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
	)
	beego.AddNamespace(ns, ns2)
}

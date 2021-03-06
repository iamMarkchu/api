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
	"api/controllers/page"
	"github.com/astaxie/beego"
)

func init() {
	// page
	beego.Router("/", &page.IndexController{})
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
			count,
		),
		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		beego.NSNamespace("/categories",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSRouter("/upload", &controllers.UploadController{}),
	)
	beego.AddNamespace(ns, ns2)
}

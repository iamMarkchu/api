package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Publish",
            Router: `/:id/publish`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["api/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Reset",
            Router: `/:id/reset`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["api/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:CategoryController"] = append(beego.GlobalControllerRouter["api/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:PublicController"] = append(beego.GlobalControllerRouter["api/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:PublicController"] = append(beego.GlobalControllerRouter["api/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

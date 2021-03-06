// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"goFrame/controllers"
	"goFrame/controllers/own"
)


func init() {
	// ns := beego.NewNamespace("/otc",
	// 	beego.NSAutoRouter(&own.UserController{}),
	// 	beego.NSAutoRouter(&controllers.ChatController{}),
	// )
	// beego.AddNamespace(ns)
	beego.AutoRouter(&own.UserController{})
	beego.AutoRouter(&controllers.ChatController{})
	beego.AutoRouter(&controllers.AboutController{})
}

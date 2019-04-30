package routers

import (
	"GooooooooG/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/code", &controllers.CodeRunnerController{})
}

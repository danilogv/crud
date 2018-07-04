package routers

import (
	"crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/altera", &controllers.AlteraController{})
	beego.Router("/consulta", &controllers.ConsultaController{})
	beego.Router("/insere", &controllers.InsereController{})
	beego.Router("/remove", &controllers.RemoveController{})
}

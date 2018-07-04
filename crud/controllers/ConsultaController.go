package controllers

import (
	"crud/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

type ConsultaController struct {
	beego.Controller
}

func (c *ConsultaController) Get() {
	c.TplName = "consulta.tpl"
}

func (c *ConsultaController) Post() {
	c.TplName = "consulta.tpl"
	codigo, err := strconv.Atoi(c.GetString("codigo"))
	p := models.Produto{Id: codigo}
	o := orm.NewOrm()
	err = o.Read(&p)
	if err == orm.ErrNoRows {
		c.Data["msg"] = "Não existe produto com esse código !"
	} else {
		codigo := strconv.Itoa(p.Id)
		nome := p.Nome
		preco := decimal.NewFromFloat(p.Preco).StringFixed(2)
		c.Data["codigo"] = codigo
		c.Data["nome"] =  nome
		c.Data["preco"] = strings.Replace(preco, ".", ",", -1)
	}
}

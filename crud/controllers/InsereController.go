package controllers

import (
	"crud/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type InsereController struct {
	beego.Controller
}

func (c *InsereController) Get() {
	c.TplName = "insere.tpl"
}

func (c *InsereController) Post() {
	c.TplName = "insere.tpl"
	codigo := 0
	nome := c.GetString("nome")
	preco, err := strconv.ParseFloat(strings.Replace(c.GetString("preco"), ",", ".", -1), 64)
	p := models.Produto{codigo, nome, preco}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&p)
	if err != nil {
		o.Rollback()
		msg := "Falha na inserção\nErro :" + err.Error() + "\nCódigo : " + strconv.FormatInt(id, 64)
		c.Data["msg"] = msg
	} else {
		o.Commit()
		msg := "Cadastro realizado com sucesso !"
		c.Data["msg"] = msg
	}
}

package controllers

import (
	"crud/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

type AlteraController struct {
	beego.Controller
}

func (c *AlteraController) Get() {
	c.TplName = "altera.tpl"
}

func (c *AlteraController) Post() {
	c.TplName = "altera.tpl"
	switch c.GetString("botao") {
	case "Consultar":
		codigo, err := strconv.Atoi(c.GetString("codigo"))
		p := models.Produto{Id: codigo}
		o := orm.NewOrm()
		err = o.Read(&p)
		if err == orm.ErrNoRows {
			c.Data["msg1"] = "Não existe produto com esse código !"
		} else {
			codigo := strconv.Itoa(p.Id)
			nome := p.Nome
			preco := decimal.NewFromFloat(p.Preco).StringFixed(2)
			c.Data["codigo"] = codigo
			c.Data["nome"] = nome
			c.Data["preco"] = strings.Replace(preco, ".", ",", -1)
		}
		break
	case "Alterar":
		codigo, err := strconv.Atoi(c.GetString("codigo"))
		p := models.Produto{Id: codigo}
		o := orm.NewOrm()
		err = o.Read(&p)
		if err == orm.ErrNoRows {
			c.Data["msg1"] = "Não existe produto com esse código !"
		} else {
			nome := c.GetString("nome")
			preco, err := strconv.ParseFloat(strings.Replace(c.GetString("preco"), ",", ".", -1), 64)
			p = models.Produto{codigo, nome, preco}
			o.Begin()
			id, err := o.Update(&p)
			if err != nil {
				o.Rollback()
				c.Data["msg2"] = "Falha na alteração\nErro :" + err.Error() + "\nCódigo : " + strconv.FormatInt(id, 64)
			} else {
				o.Commit()
				c.Data["msg2"] = "Alteração feita com sucesso !"
			}
		}
		break
	}
}

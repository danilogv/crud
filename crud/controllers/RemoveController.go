package controllers

import (
	"crud/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

type RemoveController struct {
	beego.Controller
}

func (c *RemoveController) Get() {
	c.TplName = "remove.tpl"
}

func (c *RemoveController) Post() {
	c.TplName = "remove.tpl"
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
	case "Remover":
		codigo, err := strconv.Atoi(c.GetString("codigo"))
		p := models.Produto{Id: codigo}
		o := orm.NewOrm()
		err = o.Read(&p)
		if err == orm.ErrNoRows {
			c.Data["msg1"] = "Não existe produto com esse código !"
		} else {
			o.Begin()
			id, err := o.Delete(&p)
			if err != nil {
				o.Rollback()
				c.Data["msg2"] = "Falha na exclusão\nErro :" + err.Error() + "\nCódigo : " + strconv.FormatInt(id, 64)
			} else {
				o.Commit()
				c.Data["msg2"] = "Exclusão feita com sucesso !"
			}
		}
		break
	}
}

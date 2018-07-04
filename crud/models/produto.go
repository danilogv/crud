package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Id    int     `orm:"auto;pk"`
	Nome  string  `orm:"size(30);not null;unique"`
	Preco float64 `orm:"digits(10);decimals(2);not null"`
}

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/crud?charset=utf8")
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "/users/danilo/go/src/crud/crud.db")
	orm.RegisterModel(new(Produto))
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}
}

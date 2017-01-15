package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const AliasName string = "default"
const NameDriver string = "mysql"
const Name string = "root"
const Pass string = "12345"
const NameDB string = "campeonato"

const UrlDB string = Name + ":" + Pass + "@/" + NameDB + "?charset=utf8"

func init() {
	orm.RegisterDriver(
		NameDriver,
		orm.DRMySQL,
	)

	orm.RegisterModel(
		new(Team),
		new(Person),
	)

	orm.RegisterDataBase(
		AliasName,
		NameDriver,
		UrlDB,
		100,
	)
}

func ORM() orm.Ormer {
	o := orm.NewOrm()
	o.Using(AliasName)
	return o
}

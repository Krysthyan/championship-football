package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"championship-football/tools"
	"encoding/json"
	"github.com/fatih/structs"
	"reflect"
)

const AliasName string = "default"
const NameDriver string = "mysql"
const Name string = "root"
const Pass string = "12345"
const NameDB string = "championship"

const UrlDB string = Name + ":" + Pass + "@/" + NameDB + "?charset=utf8"

func init() {
	orm.RegisterDriver(
		NameDriver,
		orm.DRMySQL,
	)

	orm.RegisterModel(
		new(Championship),
		new(Team),
		new(Team_championship),
		new(Player),
		new(Player_team),
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

func ORM_INSERT(model interface{}) (mapB []byte) {
	var error_list []tools.ErrorChampionship
	var err_tx error

	switch structs.Name(model) {
	case "Player":
		modelOrm := reflect.ValueOf(model).Interface().(Player)
		_, err_tx = ORM().Insert(&modelOrm)
	case "Team":
		modelOrm := reflect.ValueOf(model).Interface().(Team)
		_, err_tx = ORM().Insert(&modelOrm)
	case "Player_team":
		modelOrm := reflect.ValueOf(model).Interface().(Player_team)
		_, err_tx = ORM().Insert(&modelOrm)
	case "Team_championship":
		modelOrm := reflect.ValueOf(model).Interface().(Team_championship)
		_, err_tx = ORM().Insert(&modelOrm)

	}
	tools.ListError(&error_list, err_tx)

	if len(error_list) != 0 {
		mapB, _ = json.Marshal(error_list)
	}else {
		mapB, _ = json.Marshal(model)
	}

	return
}

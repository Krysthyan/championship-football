package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"championship-football/tools"
	"encoding/json"
	"github.com/fatih/structs"
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

	name_struct := structs.Name(model)

	switch name_struct {
	case "Player":
		_, err_tx = ORM().Insert(getPlayer(model))
	case "Team":
		_, err_tx = ORM().Insert(getTeam(model))
	case "Player_team":
		_, err_tx = ORM().Insert(getPlayerTeam(model))
	case "Team_championship":
		_, err_tx = ORM().Insert(getTeamChampionship(model))

	}
	tools.ListError(&error_list, err_tx)

	if len(error_list) != 0 {
		mapB, _ = json.Marshal(error_list)
	}else {
		mapB, _ = json.Marshal(model)
	}

	return
}

func getPlayer(interface_ interface{}) *Player {
	var model Player

	jsons, _ := json.Marshal(interface_)
	json.Unmarshal(jsons, &model)
	return &model
}
func getTeam(interface_ interface{}) *Team {
	var model Team

	jsons, _ := json.Marshal(interface_)
	json.Unmarshal(jsons, &model)
	return &model
}
func getPlayerTeam(interface_ interface{}) *Player_team {
	var model Player_team

	jsons, _ := json.Marshal(interface_)
	json.Unmarshal(jsons, &model)
	return &model
}
func getTeamChampionship(interface_ interface{}) *Team_championship {
	var model Team_championship

	jsons, _ := json.Marshal(interface_)
	json.Unmarshal(jsons, &model)
	return &model
}
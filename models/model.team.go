package models

import (
	"encoding/json"
	"championship-football/tools"
)

type Team struct {
	Id   string    `json:"id" orm:"pk"`
	Name string `json:"name" `
}

func GetTeamList() (mapB []byte) {
	var teams []Team

	ORM().QueryTable("team").
		OrderBy("Id").
		All(&teams)

	mapB, _ = json.Marshal(teams)
	return
}

func GetTeam(id string) (mapB []byte)  {
	var team Team
	ORM().QueryTable("team").
		Filter("id", id).
		OrderBy("Id").
		All(&team)
	mapB, _ = json.Marshal(team)
	return
}

func InsertTeam(team Team) (mapB []byte) {
	var error_list []tools.ErrorChampionship

	_, err_tx := ORM().Insert(&team)
	tools.ListError(&error_list, err_tx)

	if len(error_list) != 0 {
		mapB, _ = json.Marshal(error_list)
	}else {
		mapB, _ = json.Marshal(team)
	}

	return
}

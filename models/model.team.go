package models

import (
	"encoding/json"
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

func InsertTeam(team Team) []byte {
	return ORM_INSERT(team)
}

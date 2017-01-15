package models

import "encoding/json"

type Team struct {
	Id   int    `json:"id" orm:"pk"`
	Name string `json:"name" `
}

func GetTeam() (mapB []byte) {
	var teams []Team

	ORM().QueryTable("team").
		OrderBy("Id").
		All(&teams)

	mapB, _ = json.Marshal(teams)
	return
}

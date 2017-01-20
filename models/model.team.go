package models

import 	"encoding/json"

type Team struct {
	Id   string    `json:"id" orm:"pk"`
	Name string `json:"name" `
}

func GetListTeam() (mapB []byte) {
	var teams []Team

	ORM().QueryTable("team").
		OrderBy("Id").
		All(&teams)

	mapB, _ = json.Marshal(teams)
	return
}

func SaveTeamOrm(team Team)  {
	ORM().Insert(&team)
}

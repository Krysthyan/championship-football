package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type Team struct {
	Id   string `json:"id" orm:"pk"`
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

func GetTeam(id string) (mapB []byte) {
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

func GetTeam16() (mapB []byte) {
	var listValues []orm.ParamsList
	var teams []Team

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("team").OrderBy("RAND()").Limit(16)
	ORM().Raw(qb.String()).ValuesList(&listValues)

	for _, element := range listValues {
		teams = append(teams, Team{
			Name: element[1].(string),
			Id:   element[0].(string),
		})
	}
	mapB, _ = json.Marshal(teams)
	return
}

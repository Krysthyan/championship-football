package models

import (
	"encoding/json"
)

type Person struct {
	Id        string `json:"id" orm:"pk"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Classtype string `json:"classtype"`
	Team      *Team  `json:"team" orm:"rel(fk)"`
	Position  string `json:"position"`
	Number    int    `json:"number"`
}

func GetPerson() []byte {
	var person []Person

	ORM().QueryTable("person").
		OrderBy("Id").
		All(&person)

	return GetJsonPerson(person)
}

func GetPersonTeam(filter_team_id string) []byte {
	var person []Person

	ORM().QueryTable("person").
		Filter("team_id", filter_team_id).
		OrderBy("Id").
		All(&person)

	return GetJsonPerson(person)
}

func GetJsonPerson(person []Person) (mapB []byte) {
	for _, element := range person {
		ORM().Read(element.Team)
	}
	mapB, _ = json.Marshal(person)
	return
}

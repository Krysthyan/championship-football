package models

import (
	"championship-football/tools"
	"encoding/json"
	"log"
)

type Championship struct {
	Id   string `json:"id" orm:"column(id);pk"`
	Name string `json:"name" orm:"size(5)"`
}

func SaveChampionship(championship Championship) (mapB []byte) {
	_, err_orm := ORM().Insert(&championship)
	err := tools.ChampionshipError(err_orm, &championship)

	if err != "" {
		mapB, _ = json.Marshal(tools.ErrorChampionship{Error: err})
	}
	return
}

func GetChampionship(id string) (mapB []byte) {
	var championship Championship
	ORM().QueryTable("championship").
		Filter("Id", id).
		OrderBy("Id").
		All(&championship)

	mapB, _ = json.Marshal(championship)
	return mapB
}
func GetChampionshipName(name string)(mapB []byte) {
	var championship Championship
	ORM().QueryTable("championship").
		Filter("Name", name).
		OrderBy("Id").
		One(&championship)

	mapB, _ = json.Marshal(championship)
	return mapB
}

func GetChampionshipList() []byte {
	var championship []Championship
	ORM().QueryTable("championship").
		OrderBy("Id").
		All(&championship)

	return GetJsonChampionship(championship)
}

func DeleteChampionship(championship Championship) {
	if _, err := ORM().Delete(&championship); err == nil {
		log.Println(err)
	}
}

func UpdateChampionship(id string, name string) {
	championship := Championship{Id: id}
	if ORM().Read(&championship) == nil {
		championship.Name = name

		if _, err := ORM().Update(&championship); err == nil {
			log.Println(err)
		}
	}
}

func GetJsonChampionship(championship []Championship) (mapB []byte) {
	mapB, _ = json.Marshal(championship)
	return
}

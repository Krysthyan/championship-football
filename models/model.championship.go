package models

import (
	"championship-football/tools"
	"log"
	"encoding/json"
)

type Championship struct {
	Id string `json:"id" orm:"size(5);ref(pk)"`
	Name string `json:"name" orm:"size(5)"`
	Team []Team `orm:"reverse(many)"`
}

func SaveChampionship(championship Championship)  {
	_,err_orm := ORM().Insert(&championship)
	err := tools.ChampionshipError(err_orm, &championship)

	if err != "" {
		log.Println(err)
	}
}

func GetChampionship(id string) []byte  {
	var championship Championship
	ORM().QueryTable("championship").
		Filter("Id", id).
		OrderBy("Id").
		All(&championship)

	return GetJsonChampionship(championship)
}

func GetChampionshipList() []byte  {
	var championship []Championship
	ORM().QueryTable("championship").
		OrderBy("Id").
		All(&championship)

	return GetJsonChampionship(championship)
}

func DeleteChampionship(id string)  {
	if _, err := ORM().Delete(&Championship{Id:id}); err == nil {
		log.Println(err)
	}
}

func UpdateChampionship(id string, name string)  {
	championship := Championship{Id:id}
	if ORM().Read(&championship) == nil {
		championship.Name = name

		if _, err := ORM().Update(&championship); err == nil {
			log.Println(err)
		}
	}
}

func GetJsonChampionship(championship []Championship) (mapB []byte) {
	for _, element := range championship {
		ORM().Read(element.Team)
	}
	mapB, _ = json.Marshal(championship)
	return
}

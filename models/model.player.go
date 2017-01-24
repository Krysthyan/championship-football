package models

import (
	"championship-football/tools"
	"encoding/json"
)

type Player struct {
	Id        string `json:"id" orm:"pk"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Direction string `json:"direction"`
	Position  string `json:"position"`
	Number    int    `json:"number"`
}


func InsertPlayer(player Player) (mapB []byte) {
	var error_list []tools.ErrorChampionship

	_, err_tx := ORM().Insert(&player)
	tools.ListError(&error_list, err_tx)

	if len(error_list) != 0 {
		mapB, _ = json.Marshal(error_list)
	}else {
		mapB, _ = json.Marshal(player)
	}

	return
}

func GetPlayer(id string) (mapB []byte)  {
	var player Player

	ORM().QueryTable("player").
	Filter("id", id).
	OrderBy("Id").
	One(&player)

	mapB, _= json.Marshal(player)

	return
}
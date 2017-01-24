package models

import (
	"log"
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


func InsertPlayer(player Player)  {
	_, err_orm :=ORM().Insert(&player)
	err :=tools.ChampionshipError(err_orm, player)

	if err != "" {
		log.Println(err)
	}
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
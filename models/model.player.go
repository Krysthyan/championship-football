package models

import (
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


func InsertPlayer(player Player) []byte {
	return ORM_INSERT(player)
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

func GetPlayerList() (mapB []byte) {
	var players []Player

	ORM().QueryTable("player").
	OrderBy("Id").All(&players)

	mapB, _= json.Marshal(players)
	return
}
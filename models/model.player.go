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


func InsertPlayer(player Player, championship_id string, team_id string) (mapB []byte) {
	var error_list []tools.ErrorChampionship

	ormPlayer := ORM()

	err_tx := ormPlayer.Begin()
	tools.ListError(&error_list, err_tx)

	_, err_tx = ormPlayer.Insert(&player)
	tools.ListError(&error_list, err_tx)

	_, err_tx= ormPlayer.Raw("INSERT INTO player_team " +
		"(team_id, championship_id, player_id) " +
		"VALUES(?,?,?)",
		team_id, championship_id, player.Id,
	).Exec()
	tools.ListError(&error_list, err_tx)

	if len(error_list) != 0 {
		mapB, _ = json.Marshal(error_list)
		ormPlayer.Rollback()
	}else {
		mapB, _ = json.Marshal(player)
		ormPlayer.Commit()
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
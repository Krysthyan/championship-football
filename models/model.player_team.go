package models

import (
	"encoding/json"
)

type Player_team struct {
	Team_id string `json:"team_id" orm:"pk; ref(fk)"`
	Player_id string `json:"player_id" orm:"ref(fk)"`
	Championship_id string `json:"championship_id" orm:"ref(fk)"`
}

func InsertPlayerTeam(player_team Player_team) []byte {
	return ORM_INSERT(player_team)
}

func GetPlayersFromTeam(id_team, id_championship string) (mapB []byte) {
	var player_team []Player_team
	var players []Player

	ORM().QueryTable("player_team").
		Filter("team_id", id_team).
		Filter("championship_id", id_championship).
		OrderBy("Team_id").
		All(&player_team)

	for _, element := range player_team {
		player := Player{Id:element.Player_id}
		ORM().Read(&player)
		players = append(players, player)
	}

	mapB, _= json.Marshal(players)

	return
}
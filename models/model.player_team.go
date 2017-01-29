package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"reflect"
)

type Player_team struct {
	Team_id         string `json:"team_id" orm:"pk; ref(fk)"`
	Player_id       string `json:"player_id" orm:"ref(fk)"`
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
		player := Player{Id: element.Player_id}
		ORM().Read(&player)
		players = append(players, player)
	}

	mapB, _ = json.Marshal(players)

	return
}

func GetIdplayerRamdonDB(Team_id string, Championship_id string) (player_id string) {
	var listValues []orm.ParamsList
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("player_team").
		Where("Team_id = ? and Championship_id = ? ").
		OrderBy("RAND()").Limit(1)
	ORM().Raw(qb.String(),Team_id, Championship_id).ValuesList(&listValues)

	player_id = listValues[0][0].(string)

	return
}

func AssingPlayersTeams(team Team, Championship_id string, numAssingPlayers int) {
	var return_orm interface{}

	for numAssingPlayers > 0 {
		player_team := Player_team{
			Player_id:GetPlayerRamdonDB().Id,
			Championship_id:Championship_id,
			Team_id:team.Id,
		}

		json.Unmarshal(InsertPlayerTeam(player_team), &return_orm)

		if reflect.ValueOf(return_orm).Len() > 1 {
			numAssingPlayers--
		}
	}
}

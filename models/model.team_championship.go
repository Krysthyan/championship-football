package models

import (
	"encoding/json"
	"championship-football/tools"
	"errors"
	"strconv"
)

type Team_championship struct {
	Championship_id string`json:"championship_id" orm:"pk;ref(fk)"`
	Team_id string `json:"team_id" orm:"ref(fk)"`

}

func InsertTeamChampionship(team_championship Team_championship) []byte {
	return ORM_INSERT(team_championship)
}

func GetTeamsFromChampionship(id string) (mapB []byte)  {
	var team_championship []Team_championship
	var team []Team

	ORM().QueryTable("team_championship").
		Filter("championship_id", id).
		OrderBy("team_id").
		All(&team_championship)

	for _, element := range team_championship {
		team_tem := Team{Id:element.Team_id}
		ORM().Read(&team_tem)
		team = append(team, team_tem)
	}
	mapB, _ = json.Marshal(team)
	return
}

func AsignarEquipos(id string) (mapB []byte) {
	var error_list []tools.ErrorChampionship
	var team_championship []Team_championship
	var teams []Team
	json.Unmarshal(GetTeam16(), &teams)
	if len(teams) == 16 {
		for _, element := range teams{
			team_championship = append(team_championship, Team_championship{
				Team_id:element.Id,
				Championship_id:id,
			})
		}
		_, err_tx := ORM().InsertMulti(16, team_championship)
		tools.ListError(&error_list, err_tx)
	} else {
		tools.ListError(&error_list,errors.New(
			"No existe el limite exigido de equipos :" +
				" " + strconv.Itoa(len(teams))),
		)
	}
	if len(error_list) != 0 {
		mapB, _= json.Marshal(error_list)
	}else{
		mapB, _= json.Marshal(teams)
	}
	return
}



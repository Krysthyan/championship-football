package models

import "encoding/json"

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



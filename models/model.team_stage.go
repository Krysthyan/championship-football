package models

import (
	"encoding/json"
	"log"
)

type Team_stage struct {
	Team_id string `json:"team_id" orm:"pk"`
	Stage_id string `json:"stage_id"`
}

func InsertTeamStage(teamStage Team_stage) []byte {
	return ORM_INSERT(teamStage)
}

func GetLevelChampionship(Championship_id string) (level int) {
	ORM().Raw("CALL getRaundHigher(?)", Championship_id).QueryRow(&level)

	return
}

func GetTeamFromStage(Stage_id string) (mapB []byte) {
	var team_stage []Team_stage
	var listTeam_stage ListTeamStage

	ORM().QueryTable("team_stage").Filter("stage_id",  Stage_id).All(&team_stage)
	var teams []Team
	for _, element := range team_stage{

		var team Team
		json.Unmarshal(GetTeam(element.Team_id),&team)
		teams = append(teams, team)
	}
	log.Println(teams)
	listTeam_stage = ListTeamStage{
		Id:team_stage[0].Stage_id,
		Team:teams,
	}
	mapB, _ = json.Marshal(listTeam_stage)
	return
}

func IsChampionshipInTable(id string) (is_exist bool) {

	count, _ := ORM().QueryTable("stage_team").Filter("championship_id", id).Count()

	if int(count) > 0 {
		is_exist = true
	}
	return
}

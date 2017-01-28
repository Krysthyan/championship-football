package models

import (
	"math"
	"encoding/json"
	"championship-football/tools"
	"time"
	"math/rand"
	"errors"
)

type Stage struct {
	Id string `json:"id" orm:"pk"`
	Round int `json:"round"`
	IsStage bool `json:"is_stage"`
}

type ListTeamStage struct {
	Id string `json:"id"`
	Team []Team `json:"team"`
}

func InsertStage(stage Stage) []byte {
	return ORM_INSERT(stage)
}

func GenerateFaseGrupos(teams []Team, championship_id string) (mapB []byte) {
	var listErrors []tools.ErrorChampionship
	var listTeamStage []ListTeamStage

	x := math.Log(float64(len(teams))) / math.Log(2)

	if (x - float64(int(x)) ) == 0 && (int(x) > 2){

		var stage Stage
		newSource := rand.NewSource(time.Now().Unix())
		randNew := rand.New(newSource)
		var isNewGoup int
		var teams_alt []Team
		for len(teams) > 0 {
			ramdon := randNew.Intn(len(teams))

			if isNewGoup == 0 || isNewGoup == 4 {
				if isNewGoup == 4 {
					listTeamStage = append(listTeamStage, ListTeamStage{
						Id:stage.Id,
						Team:teams_alt,
					})
					teams_alt = nil
				}

				json.Unmarshal(InsertStage(GetNewStage()),&stage)
				isNewGoup = 0
			}
			team := teams[ramdon]
			InsertTeamStage(Team_stage{
				Championship_id:championship_id,
				Stage_id:stage.Id,
				Team_id:team.Id,
			})

			teams = append(teams[:ramdon],teams[ramdon+1:]...)
			teams_alt = append(teams_alt, team)
			isNewGoup++


		}
		listTeamStage = append(listTeamStage, ListTeamStage{
			Id:stage.Id,
			Team:teams_alt,
		})
	} else {

		tools.ListError(&listErrors, errors.New("No se ha ingresado el " +
			"número permitido de equipos, ejm 8, 16, 32, 64 ..."))

	}

	if len(listErrors) == 0 {
		mapB, _ = json.Marshal(listTeamStage)
	} else {
		mapB, _ = json.Marshal(listErrors)
	}

	return
}

func GetNewStage() Stage {
	return Stage{
		Id:tools.ChampionshipToken(3),
		Round:0,
		IsStage:false,
	}
}
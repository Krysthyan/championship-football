package models

import (
	"math"
	"encoding/json"
	"championship-football/tools"
	"time"
	"math/rand"
	"errors"
	"log"
)

type Stage struct {
	Id string `json:"id" orm:"pk"`
	Round int `json:"round"`
	IsStage int `json:"is_stage"`
	Championship_id string `json:"championship_id"`
}

type ListTeamStage struct {
	Id string `json:"id"`
	Team []Team `json:"team"`
}

func InsertStage(stage Stage) []byte {
	return ORM_INSERT(stage)
}

func UpdateStage(Stage_id string, ImcrementRound int)  {
	stage := Stage{Id: Stage_id}
	if ORM().Read(&stage) == nil {
		stage.Round += ImcrementRound

		if _, err := ORM().Update(&stage); err == nil {
			log.Println(err)
		}
	}
}

func GenerateFaseGrupos(teams []Team, championship_id string) (mapB []byte) {
	var listErrors []tools.ErrorChampionship
	var listTeamStage []ListTeamStage

	x := math.Log(float64(len(teams))) / math.Log(2)

	if (x - float64(int(x)) ) == 0 && (int(x) > 2){

		newSource := rand.NewSource(time.Now().Unix())
		randNew := rand.New(newSource)
		var isNewGoup int

		var teams_alt []Team
		var stage Stage

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

				json.Unmarshal(InsertStage(GetNewStage(championship_id)),&stage)
				isNewGoup = 0
			}

			team := teams[ramdon]
			InsertTeamStage(Team_stage{
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


func PlayARound(Championship_id string) (mapB []byte) {
	var listTeamStage []ListTeamStage
	levelChampionship := GetLevelChampionship(Championship_id)

	json.Unmarshal(GetListStageFromTeam(Championship_id), & listTeamStage)

	for _, element := range listTeamStage{
		combinations := CombinationsTeam(element.Team)
		Play(combinations[levelChampionship], element.Id, Championship_id)
		Play(combinations[5 - levelChampionship], element.Id, Championship_id)
		UpdateStage(element.Id, 1)
	}

	mapB, _ = json.Marshal("Todo bien....")

	return

}

func Play(combination Combinations, Stage_id string, Championship_id string) {
	rand.Seed(time.Now().UnixNano())

	goalEquipo1 := rand.Intn(5)
	goalEquipo2 := rand.Intn(5)

	var match Match
	match.Id = tools.ChampionshipToken(3)
	match.Is_draw = 0
	match.Stage_id = Stage_id

	log.Println(goalEquipo1, goalEquipo2)

	if goalEquipo1 > goalEquipo2 {

		match.Team_winner = combination.team1.Id
		match.Team_losser = combination.team2.Id

	} else if goalEquipo1 < goalEquipo2 {

		match.Team_winner = combination.team2.Id
		match.Team_losser = combination.team1.Id

	}else {
		match.Is_draw = 1
		match.Team_winner = combination.team1.Id
		match.Team_losser = combination.team2.Id
	}

	json.Unmarshal(InsertMatch(match),&match)

	for index:=0 ; index < goalEquipo1; index++ {
		SaveUpdateGoal(match.Id, combination.team1.Id, Championship_id)
	}
	for index:=0 ; index < goalEquipo2; index++ {
		SaveUpdateGoal(match.Id, combination.team2.Id, Championship_id)
	}

}


func SaveUpdateGoal(Match_id string, Team_id string, Championship_id string)  {
	player_id := GetIdplayerRamdonDB(Team_id, Championship_id )
	goal := Goal{
		Match_id:Match_id,
		Player_id:player_id,
	}
	if IsExistPlayerGoal(goal){
		UpdateGoal(goal)
	}else {
		goal.Number_goals = 1
		InsertGoal(goal)
	}
}

func GetListStageFromTeam(Championship_id string) (mapB []byte) {
	var stages []Stage
	var listTeamStage []ListTeamStage
	json.Unmarshal(GetStageFromChampionship(Championship_id), &stages)

	for _, stage := range stages{
		var team_stage_list ListTeamStage
		json.Unmarshal(GetTeamFromStage(stage.Id),&team_stage_list)
		listTeamStage = append(listTeamStage, team_stage_list)
	}
	mapB, _ = json.Marshal(listTeamStage)
	return
}

func GetStageFromChampionship(Championship_id string) (mapB []byte) {
	var stage []Stage

	ORM().QueryTable("stage").
		Filter("championship_id", Championship_id).
		All(&stage)
	log.Println("estamos aqui")
	mapB, _ = json.Marshal(stage)

	return
}

func GetNewStage(Championship_id string) Stage {
	return Stage{
		Id:tools.ChampionshipToken(3),
		Round:0,
		IsStage:0,
		Championship_id:Championship_id,
	}
}
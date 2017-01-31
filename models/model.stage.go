package models

import (
	"math"
	"encoding/json"
	"championship-football/tools"
	"time"
	"math/rand"
	"errors"
	"log"
	"github.com/astaxie/beego/orm"
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

type Team_winner struct {
	Id_stage string `json:"id_stage"`
	Id_team string `json:"id_team"`
	Championship_id string `json:"championship_id"`
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


	/* len(numTeams) / 2^(Round - 3) = numTems , equipos que pertenecen a la ronda*/
	num_teams := (float64( CountTeamInChampionship(Championship_id) ))/(math.Pow(2.0, (float64(levelChampionship) - 3.0)))

	if levelChampionship < 3 {
		log.Println("estamos en la opcopmm 1")
		json.Unmarshal(GetListStageFromTeam(Championship_id, "round__lt"), & listTeamStage)

		for _, element := range listTeamStage{
			combinations := CombinationsTeam(element.Team)
			Play(combinations[levelChampionship], element.Id, Championship_id)
			Play(combinations[5 - levelChampionship], element.Id, Championship_id)
			UpdateStage(element.Id, 1)
		}
		mapB, _ = json.Marshal("Todo bien....")

	}else if levelChampionship == 3 {
		log.Println("estamos en la opcopmm 2")
		mapB = PlayOff(Championship_id, levelChampionship + 1)
	}else if levelChampionship > 3 && num_teams > 1 {

		mapB = GetWinner(Championship_id, int(num_teams),levelChampionship + 1)

	}else if num_teams == 1 {
		mapB, _ = json.Marshal("El campeonato a finalizado")
	}
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

func GetListStageFromTeam(Championship_id string, condition string) (mapB []byte) {
	var stages []Stage
	var listTeamStage []ListTeamStage
	json.Unmarshal(GetStageFromChampionship(Championship_id, condition), &stages)

	for _, stage := range stages{
		var team_stage_list ListTeamStage
		json.Unmarshal(GetTeamFromStage(stage.Id),&team_stage_list)
		listTeamStage = append(listTeamStage, team_stage_list)
	}
	mapB, _ = json.Marshal(listTeamStage)
	return
}

func GetStageFromChampionship(Championship_id string, condition string) (mapB []byte) {
	var stage []Stage

	ORM().QueryTable("stage").
		Filter("championship_id", Championship_id).Filter(condition, 4).
		All(&stage)
	mapB, _ = json.Marshal(stage)

	return
}


func PlayOff(Championship_id string, level int) []byte {
	log.Println(Championship_id)
	num_teams := CountTeamInChampionship(Championship_id)

	return GetWinner(Championship_id, num_teams, level)
}

func GetWinner(Championship_id string, num_team int, level int) (mapB []byte){
	log.Println(" \n\n\n\n")
	log.Println(num_team)
	var team_winners []Team_winner
	var combinations []Combinations
	var listValues []orm.ParamsList

	stage := Stage{
		Championship_id:Championship_id,
		Id:tools.ChampionshipToken(3),
		IsStage:0,
		Round:level,
	}

	qb, _ := orm.NewQueryBuilder("mysql")
	if level == 4 {
		qb.Select("*").From("get_winner").Where("Championship_id = ? ").Limit(num_team/2)
	} else {
		qb.Select("*").From("get_winner").Where("Championship_id = ? ").OrderBy("Round").Desc().Limit(num_team)
	}

	ORM().Raw(qb.String(), Championship_id).ValuesList(&listValues)

	for _, element := range listValues {

		team_winners = append(team_winners, Team_winner{
			Id_team:element[0].(string),
			Championship_id:element[1].(string),
		})
	}
	InsertStage(stage)

	for index:= 0; index < num_team / 4; index ++ {
		var team1 Team
		var team2 Team
		json.Unmarshal(GetTeam(team_winners[index].Id_team), &team1)
		json.Unmarshal(GetTeam(team_winners[( (num_team / 2) - 1 ) - index].Id_team), &team2)

		InsertTeamStage(Team_stage{
			Stage_id:stage.Id,
			Team_id:team1.Id,
		})

		InsertTeamStage(Team_stage{
			Stage_id:stage.Id,
			Team_id:team2.Id,
		})


		combination := Combinations{
			team1:team1,
			team2:team2,
		}
		log.Println("\n\n\n")
		log.Println(combination)
		log.Println("\n\n\n")
		Play(combination, stage.Id, Championship_id)

		combinations = append(combinations, combination)
	}
	log.Println(combinations)
	mapB, _ = json.Marshal(combinations)

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


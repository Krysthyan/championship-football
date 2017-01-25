package controllers

import (
	"net/http"
	"strconv"
	"championship-football/models"
	"encoding/json"
	"championship-football/tools"
	"log"
)

func GetTeamList(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "/team/list", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetTeamList())
}

func SaveTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("POST", "team/save", strconv.Itoa(http.StatusOK))
	var team models.Team

	err := json.NewDecoder(h.Body).Decode(&team)

	if err != nil {
		panic(err)
	}
	team.Id = tools.ChampionshipToken(5)

	w = Set_ResponseWrite(w)
	w.Write(models.InsertTeam(team))
}

func GetTeamsFromChampionship(w http.ResponseWriter, h *http.Request){
	LogChampionship("GET", "team/getFromChampionship", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetTeamsFromChampionship(h.URL.Query().Get("championship_id")))
}

func AssingTeamToChampionship(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("POST", "team/assingTeamToChampionship", strconv.Itoa(http.StatusOK))
	var team_championship models.Team_championship
	err := json.NewDecoder(h.Body).Decode(&team_championship)

	if err != nil {
		log.Println(err)
	}
	w = Set_ResponseWrite(w)
	w.Write(models.InsertTeamChampionship(team_championship))
}



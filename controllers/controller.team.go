package controllers

import (
	"net/http"
	"strconv"
	"championship-football/models"
	"encoding/json"
	"championship-football/tools"
)

func GetTeamList(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "/team/list", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetTeamList())
}

func SaveTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("POST", "team/save", strconv.Itoa(http.StatusOK))
	parameter := h.URL.Query().Get("championship_id")
	var team models.Team

	err := json.NewDecoder(h.Body).Decode(&team)

	if err != nil {
		panic(err)
	}
	team.Id = tools.ChampionshipToken(5)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(models.InsertTeam(team, parameter))
}

func GetTeamsFromChampionship(w http.ResponseWriter, h *http.Request){
	LogChampionship("GET", "team/getFromChampionship", strconv.Itoa(http.StatusOK))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(models.GetTeamsFromChampionship(h.URL.Query().Get("championship_id")))
}



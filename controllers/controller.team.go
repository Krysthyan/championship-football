package controllers

import (
	"net/http"
	"strconv"
	"championship-football/models"
	"encoding/json"
	"championship-football/tools"
)

func GetListTeam(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "/team/list", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetListTeam())
}

func SaveTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("POST", "team/save", strconv.Itoa(http.StatusOK))

	var team models.Team
	err := json.NewDecoder(h.Body).Decode(&team)

	if err != nil {
		panic(err)
	}
	team.Id = tools.ChampionshipToken(5)
	models.SaveTeamOrm(team)


	w.Header().Set("Content-Type", "application/json")
	getJsonTeam, err := json.Marshal(team)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(getJsonTeam)
}

package controllers

import (
	"championship-football/models"
	"net/http"
	"strconv"
	"encoding/json"
	//"championship-football/tools"
)


func InsertPlayer(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "player/save", strconv.Itoa(http.StatusOK))

	var team models.Player
	err := json.NewDecoder(h.Body).Decode(&team)

	if err != nil {
		panic(err)
	}

	models.InsertPlayer(team)

	w.Header().Set("Content-Type", "application/json")
	getJsonTeam, err := json.Marshal(team)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(getJsonTeam)
}

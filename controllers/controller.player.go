package controllers

import (
	"championship-football/models"
	"net/http"
	"strconv"
	"encoding/json"
)


func InsertPlayer(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "player/save", strconv.Itoa(http.StatusOK))

	championship_id := h.URL.Query().Get("championship_id")
	team_id := h.URL.Query().Get("team_id")

	var team models.Player
	err := json.NewDecoder(h.Body).Decode(&team)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(models.InsertPlayer(
		team,
		championship_id,
		team_id,
	))
}

func GetPlayer(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "player/get", strconv.Itoa(http.StatusOK))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(models.GetPlayer(h.URL.Query().Get("player_id")))
}

func GetPlayersFromTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "player/getFromTeam", strconv.Itoa(http.StatusOK))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(models.GetPlayersFromTeam(
		h.URL.Query().Get("team_id"),
		h.URL.Query().Get("championship_id"),
	))
}


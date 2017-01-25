package controllers

import (
	"championship-football/models"
	"net/http"
	"strconv"
	"encoding/json"
	"log"
)


func InsertPlayer(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "player/save", strconv.Itoa(http.StatusOK))

	var player models.Player
	err := json.NewDecoder(h.Body).Decode(&player)

	if err != nil {
		log.Println(err)
	}

	w = Set_ResponseWrite(w)
	w.Write(models.InsertPlayer(player))
}

func GetPlayer(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "player/get", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayer(h.URL.Query().Get("player_id")))
}

func GetPlayersFromTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "player/getFromTeam", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayersFromTeam(
		h.URL.Query().Get("team_id"),
		h.URL.Query().Get("championship_id"),
	))
}

func AssingPlayerToTeam(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("POST", "player/assingPlayerToTeam", strconv.Itoa(http.StatusOK))
	var player_team models.Player_team
	err := json.NewDecoder(h.Body).Decode(&player_team)

	if err != nil {
		log.Println(err)
	}
	w = Set_ResponseWrite(w)
	w.Write(models.InsertPlayerTeam(player_team))
}

func GetPlayerList(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "player/getList", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayerList())

}


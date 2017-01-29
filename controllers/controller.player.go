package controllers

import (
	"championship-football/models"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"
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

func GetPlayer(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "player/get", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayer(h.URL.Query().Get("player_id")))
}

func GetPlayersFromTeam(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "player/getFromTeam", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayersFromTeam(
		h.URL.Query().Get("team_id"),
		h.URL.Query().Get("championship_id"),
	))
}

func AssingPlayerToTeam(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "player/assingPlayerToTeam", strconv.Itoa(http.StatusOK))
	var teams []models.Team
	json.NewDecoder(h.Body).Decode(&teams)
	numberElements, _ := strconv.Atoi(h.URL.Query().Get("num_players"))
	championship_id := h.URL.Query().Get("championship_id")

	for _, team := range teams{
		models.AssingPlayersTeams(team, championship_id,numberElements)
	}

	w = Set_ResponseWrite(w)

}

func GetPlayerList(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "player/getList", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetPlayerList())

}

func AddPlayersRandom(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "player/addPlayersRamdon", strconv.Itoa(http.StatusOK))

	numberElements, _ := strconv.Atoi(h.URL.Query().Get("num_players"))
	var players []interface{}
	var return_orm interface{}
	for numberElements > 0 {

		json.Unmarshal(models.InsertPlayer(models.GetPlayerRamdon()), &return_orm)

		if reflect.ValueOf(return_orm).Len() > 1 {
			players = append(players, return_orm)
			numberElements--
		}
	}

	return_json, _ := json.Marshal(players)
	w = Set_ResponseWrite(w)
	w.Write(return_json)
}

func GetPlayerListRandom(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "player/getPlayerListRandom", strconv.Itoa(http.StatusOK))
	numberElements, _ := strconv.Atoi(h.URL.Query().Get("num_players"))

	w = Set_ResponseWrite(w)

	w.Write(models.GetPlayerListRamdon(numberElements))
}

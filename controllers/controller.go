package controllers

import (
	"championship-football/models"
	"net/http"
	"strconv"
)

func GetTeam(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "localhost:8080/getTeam", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetTeam())

}

func GetPerson(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "localhost:8080/getPerson", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetPerson())

}

func GetPersonTeam(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "localhost:8080/getPersontTeam", strconv.Itoa(http.StatusOK))
	id_team, _ := strconv.Atoi(h.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetPersonTeam(id_team))

}

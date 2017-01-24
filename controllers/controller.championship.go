package controllers

import (
	"net/http"
	"strconv"
	"championship-football/models"
	"encoding/json"
	"log"
	"championship-football/tools"
)

func InsertChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "championship/save", strconv.Itoa(http.StatusOK))
	var championship models.Championship
	err := json.NewDecoder(h.Body).Decode(&championship)

	if err != nil {
		log.Println(err)
	}

	championship.Id = tools.ChampionshipToken(5)
	models.SaveChampionship(championship)
}

func GetChampionship(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "championship/get", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetChampionship(h.URL.Query().Get("id")))
}

func GetChampionshipList(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("GET", "championship/getList", strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(models.GetChampionshipList())
}
func DeleteChampionship(w http.ResponseWriter, h *http.Request)  {
	LogChampionship("DELETE", "championship/delete", strconv.Itoa(http.StatusOK))
	var championship models.Championship
	err := json.NewDecoder(h.Body).Decode(&championship)

	if err != nil {
		log.Println(err)
	}
	models.DeleteChampionship(championship)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mapB, _ := json.Marshal("delete")
	w.Write(mapB)

}
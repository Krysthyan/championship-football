package controllers

import (
	"championship-football/models"
	"championship-football/tools"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func InsertChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "championship/save", strconv.Itoa(http.StatusOK))
	var championship models.Championship
	championship.Id = tools.ChampionshipToken(5)
	err := json.NewDecoder(h.Body).Decode(&championship)

	if err != nil {
		log.Println(err)
	}

	w = Set_ResponseWrite(w)
	w.Write(models.SaveChampionship(championship))
}

func GetChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "championship/get", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetChampionship(h.URL.Query().Get("id")))
}

func GetChampionshipList(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "championship/getList", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetChampionshipList())
}

func DeleteChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("DELETE", "championship/delete", strconv.Itoa(http.StatusOK))

	var championship models.Championship
	err := json.NewDecoder(h.Body).Decode(&championship)

	if err != nil {
		log.Println(err)
	}

	models.DeleteChampionship(championship)

	w = Set_ResponseWrite(w)
	mapB, _ := json.Marshal("delete")
	w.Write(mapB)

}

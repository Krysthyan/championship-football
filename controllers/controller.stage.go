package controllers

import (
	"net/http"
	"strconv"
	"championship-football/models"
	"encoding/json"
)

func GenerateStage(w http.ResponseWriter, h *http.Request) {

	LogChampionship("POST", "stage/generateStage", strconv.Itoa(http.StatusOK))

	var teams []models.Team

	championship_id := h.URL.Query().Get("championship_id")
	err := json.NewDecoder(h.Body).Decode(&teams)
	if err != nil {
		panic(err)
	}
	w = Set_ResponseWrite(w)
	w.Write(models.GenerateFaseGrupos(teams, championship_id ))

}

func PlayGroupController(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "stage/playGroup", strconv.Itoa(http.StatusOK))

	championship_id := h.URL.Query().Get("championship_id")

	w = Set_ResponseWrite(w)

	w.Write(models.PlayARound(championship_id))

}

func GetListStageFromChampionship(w http.ResponseWriter, h *http.Request) {
	championship_id := h.URL.Query().Get("championship_id")
	w = Set_ResponseWrite(w)

	w.Write(models.GetListStageFromTeam(championship_id))
}

func GetTablePosicion(w http.ResponseWriter, h *http.Request) {
	stage_id := h.URL.Query().Get("stage_id")
	w = Set_ResponseWrite(w)

	w.Write(models.GetTablePosition(stage_id))
}
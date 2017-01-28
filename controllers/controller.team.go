package controllers

import (
	"championship-football/models"
	"championship-football/tools"
	"encoding/json"
	"github.com/icrowley/fake"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func GetTeamList(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "/team/list", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetTeamList())
}

func SaveTeam(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "team/save", strconv.Itoa(http.StatusOK))
	var team models.Team

	err := json.NewDecoder(h.Body).Decode(&team)

		if err != nil {
			panic(err)
		}
	team.Id = tools.ChampionshipToken(5)

	w = Set_ResponseWrite(w)
	w.Write(models.InsertTeam(team))
}

func GetTeamsFromChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("GET", "team/getFromChampionship", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.GetTeamsFromChampionship(h.URL.Query().Get("championship_id")))
}

func AssingTeamToChampionship(w http.ResponseWriter, h *http.Request) {
	LogChampionship("POST", "team/assingTeamToChampionship", strconv.Itoa(http.StatusOK))
	var listTeam_championship []models.Team_championship
	var team []models.Team
	 // comprobar que pasa si se envia solamente un equipo.
	err := json.NewDecoder(h.Body).Decode(&team)
	championship_id := h.URL.Query().Get("championship_id")

	for _, element := range team{
		var team_championship models.Team_championship
		json_ := models.InsertTeamChampionship(models.Team_championship{
			Team_id:element.Id,
			Championship_id:championship_id,
		})
		json.Unmarshal(json_,&team_championship)
		listTeam_championship = append(listTeam_championship,team_championship)
	}
	if err != nil {
		log.Println(err)
	}
	w = Set_ResponseWrite(w)
	mapB, _ := json.Marshal(listTeam_championship)
	w.Write(mapB)
}

func AsignarEquiposAleatorios(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "team/asignarEquipos", strconv.Itoa(http.StatusOK))

	w = Set_ResponseWrite(w)
	w.Write(models.AsignarEquipos(h.URL.Query().Get("championship_id")))

}

func AddTeamsRandom(w http.ResponseWriter, h *http.Request) {

	LogChampionship("GET", "team/addTeamsRamdon", strconv.Itoa(http.StatusOK))

	numberElements, _ := strconv.Atoi(h.URL.Query().Get("num_team"))
	var teams []interface{}
	var team models.Team
	var return_orm interface{}
	for numberElements > 0 {

		team.Id = tools.ChampionshipToken(3)
		team.Name = fake.City()
		json.Unmarshal(models.InsertTeam(team), &return_orm)

		if reflect.ValueOf(return_orm).Len() > 1 {
			teams = append(teams, return_orm)
			numberElements--
		}
	}

	return_json, _ := json.Marshal(teams)
	w = Set_ResponseWrite(w)
	w.Write(return_json)
}

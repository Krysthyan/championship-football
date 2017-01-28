package routers

import (
	"championship-football/controllers"
	"github.com/gorilla/mux"
)

func RouterTeam(championship *mux.Router) {

	championship.HandleFunc(
		"/team/list",
		controllers.GetTeamList,
	).Methods("GET")

	championship.HandleFunc(
		"/team/save",
		controllers.SaveTeam,
	).Methods("POST")

	championship.HandleFunc(
		"/team/getFromChampionship",
		controllers.GetTeamsFromChampionship,
	).Methods("GET")

	championship.HandleFunc(
		"/team/assingTeamToChampionship",
		controllers.AssingTeamToChampionship,
	).Methods("POST")

	championship.HandleFunc(
		"/team/asignarEquipos",
		controllers.AsignarEquiposAleatorios,
	).Methods("GET")

	championship.HandleFunc(
		"/team/addTeamsRamdon",
		controllers.AddTeamsRandom,
	).Methods("GET")

}

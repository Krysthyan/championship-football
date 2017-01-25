package routers

import (
	"github.com/gorilla/mux"
	"championship-football/controllers"
)


func RouterTeam(championship *mux.Router)  {

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

}

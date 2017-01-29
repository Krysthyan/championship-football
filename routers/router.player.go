package routers

import (
	"championship-football/controllers"
	"github.com/gorilla/mux"
)

func RouterPerson(championship *mux.Router) {
	championship.HandleFunc(
		"/player/save",
		controllers.InsertPlayer,
	).Methods("POST")

	championship.HandleFunc(
		"/player/get",
		controllers.GetPlayer,
	).Methods("GET")

	championship.HandleFunc(
		"/player/getFromTeam",
		controllers.GetPlayersFromTeam,
	).Methods("GET")

	championship.HandleFunc(
		"/player/assingPlayerToTeam",
		controllers.AssingPlayerToTeam,
	).Methods("POST")

	championship.HandleFunc(
		"/player/getList",
		controllers.GetPlayerList,
	).Methods("GET")

	championship.HandleFunc(
		"/player/addPlayersRandom",
		controllers.AddPlayersRandom,
	).Methods("GET")

	championship.HandleFunc(
		"/player/getPlayerListRandom",
		controllers.GetPlayerListRandom,
	).Methods("GET")
}

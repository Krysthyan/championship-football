package routers

import (
	"championship-football/controllers"
	"github.com/gorilla/mux"
)

func RouterChampionship(championship *mux.Router) {
	championship.HandleFunc(
		"/championship/save",
		controllers.InsertChampionship,
	).Methods("POST")

	championship.HandleFunc(
		"/championship/get",
		controllers.GetChampionship,
	).Methods("GET")

	championship.HandleFunc(
		"/championship/getFromName",
		controllers.GetChampionshipName,
	).Methods("GET")

	championship.HandleFunc(
		"/championship/getList",
		controllers.GetChampionshipList,
	).Methods("GET")

	championship.HandleFunc(
		"/championship/delete",
		controllers.DeleteChampionship,
	).Methods("POST")
}

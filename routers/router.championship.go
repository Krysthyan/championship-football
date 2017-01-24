package routers

import (
	"github.com/gorilla/mux"
	"championship-football/controllers"
)

func RouterChampionship(championship *mux.Router)  {
	championship.HandleFunc(
		"/championship/save",
		controllers.InsertChampionship,
	).Methods("POST")

	championship.HandleFunc(
		"/championship/get",
		controllers.GetChampionship,
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
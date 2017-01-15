package routers

import (
	"championship-football/controllers"
	"github.com/gorilla/mux"
)

func InitRouter(championship *mux.Router) *mux.Router {

	championship.HandleFunc(
		"/getTeam",
		controllers.GetTeam,
	).Methods("GET")

	championship.HandleFunc(
		"/getPerson",
		controllers.GetPerson,
	).Methods("GET")

	championship.HandleFunc(
		"/getPersonTeam",
		controllers.GetPersonTeam,
	).Methods("GET")

	return championship
}

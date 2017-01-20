package routers

import (
	"github.com/gorilla/mux"
	"championship-football/controllers"
)

func RouterTeam(championship *mux.Router)  {

	championship.HandleFunc(
		"/team/list",
		controllers.GetListTeam,
	).Methods("GET")

	championship.HandleFunc(
		"/team/save",
		controllers.SaveTeam,
	).Methods("POST")

}

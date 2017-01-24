package routers

import (
	"github.com/gorilla/mux"
	"championship-football/controllers"
)

func RouterPerson(championship *mux.Router)  {
	championship.HandleFunc(
		"/player/save",
		controllers.InsertPlayer,
	).Methods("POST")
}
